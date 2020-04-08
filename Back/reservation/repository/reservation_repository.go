package repository

import (
	"api-dashboard/helpers"
	"api-dashboard/models"
	"api-dashboard/reservation"
	"context"
	"time"

	"github.com/jinzhu/gorm"
)

type reservationRepository struct {
	reservations *gorm.DB
}

//NewreservationRepository creates repository
func NewReservationRepository(Conn *gorm.DB) reservation.Repository {
	reservations := Conn.Model(&models.Reservation{})

	return &reservationRepository{reservations: reservations}
}

func (rr *reservationRepository) CheckIn(ctx context.Context, r *models.Reservation) (*models.Reservation, error) {
	if err := rr.reservations.First(r).Update("status", helpers.OnBoard).Error; err != nil {
		return nil, err
	}

	return r, nil
}

func (rr *reservationRepository) Create(ctx context.Context, reserv *models.Reservation) (*models.Reservation, error) {
	if err := rr.reservations.Create(&reserv).Error; err != nil {
		return nil, err
	}
	return reserv, nil
}

//Search reservation by id
func (rr *reservationRepository) GetById(ctx context.Context, id string) (*models.Reservation, error) {
	result := &models.Reservation{}
	if err := rr.reservations.Preload("Passenger").Preload("Travel").Where("ID = ?", id).First(result).Error; err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (rr *reservationRepository) Delete(ctx context.Context, id int) (int, error) {
	result := rr.reservations.Where("id = ?", id).Update("status", helpers.Cancelled).Delete(&models.Reservation{})
	if result.Error != nil {
		return 0, result.Error
	}
	return id, nil
}

//Lists reservations by passenger id
func (pr *reservationRepository) GetListReservations(ctx context.Context, passengerid string) (*[]models.Reservation, error) {
	aux := &[]models.Reservation{}
	result := &[]models.Reservation{}
	now := time.Now().UTC().Add(-time.Minute * time.Duration(helpers.TimeForCheckIn))
	if err := pr.reservations.Preload("Travel").Where("passenger_id = ?", passengerid).Find(aux).Error; err != nil {
		return nil, err
	} else {
		for _, r := range *aux {
			if r.Travel != nil {
				if r.Travel.Departure.After(now) {
					*result = append(*result, r)
				}
			}
		}
		return result, nil
	}

}
