package repository

import (
	"api-dashboard/helpers"
	"api-dashboard/models"
	"api-dashboard/travel"
	"context"
	"sort"
	"time"

	"github.com/jinzhu/gorm"
)

type travelRepository struct {
	travels *gorm.DB
}

//NewtravelRepository creates repository
func NewTravelRepository(Conn *gorm.DB) travel.Repository {
	travels := Conn.Model(&models.Travel{})
	return &travelRepository{travels: travels}
}

func (tr *travelRepository) Insert(ctx context.Context, t *models.Travel) (*models.Travel, error) {
	err := tr.travels.Create(&t)

	if err.Error != nil {
		return nil, err.Error
	}

	return t, nil
}

//Search travel by id
func (tr *travelRepository) GetById(ctx context.Context, id string) (*models.Travel, error) {
	result := &models.Travel{}
	if err := tr.travels.Where("id = ?", id).Where("template = ?", false).Preload("Reservations").Preload("Reservations.Passenger").Find(result).Error; err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//Lists travels by status
func (tr *travelRepository) GetCurrentsTravels(ctx context.Context) (*[]models.Travel, error) {
	result := &[]models.Travel{}
	now := time.Now().UTC().Add(-time.Minute * time.Duration(helpers.TimeForCheckIn))
	if err := tr.travels.Where("Departure >= ?", now).Where("template = ?", false).Find(result).Error; err != nil {
		return nil, err
	} else {
		sort.Slice(*result, func(i, j int) bool {
			return (*result)[i].Departure.Before((*result)[j].Departure)
		})
		return result, nil
	}
}

func (tr *travelRepository) GetTemplates(ctx context.Context) (*[]models.Travel, error) {
	result := &[]models.Travel{}
	if err := tr.travels.Where("Template = ?", true).Find(result).Error; err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (tr *travelRepository) Delete(ctx context.Context, id int) (int, error) {
	result := tr.travels.Where("id = ?", id).Delete(&models.Travel{})
	if result.Error != nil {
		return 0, result.Error
	}
	return id, nil
}

func (tr *travelRepository) UpdateTravel(ctx context.Context, t *models.Travel) (*models.Travel, error) {
	result := tr.travels.Where("id = ?", t.ID).Update(t)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}

func (tr *travelRepository) GetTravelInfo(ctx context.Context, start time.Time, end time.Time) (*[]models.Travel, error) {
	result := &[]models.Travel{}
	if err := tr.travels.Where("Departure BETWEEN ? AND ?", start.UTC(), end.UTC()).Where("template = ?", false).Preload("Reservations").Find(result).Error; err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (tr *travelRepository) Notify(ctx context.Context) (*[]models.Travel, error) {
	travel := &[]models.Travel{}
	to := time.Now()
	var from time.Time
	if to.Hour() < 18 {
		from = time.Date(to.Year(), to.Month(), to.Day(), 0, 0, 0, 0, time.UTC)
	} else {
		from = time.Date(to.Year(), to.Month(), to.Day(), 15, 0, 0, 0, time.UTC)
	}
	if err := tr.travels.Where("Departure between ? and ?", from, to).Where("template = ?", false).Preload("Reservations").Preload("Reservations.Passenger").Find(travel).Error; err != nil {
		return nil, err
	}
	return travel, nil
}
