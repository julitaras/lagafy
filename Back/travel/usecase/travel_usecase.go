package usecase

import (
	"api-dashboard/models"
	"api-dashboard/travel"
	"bytes"
	"context"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
	"time"
)

type travelUsecase struct {
	tr travel.Repository
	t  time.Duration
}

type Sender struct {
	user     string
	password string
}

//NewTravelUsecase returns travel usecas
func NewTravelUsecase(tr travel.Repository, t time.Duration) travel.Usecase {
	return &travelUsecase{
		tr: tr,
		t:  t,
	}
}

func (uc *travelUsecase) Insert(ctx context.Context, v *[]models.Travel) (*[]models.Travel, error) {
	travels := []models.Travel{}

	for _, travel := range *v {
		res, err := uc.tr.Insert(ctx, &travel)
		if err != nil {
			return &travels, err
		}
		travels = append(travels, *res)

	}
	return &travels, nil
}

func (uc *travelUsecase) GetById(ctx context.Context, id string) (*models.Travel, error) {
	res, err := uc.tr.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc *travelUsecase) GetCurrentsTravels(ctx context.Context) (*[]models.Travel, error) {
	res, err := uc.tr.GetCurrentsTravels(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc *travelUsecase) GetTemplates(ctx context.Context) (*[]models.Travel, error) {
	res, err := uc.tr.GetTemplates(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc *travelUsecase) Delete(ctx context.Context, id int) (int, error) {
	res, err := uc.tr.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (uc *travelUsecase) UpdateTravel(ctx context.Context, t *models.Travel) (*models.Travel, error) {
	res, err := uc.tr.UpdateTravel(ctx, t)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc *travelUsecase) GetTravelInfo(ctx context.Context, s string, e string) (*[]models.TravelInformation, error) {
	start, err := time.Parse("2006-01-02", s)
	end, err := time.Parse("2006-01-02", e)

	res, err := uc.tr.GetTravelInfo(ctx, start, end)

	if err != nil {
		return nil, err
	}

	response, err2 := ParseModel(res)

	if err2 != nil {
		return nil, err
	}

	return response, nil
}

func ParseModel(t *[]models.Travel) (*[]models.TravelInformation, error) {
	travelsInformation := []models.TravelInformation{}
	travelInformation := models.TravelInformation{}

	for _, travel := range *t {
		travelInformation.ID = travel.ID
		travelInformation.Capacity = travel.Capacity
		travelInformation.Arrival = travel.Arrival
		travelInformation.Departure = travel.Departure
		travelInformation.Origin = travel.Origin
		travelInformation.Destination = travel.Destination
		pending := 0
		onBoard := 0
		confirmed := 0
		cancelled := 0

		for _, reservation := range travel.Reservations {
			switch reservation.Status {
			case "pending":
				pending++
				break
			case "onboard":
				onBoard++
				break
			case "confirmed":
				confirmed++
				break
			case "cancelled":
				cancelled++
				break
			}
		}
		travelInformation.OnBoard = onBoard
		travelInformation.Cancelled = cancelled
		travelInformation.Pending = pending
		travelInformation.Confirmed = confirmed
		travelsInformation = append(travelsInformation, travelInformation)
	}

	return &travelsInformation, nil
}

func (uc *travelUsecase) Notify(ctx context.Context) (*[]models.Travel, error) {
	res, err := uc.tr.Notify(ctx)
	if err != nil {
		return nil, err
	}
	for _, t := range *res {
		for _, r := range t.Reservations {
			if r.Status != "onboard" {
				sender := NewSender("ejemplodemail@ejemplo.com", "contraseña")
				receiver := []string{r.Passenger.Email}

				subject := "Amistoso recordatorio de parte de Lagafy"
				message := `
				<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
				<html>
				<head>
				<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
				</head>
				<body>
				<p>Por favor, recordar hacer el check-in cuando subamos a la combi.</p>
				<p><b>Muchas gracias.</b></p>
				<br>
				<br>
				<p style="color:blue";>Equipo Lagafy</p>
				</body>
				</html>
				`
				bodyMessage := sender.WriteEmail(receiver, "text/html", subject, message)
				sender.SendMail(receiver, subject, bodyMessage)
			}
		}
	}
	return res, nil
}

//NewSender initialize sender
func NewSender(username, password string) Sender {

	return Sender{username, password}
}

//WriteEmail put mail style
func (sender Sender) WriteEmail(dest []string, contentType string, subject string, bodyMessage string) string {

	header := make(map[string]string)
	header["From"] = sender.user
	receipient := ""

	for _, user := range dest {
		receipient = receipient + user
	}
	header["To"] = receipient
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()
	message += "\r\n" + encodedMessage.String()

	return message
}

//SendMail send mail
func (sender Sender) SendMail(dest []string, subject string, bodyMessage string) {

	msg := "From: " + sender.user + "\n" +
		"To: " + strings.Join(dest, ",") + "\n" +
		"Subject: " + subject + "\n" + bodyMessage

	errors := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", sender.user, sender.password, "smtp.gmail.com"),
		sender.user, dest, []byte(msg))

	if errors != nil {

		fmt.Printf("smtp error: %s", errors)
		return
	}

	fmt.Println("Mail sent successfully!")
}
