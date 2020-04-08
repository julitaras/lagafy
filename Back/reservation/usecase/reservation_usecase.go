package usecase

import (
	"api-dashboard/helpers"
	"api-dashboard/models"
	"api-dashboard/passenger"
	"api-dashboard/pkg/setting"
	"api-dashboard/reservation"
	"api-dashboard/travel"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	_ "golang.org/x/net/context/ctxhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

type reservationUsecase struct {
	rr reservation.Repository
	pr passenger.Repository
	tr travel.Repository
	t  time.Duration
}

//NewReservationUsecase returns reservation uscas
func NewReservationUsecase(rr reservation.Repository, pr passenger.Repository, tr travel.Repository, t time.Duration) reservation.Usecase {
	return &reservationUsecase{
		rr: rr,
		pr: pr,
		tr: tr,
		t:  t,
	}
}

func (uc *reservationUsecase) GetById(ctx context.Context, id string) (*models.Reservation, error) {
	res, err := uc.rr.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc *reservationUsecase) CheckIn(ctx context.Context, id uint) (*models.Reservation, error) {
	reserv, err := uc.GetById(ctx, strconv.Itoa(int(id)))

	if err != nil {
		return nil, err
	}

	if reserv.Status == helpers.Pending {
		return nil, errors.New("Su reserva se encuentra pendiente")
	} else if reserv.Status == helpers.OnBoard {
		return nil, errors.New("Ya has hecho el check-in")
	}

	now := time.Now().UTC()
	start := reserv.Travel.Departure
	end := reserv.Travel.Departure.Add(time.Minute * time.Duration(helpers.TimeForCheckIn))
	if !inTimeSpan(start, end, now) {
		return nil, errors.New("No es posible hacer el check in en este horario")
	}

	res, err := uc.rr.CheckIn(ctx, reserv)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc *reservationUsecase) Create(ctx context.Context, travelId string, email string, name string) (*models.Reservation, error) {
	t, err := uc.tr.GetById(ctx, travelId)
	p, err2 := uc.pr.GetOrCreate(ctx, email, name)
	reserv := models.Reservation{}

	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}

	if t.Departure.Add(time.Minute * time.Duration(helpers.TimeForCheckIn)).Before(time.Now().UTC()) {
		return nil, errors.New("No se pudo reservar. El viaje ya pasó.")
	}

	noon := time.Date(time.Now().UTC().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(), helpers.Noon, 0, 0, 0, time.UTC)
	limitDateTime := t.Departure.Add(-time.Minute * time.Duration(helpers.MinutesBeforeReservation))

	if t.Departure.After(noon) && time.Now().UTC().After(limitDateTime){
		return nil, errors.New("No se puede hacer una reserva para este viaje pasadas las " + strconv.Itoa(limitDateTime.Hour()) + "hs")
	}

	for _, r := range t.Reservations {
		if r.PassengerID == p.ID {
			return nil, errors.New("Ya tienes una reserva a tu nombre.")
		}
	}

	reserv.Travel = t
	reserv.TravelID = t.ID
	reserv.Passenger = *p
	reserv.PassengerID = p.ID

	if len(t.Reservations) >= t.Capacity {
		if err := SendMessage(&reserv, false); err != nil {
			return nil, err
		}
		reserv.Status = helpers.Pending
	} else {
		reserv.Status = helpers.Confirmed
	}
	res, err := uc.rr.Create(ctx, &reserv)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SendMessage(res *models.Reservation, isCanceled bool) error {
	var status string
	if isCanceled {
		status = "🚨 Reserva cancelada 🚨"
	} else {
		status = "🚨 Reserva pendiente 🚨"
	}
	message := status +
		"\nPasajero: " + res.Passenger.Name +
		"\nEmail: " + res.Passenger.Email +
		"\n📅 Fecha de viaje: " + res.Travel.Departure.Format(" 02-01-2006 15:04") +
		"\n📍Origen: " + res.Travel.Origin +
		"\n📍Destino: " + res.Travel.Destination

	type Recipient struct {
		ThreadKey string `json:"thread_key"`
	}

	type Message struct {
		Text string `json:"text"`
	}
	type RequestBody struct {
		Recipient *Recipient `json:"recipient"`
		Message   *Message   `json:"message"`
	}

	data := RequestBody{Message: &Message{Text: message}, Recipient: &Recipient{ThreadKey: setting.AppSetting.ThreadId}}

	requestBody, err := json.Marshal(data)
	req, err := http.NewRequest("POST", "https://graph.facebook.com/v4.0/me/messages", bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}

	log.Println(string(requestBody))

	req.Host = setting.AppSetting.Host
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+setting.AppSetting.AccessToken)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return nil
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func (uc *reservationUsecase) Delete(ctx context.Context, id int) (int, error) {
	reserv, err := uc.GetById(ctx, strconv.Itoa(int(id)))
	res, err := uc.rr.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	SendMessage(reserv, true)
	return res, nil
}

func (uc *reservationUsecase) GetListReservations(ctx context.Context, email string, name string) (*[]models.Reservation, error) {
	p, err1 := uc.pr.GetOrCreate(ctx, email, name)
	if err1 != nil {
		return nil, err1
	}
	passengerid := p.ID
	res, err := uc.rr.GetListReservations(ctx, strconv.Itoa(int(passengerid)))
	if err != nil {
		return nil, err
	}
	return res, nil
}
