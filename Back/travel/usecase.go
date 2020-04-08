package travel

import (
	"api-dashboard/models"
	"context"
)

//Usecase interfaces
type Usecase interface {
	Insert(context.Context, *[]models.Travel) (*[]models.Travel, error)
	GetById(context.Context, string) (*models.Travel, error)
	GetCurrentsTravels(context.Context) (*[]models.Travel, error)
	GetTemplates(context.Context) (*[]models.Travel, error)
	Delete(context.Context, int) (int, error)
	UpdateTravel(context.Context, *models.Travel) (*models.Travel, error)
	GetTravelInfo(context.Context, string, string) (*[]models.TravelInformation, error)
	Notify(context.Context) (*[]models.Travel, error)
}
