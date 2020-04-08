package usecase

import (
	_ "api-dashboard/helpers"
	"api-dashboard/models"
	"api-dashboard/passenger"

	"context"
	"time"
)

type passengerUsecase struct {
	pr passenger.Repository
	t  time.Duration
}

//NewPassengerUsecase returns passenger usecas
func NewPassengerUsecase(pr passenger.Repository, t time.Duration) passenger.Usecase {
	return &passengerUsecase{
		pr: pr,
		t:  t,
	}
}
func (uc *passengerUsecase) GetOrCreate(ctx context.Context, email string, name string) (*models.Passenger, error) {
	res, err := uc.pr.GetOrCreate(ctx, email, name)
	if err != nil {
		return nil, err
	}
	return res, nil
}
