package reservation

import (
	"api-dashboard/models"
	"context"
)

//Usecase interfaces
type Usecase interface {
	CheckIn(context.Context, uint) (*models.Reservation, error)
	Create(context.Context, string, string, string) (*models.Reservation, error)
	GetById(context.Context, string) (*models.Reservation, error)
	Delete(context.Context, int) (int, error)
	GetListReservations(context.Context, string, string) (*[]models.Reservation, error)
}
