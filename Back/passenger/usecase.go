package passenger

import (
	"api-dashboard/models"
	"context"
)

//Usecase interfaces
type Usecase interface {
	GetOrCreate(context.Context, string, string) (*models.Passenger, error)
}
