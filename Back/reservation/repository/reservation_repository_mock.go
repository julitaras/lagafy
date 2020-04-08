package repository

import (
	"api-dashboard/helpers"
	"api-dashboard/models"
	"api-dashboard/reservation"
	"context"
	"errors"
	"strconv"
	"time"
)

type reservationRepositoryMock struct {
	reservations *[]models.Reservation
}

//NewpassengerRepository creates repository
func NewReservationRepositoryMock(ps *[]models.Reservation) reservation.Repository {
	reservations := ps
	return &reservationRepositoryMock{reservations: reservations}
}

func (rr *reservationRepositoryMock) GetById(ctx context.Context, id string) (*models.Reservation, error) {
	ID, _ := strconv.Atoi(id)
	for _, p := range *rr.reservations {
		if int(p.ID) == ID {
			return &p, nil
		}
	}

	return nil, errors.New("No se encontro la reserva")
}

func (rr *reservationRepositoryMock) CheckIn(ctx context.Context, r *models.Reservation) (*models.Reservation, error) {

	for _, p := range *rr.reservations {
		if p.ID == r.ID {
			p.Status = helpers.OnBoard
			return &p, nil
		}
	}

	return nil, errors.New("no existe la reserva")
}

func GetNextId(res []models.Reservation) uint {
	var lastID uint = 0
	for _, r := range res {
		if r.ID > lastID {
			lastID = r.ID
		}
	}
	return lastID + 1
}

func (rr *reservationRepositoryMock) Create(ctx context.Context, reserv *models.Reservation) (*models.Reservation, error) {
	reserv.ID = GetNextId(*rr.reservations)
	*rr.reservations = append(*rr.reservations, *reserv)
	return reserv, nil
}

func (rr *reservationRepositoryMock) Delete(ctx context.Context, id int) (int, error) {
	newReservations := &[]models.Reservation{}
	deletedId := 0

	for _, p := range *rr.reservations {
		if int(p.ID) != id {
			*newReservations = append(*newReservations, p)
		} else {
			deletedId = int(p.ID)
		}
	}

	rr.reservations = newReservations

	if deletedId == 0 {
		return deletedId, errors.New("No se pudo encontrar la reserva")
	}

	return deletedId, nil
}

func (rr *reservationRepositoryMock) GetListReservations(ctx context.Context, passengerid string) (*[]models.Reservation, error) {
	result := &[]models.Reservation{}
	now := time.Now().UTC().Add(-time.Minute * time.Duration(helpers.TimeForCheckIn))
	idPassenger, _ := strconv.Atoi(passengerid)
	for _, r := range *rr.reservations {
		if int(r.PassengerID) == idPassenger {
			if r.Travel.Departure.After(now) {
				*result = append(*result, r)
			}
		}
	}
	return result, nil
}