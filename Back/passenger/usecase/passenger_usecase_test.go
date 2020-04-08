package usecase

import (
	"api-dashboard/models"
	"api-dashboard/passenger/repository"
	_ "api-dashboard/passenger/repository"
	"context"
	"testing"
	"time"
)

func TestPassengerUsecase_GetOrCreate(t *testing.T) {
	mockDB := []models.Passenger{}
	mockDB = append(mockDB, models.Passenger{1, "Laila Weil", "lailaw@lagash.com", nil, time.Now(), time.Now().UTC()})

	mockRepo := repository.NewPassengerRepositoryMock(&mockDB)

	timeoutContext := time.Duration(15 * time.Second)

	useCase := NewPassengerUsecase(mockRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	oldPassenger, errOldPassenger := useCase.GetOrCreate(ctx, "lailaw@lagash.com", "Laila Weil")
	newPassenger, errNewPassenger := useCase.GetOrCreate(ctx, "gonzagr@lagash.com", "Gonzalo Greco")

	if oldPassenger == nil {
		t.Fatal("Expecting passenger", oldPassenger)
	}
	if errOldPassenger != nil {
		t.Error("Expecting error to be nil", errOldPassenger)
	}
	if oldPassenger.ID != 1 {
		t.Error("Expecting ID 1", oldPassenger.ID)
	}
	if oldPassenger.Email != "lailaw@lagash.com" {
		t.Error("Expecting lailaw@lagash.com", oldPassenger.Email)
	}
	if oldPassenger.DeletedAt != nil {
		t.Error("Expecting nil", oldPassenger.DeletedAt)
	}

	if errNewPassenger != nil {
		t.Error("Expecting error to be nil", errNewPassenger)
	}
	if newPassenger == nil {
		t.Fatal("Expecting passenger", newPassenger)
	}
	if newPassenger.ID == 1 {
		t.Error("Expecting not 1", newPassenger.ID)
	}
	if newPassenger.Email != "gonzagr@lagash.com" {
		t.Error("Expecting not 1", newPassenger.Email)
	}
	if newPassenger.Name != "Gonzalo Greco" {
		t.Error("Expecting not 1", newPassenger.Name)
	}

}
