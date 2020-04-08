package repository

import (
	"api-dashboard/models"
	"api-dashboard/passenger"
	"context"

	"github.com/jinzhu/gorm"
)

type passengerRepository struct {
	passengers *gorm.DB
}

//NewpassengerRepository creates repository
func NewPassengerRepository(Conn *gorm.DB) passenger.Repository {
	passengers := Conn.Model(&models.Passenger{})
	return &passengerRepository{passengers: passengers}
}
func (pr *passengerRepository) GetOrCreate(ctx context.Context, email string, name string) (*models.Passenger, error) {
	pg := models.Passenger{}
	pg.Name = name
	pg.Email = email
	if err := pr.passengers.Where("email = ?", email).FirstOrCreate(&pg).Error; err != nil {
		return nil, err
	}
	return &pg, nil
}
