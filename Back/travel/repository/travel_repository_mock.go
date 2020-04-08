package repository

import (
	"api-dashboard/models"
	"api-dashboard/travel"
	"context"
	"errors"
	"strconv"
	"time"
)

type travelRepositoryMock struct {
	travels *[]models.Travel
}

//NewTravelRepositoryMock creates repository
func NewTravelRepositoryMock(t *[]models.Travel) travel.Repository {
	travels := t
	return &travelRepositoryMock{travels: travels}
}

func (tr *travelRepositoryMock) Insert(ctx context.Context, t *models.Travel) (*models.Travel, error) {
	newTravel := models.Travel{}
	newTravel.ID = t.ID
	newTravel.HasWifi = t.HasWifi
	newTravel.Capacity = t.Capacity
	newTravel.Driver = t.Driver
	newTravel.Departure = t.Departure
	newTravel.Arrival = t.Arrival
	newTravel.Origin = t.Origin
	newTravel.Destination = t.Destination
	newTravel.Status = t.Status
	newTravel.DeletedAt = t.DeletedAt
	newTravel.CreatedAt = t.CreatedAt
	newTravel.UpdatedAt = t.UpdatedAt
	newTravel.Template = t.Template
	newTravel.DepartureAddress = t.DepartureAddress
	newTravel.ArrivalAddress = t.ArrivalAddress

	*tr.travels = append(*tr.travels, newTravel)
	return &newTravel, nil
}

func (tr *travelRepositoryMock) GetById(ctx context.Context, id string) (*models.Travel, error) {
	idTravel, _ := strconv.Atoi(id)

	for _, p := range *tr.travels {
		if p.ID == uint(idTravel) {
			return &p, nil
		}
	}

	return nil, errors.New("No hay ningún viaje con ese Id.")
}

func (tr *travelRepositoryMock) GetCurrentsTravels(ctx context.Context) (*[]models.Travel, error) {
	now := time.Now().UTC()
	result := &[]models.Travel{}
	for _, t := range *tr.travels {
		if t.Departure.After(now) && t.Template != true {

			*result = append(*result, t)
		}
	}

	return result, nil
}

func (tr *travelRepositoryMock) GetTemplates(ctx context.Context) (*[]models.Travel, error) {
	result := &[]models.Travel{}
	for _, t := range *tr.travels {
		if t.Template == true {

			*result = append(*result, t)
		}
	}

	return result, nil
}

func (tr *travelRepositoryMock) Delete(ctx context.Context, id int) (int, error) {
	travels := &[]models.Travel{}
	deletedId := 0
	for _, t := range *tr.travels {
		if int(t.ID) != id {
			*travels = append(*travels, t)
		} else {
			deletedId = int(t.ID)
		}
	}
	tr.travels = travels

	if deletedId == 0 {
		return deletedId, errors.New("No se pudo encontrar el viaje")
	}
	return deletedId, nil
}

func (tr *travelRepositoryMock) UpdateTravel(ctx context.Context, updated *models.Travel) (*models.Travel, error) {
	for _, t := range *tr.travels {
		if t.ID == updated.ID {
			t.HasWifi = updated.HasWifi
			t.Capacity = updated.Capacity
			t.Driver = updated.Driver
			t.Departure = updated.Departure
			t.Arrival = updated.Arrival
			t.Origin = updated.Origin
			t.Destination = updated.Destination
			t.Status = updated.Status
			t.DeletedAt = updated.DeletedAt
			t.CreatedAt = updated.CreatedAt
			t.UpdatedAt = updated.UpdatedAt
			t.Template = updated.Template
			t.DepartureAddress = updated.DepartureAddress
			t.ArrivalAddress = updated.ArrivalAddress

			return &t, nil
		}
	}
	return nil, errors.New("Id not found")
}

func (tr *travelRepositoryMock) GetTravelInfo(ctx context.Context, start time.Time, end time.Time) (*[]models.Travel, error) {
	result := &[]models.Travel{}

	for _, t := range *tr.travels {

		if isBetween(start.UTC(), end.UTC(), t.Departure) && t.Template != true {

			*result = append(*result, t)
		}
	}

	return result, nil
}

func isBetween(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func (tr *travelRepositoryMock) Notify(ctx context.Context) (*[]models.Travel, error) {
	result := &[]models.Travel{}
	to := time.Date(2020, 4, 12, 20, 0, 0, 0, time.UTC)
	from := time.Date(2020, 4, 12, 15, 0, 0, 0, time.UTC)
	for _, t := range *tr.travels {
		if t.Departure.Before(to) && t.Departure.After(from) {
			*result = append(*result, t)
		}
	}
	return result, nil
}
