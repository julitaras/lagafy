package reservation

import (
	"api-dashboard/models"
	"context"
)

//Repository interfaces
type Repository interface {
	CheckIn(context.Context, *models.Reservation) (*models.Reservation, error)
	Create(context.Context, *models.Reservation) (*models.Reservation, error)
	GetById(context.Context, string) (*models.Reservation, error)
	Delete(context.Context, int) (int, error)
	GetListReservations(context.Context, string) (*[]models.Reservation, error)
}
