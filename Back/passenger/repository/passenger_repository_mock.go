package repository

import (
	"api-dashboard/models"
	"api-dashboard/passenger"
	"context"
	"time"
)

type passengerRepositoryMock struct {
	passengers *[]models.Passenger
}

//NewpassengerRepository creates repository
func NewPassengerRepositoryMock(ps *[]models.Passenger) passenger.Repository {
	passengers := ps
	return &passengerRepositoryMock{passengers: passengers}
}

func (pr *passengerRepositoryMock) GetOrCreate(ctx context.Context, email string, name string) (*models.Passenger, error) {
	pg := models.Passenger{}
	pg.Name = name
	pg.Email = email
	pg.ID = GetLastId(*pr.passengers)
	pg.CreatedAt = time.Now().UTC()
	pg.UpdatedAt = time.Now().UTC()

	for _, p := range *pr.passengers {
		if p.Email == email {
			return &p, nil
		}
	}

	*pr.passengers = append(*pr.passengers, pg)
	return &pg, nil
}

func GetLastId(ps []models.Passenger) uint {
	var lastID uint = 1
	for _, p := range ps {
		if p.ID > lastID {
			lastID = p.ID
		}
	}
	return lastID + 1
}
