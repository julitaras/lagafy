package passenger

import (
	"api-dashboard/models"
	"context"
)

//Repository interfaces
type Repository interface {
	GetOrCreate(context.Context, string, string) (*models.Passenger, error)
}
