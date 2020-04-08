package travel

import (
	"api-dashboard/models"
	"context"
	"time"
)

//Repository interfaces
type Repository interface {
	Insert(context.Context, *models.Travel) (*models.Travel, error)
	GetById(context.Context, string) (*models.Travel, error)
	GetCurrentsTravels(context.Context) (*[]models.Travel, error)
	GetTemplates(context.Context) (*[]models.Travel, error)
	Delete(context.Context, int) (int, error)
	UpdateTravel(context.Context, *models.Travel) (*models.Travel, error)
	GetTravelInfo(context.Context, time.Time, time.Time) (*[]models.Travel, error)
	Notify(context.Context) (*[]models.Travel, error)
}
