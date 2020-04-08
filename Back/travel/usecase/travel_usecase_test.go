package usecase

import (
	"api-dashboard/helpers"
	"api-dashboard/models"
	"api-dashboard/travel/repository"
	_ "api-dashboard/travel/repository"
	"context"
	"testing"
	"time"
)

func TestTravelUsecase_Insert(t *testing.T) {
	mockDB := []models.Travel{}

	mockDB = append(mockDB, models.Travel{ID: 1, HasWifi: true, Capacity: 15, Driver: "conductor", Departure: time.Now().UTC(), Arrival: time.Now().UTC(), Origin: "origen", Destination: "destino", Status: "arrived", DeletedAt: nil, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), Template: true, DepartureAddress: "Calle 123", ArrivalAddress: "Calle 321"})
	mockDB = append(mockDB, models.Travel{ID: 2, HasWifi: true, Capacity: 15, Driver: "conductor", Departure: time.Now().UTC(), Arrival: time.Now().UTC(), Origin: "origen", Destination: "destino", Status: "arrived", DeletedAt: nil, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), Template: true, DepartureAddress: "Calle 123", ArrivalAddress: "Calle 321"})

	mockRepo := repository.NewTravelRepositoryMock(&mockDB)

	timeoutContext := time.Duration(15 * time.Second)

	useCase := NewTravelUsecase(mockRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	createdTravels, err := useCase.Insert(ctx, &mockDB)

	if createdTravels == nil {
		t.Fatal("Expecting new travel", createdTravels)
	}

	if err != nil {
		t.Error("Expecting Error to be nil", err)
	}

	if len(*createdTravels) != 2 {
		t.Error("Expecing travels to be of length of 2", len(*createdTravels))
	}

	if (*createdTravels)[0].ID != 1 {
		t.Error("Expecting travel ID 1", (*createdTravels)[0].ID)
	}

	if (*createdTravels)[0].HasWifi != true {
		t.Error("Expecting travel HasWifi true", (*createdTravels)[0].HasWifi)
	}

	if (*createdTravels)[0].Driver != "conductor" {
		t.Error("Expecting travel Driver conductor", (*createdTravels)[0].Driver)
	}

	if (*createdTravels)[0].Origin != "origen" {
		t.Error("Expecting travel Origin origen", (*createdTravels)[0].Origin)
	}

	if (*createdTravels)[0].Destination != "destino" {
		t.Error("Expecting travel Destination destino", (*createdTravels)[0].Destination)
	}

	if (*createdTravels)[0].Status != "arrived" {
		t.Error("Expecting travel Status arrived", (*createdTravels)[0].Status)
	}

	if (*createdTravels)[0].DepartureAddress != "Calle 123" {
		t.Error("Expecting travel departure address Calle 123", (*createdTravels)[0].DepartureAddress)
	}

	if (*createdTravels)[0].ArrivalAddress != "Calle 321" {
		t.Error("Expecting travel arrival address Calle 321", (*createdTravels)[0].ArrivalAddress)
	}

	if (*createdTravels)[1].ID != 2 {
		t.Error("Expecting travel ID 2", (*createdTravels)[1].ID)

	}
}

func TestTravelUsecase_GetCurrentsTravels(t *testing.T) {
	mockDB := []models.Travel{}

	travel := models.Travel{Departure: time.Now().UTC().Add(time.Minute * time.Duration(20)), ID: 1, Template: false}
	travel2 := models.Travel{Departure: time.Now().UTC().Add(time.Minute * time.Duration(45)), ID: 2, Template: false}
	travel3 := models.Travel{Departure: time.Now().UTC().Add(-time.Minute * time.Duration(45)), ID: 3, Template: false}
	travel4 := models.Travel{Departure: time.Now().UTC().Add(-time.Minute * time.Duration(45)), ID: 4, Template: true}

	mockDB = append(mockDB, travel)
	mockDB = append(mockDB, travel2)
	mockDB = append(mockDB, travel3)
	mockDB = append(mockDB, travel4)
	mockRepo := repository.NewTravelRepositoryMock(&mockDB)

	timeoutContext := time.Duration(15 * time.Second)

	useCase := NewTravelUsecase(mockRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	currentsTravels, err := useCase.GetCurrentsTravels(ctx)

	if err != nil {
		t.Error("Expecting error to be nil", err)
	}

	if len(*currentsTravels) != 2 {
		t.Error("Expecting two travels")
	}
}

func TestTravelUsecase_GetById(t *testing.T) {
	mockDB := []models.Travel{}

	travel := models.Travel{Departure: time.Now().UTC().Add(time.Minute * time.Duration(20)), ID: 1, Reservations: []*models.Reservation{}}
	travel2 := models.Travel{Departure: time.Now().UTC().Add(time.Minute * time.Duration(45)), ID: 2, Reservations: []*models.Reservation{}}

	mockDB = append(mockDB, travel)
	mockDB = append(mockDB, travel2)
	mockRepo := repository.NewTravelRepositoryMock(&mockDB)

	timeoutContext := time.Duration(15 * time.Second)

	useCase := NewTravelUsecase(mockRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	foundTravel, foundError := useCase.GetById(ctx, "1")
	notFoundTravel, notFoundError := useCase.GetById(ctx, "3")

	//found travel
	if foundTravel == nil {
		t.Fatal("Expecting travel not to be nil", foundTravel)
	}

	if foundError != nil {
		t.Error("Expecting error to be nil", foundError)
	}

	if foundTravel.Reservations == nil {
		t.Error("Expecting reservations to be an array", foundTravel.Reservations)
	}

	//not found travel
	if notFoundTravel != nil {
		t.Fatal("Expecting travel not to be nil", notFoundTravel)
	}

	if notFoundError == nil {
		t.Error("Expecting error to be nil", notFoundError)
	}

}

func TestTravelUsecase_GetTemplates(t *testing.T) {
	mockDB := []models.Travel{}

	travel := models.Travel{ID: 1, Template: true}
	travel2 := models.Travel{ID: 2, Template: true}
	travel3 := models.Travel{ID: 3, Template: false}

	mockDB = append(mockDB, travel)
	mockDB = append(mockDB, travel2)
	mockDB = append(mockDB, travel3)

	mockRepo := repository.NewTravelRepositoryMock(&mockDB)
	timeoutContext := time.Duration(15 * time.Second)

	useCase := NewTravelUsecase(mockRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	templates, err := useCase.GetTemplates(ctx)

	if err != nil {
		t.Error("Expecting Error to be nil", err)
	}

	if len(*templates) != 2 {
		t.Error("Expecting two templates", len(*templates))
	}

	if (*templates)[0].Template != true {
		t.Error("Expecting template be true", (*templates)[0].Template)
	}

	if (*templates)[1].Template != true {
		t.Error("Expecting template be true", (*templates)[1].Template)
	}
}

func TestTravelUsecase_Update(t *testing.T) {
	mockDB := []models.Travel{}

	travel := models.Travel{ID: 1, HasWifi: true, Capacity: 15, Driver: "conductor", Departure: time.Now().UTC(), Arrival: time.Now().UTC(), Origin: "origen", Destination: "destino", Status: "arrived", DeletedAt: nil, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), Template: true, DepartureAddress: "Calle 123", ArrivalAddress: "Calle 321"}

	mockDB = append(mockDB, travel)
	mockRepo := repository.NewTravelRepositoryMock(&mockDB)

	timeoutContext := time.Duration(15 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	useCase := NewTravelUsecase(mockRepo, timeoutContext)

	newTravel := models.Travel{ID: 1, HasWifi: true, Capacity: 20, Driver: "Gonzalo", Departure: time.Now().UTC(), Arrival: time.Now().UTC(), Origin: "Lagash PP", Destination: "Meli DOT", Status: "arrived", DeletedAt: nil, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), Template: true, DepartureAddress: "Avenida 456", ArrivalAddress: "Avenida 456"}
	invalidIdTravel := models.Travel{ID: 2, HasWifi: true, Capacity: 20, Driver: "Gonzalo", Departure: time.Now().UTC(), Arrival: time.Now().UTC(), Origin: "Lagash PP", Destination: "Meli DOT", Status: "arrived", DeletedAt: nil, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), Template: true, DepartureAddress: "Avenida 456", ArrivalAddress: "Avenida 456"}

	updatededTravel, err := useCase.UpdateTravel(ctx, &newTravel)
	invalidId, IdErr := useCase.UpdateTravel(ctx, &invalidIdTravel)

	if updatededTravel == nil {
		t.Fatal("Expecting new travel", updatededTravel)
	}

	if err != nil {
		t.Error("Expecting Error to be nil", err)
	}

	if updatededTravel.Driver != "Gonzalo" {
		t.Error("Expecting Gonzalo", updatededTravel.Driver)
	}

	if updatededTravel.Origin != "Lagash PP" {
		t.Error("Expecting Lagash PP", updatededTravel.Origin)
	}

	if updatededTravel.Destination != "Meli DOT" {
		t.Error("Expecting Meli DOT", updatededTravel.Destination)
	}

	if updatededTravel.Capacity != 20 {
		t.Error("Expecting 20", updatededTravel.Capacity)
	}

	if updatededTravel.DepartureAddress != "Avenida 456" {
		t.Error("Expecting 20", updatededTravel.DepartureAddress)
	}

	if updatededTravel.ArrivalAddress != "Avenida 456" {
		t.Error("Expecting Avenida 456", updatededTravel.ArrivalAddress)
	}

	if invalidId != nil {
		t.Fatal("Expecting nil", updatededTravel)
	}

	if IdErr == nil {
		t.Error("Expecting Error", err)
	}
}

func TestTravelUsecase_Delete(t *testing.T) {
	mockDB := []models.Travel{}

	travel := models.Travel{ID: 1, CreatedAt: time.Now()}
	travel2 := models.Travel{ID: 2, CreatedAt: time.Now()}
	travel3 := models.Travel{ID: 3, CreatedAt: time.Now()}

	mockDB = append(mockDB, travel)
	mockDB = append(mockDB, travel2)
	mockDB = append(mockDB, travel3)

	mockRepo := repository.NewTravelRepositoryMock(&mockDB)
	timeoutContext := time.Duration(15 * time.Second)
	useCase := NewTravelUsecase(mockRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	deleted1, err := useCase.Delete(ctx, 1)
	deleted2, err2 := useCase.Delete(ctx, 2)
	notDeleted, err3 := useCase.Delete(ctx, 10)

	//Check travel found
	if deleted1 == 0 {
		t.Error("Expecting travel to be deleted", err)
	}

	if err != nil {
		t.Error("Expecting Error to be nil", err)
	}

	//Check travel found
	if deleted2 == 0 {
		t.Error("Expecting travel2 to be deleted", err2)
	}

	if err2 != nil {
		t.Error("Expecting Error to be nil", err2)
	}

	//Check travel not found
	if notDeleted != 0 {
		t.Error("Expecting not to find this travel", err3)
	}

	if err3 == nil {
		t.Error("Expecting Error to be 'No se pudo encontrar el viaje'", err3)
	}

}

func TestTravelUsecase_GetTravelInfo(t *testing.T) {

	mockDBReservation1 := []*models.Reservation{}
	mockDBReservation2 := []*models.Reservation{}
	mockDBReservation3 := []*models.Reservation{}

	reservation1 := models.Reservation{ID: 1, TravelID: 1, Status: helpers.Confirmed}
	reservation2 := models.Reservation{ID: 2, TravelID: 1, Status: helpers.Pending}
	reservation3 := models.Reservation{ID: 3, TravelID: 2, Status: helpers.OnBoard}
	reservation4 := models.Reservation{ID: 4, TravelID: 2, Status: helpers.Cancelled}
	reservation5 := models.Reservation{ID: 5, TravelID: 3, Status: helpers.Confirmed}
	reservation6 := models.Reservation{ID: 6, TravelID: 3, Status: helpers.OnBoard}
	reservation7 := models.Reservation{ID: 7, TravelID: 3, Status: helpers.OnBoard}

	mockDBReservation1 = append(mockDBReservation1, &reservation1)
	mockDBReservation1 = append(mockDBReservation1, &reservation2)
	mockDBReservation2 = append(mockDBReservation2, &reservation3)
	mockDBReservation2 = append(mockDBReservation2, &reservation4)
	mockDBReservation3 = append(mockDBReservation3, &reservation5)
	mockDBReservation3 = append(mockDBReservation3, &reservation6)
	mockDBReservation3 = append(mockDBReservation3, &reservation7)

	mockDB := []models.Travel{}
	travel1 := models.Travel{ID: 1, Departure: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), Arrival: time.Date(2020, 1, 10, 1, 0, 0, 0, time.UTC), Reservations: mockDBReservation1, Template: false, Capacity: 15, Origin: "Origen", Destination: "Destino"}
	travel2 := models.Travel{ID: 2, Departure: time.Date(2020, 2, 15, 0, 0, 0, 0, time.UTC), Arrival: time.Date(2020, 2, 15, 1, 0, 0, 0, time.UTC), Reservations: mockDBReservation2, Template: false, Capacity: 15, Origin: "Origen", Destination: "Destino"}
	travel3 := models.Travel{ID: 3, Departure: time.Date(2020, 3, 12, 0, 0, 0, 0, time.UTC), Arrival: time.Date(2020, 3, 12, 1, 0, 0, 0, time.UTC), Reservations: mockDBReservation3, Template: false, Capacity: 15, Origin: "Origen", Destination: "Destino"}
	travel4 := models.Travel{ID: 4, Departure: time.Date(2020, 4, 10, 0, 0, 0, 0, time.UTC), Arrival: time.Date(2020, 4, 10, 1, 0, 0, 0, time.UTC), Reservations: []*models.Reservation{}, Template: false, Capacity: 15, Origin: "Origen", Destination: "Destino"}
	travel5 := models.Travel{ID: 5, Departure: time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC), Arrival: time.Date(2020, 1, 20, 1, 0, 0, 0, time.UTC), Reservations: []*models.Reservation{}, Template: false, Capacity: 15, Origin: "Origen", Destination: "Destino"}
	travel6 := models.Travel{ID: 6, Departure: time.Date(2020, 4, 20, 0, 0, 0, 0, time.UTC), Arrival: time.Date(2020, 4, 20, 1, 0, 0, 0, time.UTC), Reservations: []*models.Reservation{}, Template: true, Capacity: 15, Origin: "Origen", Destination: "Destino"}

	mockDB = append(mockDB, travel1)
	mockDB = append(mockDB, travel2)
	mockDB = append(mockDB, travel3)
	mockDB = append(mockDB, travel4)
	mockDB = append(mockDB, travel5)
	mockDB = append(mockDB, travel6)

	mockRepo := repository.NewTravelRepositoryMock(&mockDB)
	timeoutContext := time.Duration(15 * time.Second)
	useCase := NewTravelUsecase(mockRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	travels, err := useCase.GetTravelInfo(ctx, "2020-01-05", "2020-03-20")

	if err != nil {
		t.Error("Expecting Error to be nil", err)
	}

	if len(*travels) != 4 {
		t.Error("Expecting four travels, result: ", len(*travels))
	}

	if (*travels)[0].ID != 1 {
		t.Error("Expecting ID = 1, result: ", (*travels)[0].ID)
	}

	if (*travels)[1].ID != 2 {
		t.Error("Expecting ID = 2, result: ", (*travels)[1].ID)
	}

	if (*travels)[2].ID != 3 {
		t.Error("Expecting ID = 3, result: ", (*travels)[2].ID)
	}

	if (*travels)[3].ID != 5 {
		t.Error("Expecting ID = 3, result: ", (*travels)[2].ID)
	}

	if (*travels)[0].Departure != time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC) {
		t.Error("Expecting Departure: ", time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), " Result: ", (*travels)[0].Departure)
	}

	if (*travels)[1].Departure != time.Date(2020, 2, 15, 0, 0, 0, 0, time.UTC) {
		t.Error("Expecting Departure: ", time.Date(2020, 2, 15, 0, 0, 0, 0, time.UTC), " Result: ", (*travels)[1].Departure)
	}

	if (*travels)[2].Departure != time.Date(2020, 3, 12, 0, 0, 0, 0, time.UTC) {
		t.Error("Expecting Departure: ", time.Date(2020, 3, 12, 0, 0, 0, 0, time.UTC), " Result: ", (*travels)[2].Departure)
	}

	if (*travels)[3].Departure != time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC) {
		t.Error("Expecting Departure: ", time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC), " Result: ", (*travels)[3].Departure)
	}

	if (*travels)[0].Arrival != time.Date(2020, 1, 10, 1, 0, 0, 0, time.UTC) {
		t.Error("Expecting Arrival: ", time.Date(2020, 1, 10, 1, 0, 0, 0, time.UTC), " Result: ", (*travels)[0].Arrival)
	}

	if (*travels)[1].Arrival != time.Date(2020, 2, 15, 1, 0, 0, 0, time.UTC) {
		t.Error("Expecting Arrival: ", time.Date(2020, 2, 15, 1, 0, 0, 0, time.UTC), " Result: ", (*travels)[1].Arrival)
	}

	if (*travels)[2].Arrival != time.Date(2020, 3, 12, 1, 0, 0, 0, time.UTC) {
		t.Error("Expecting Arrival: ", time.Date(2020, 3, 12, 1, 0, 0, 0, time.UTC), " Result: ", (*travels)[2].Arrival)
	}

	if (*travels)[3].Arrival != time.Date(2020, 1, 20, 1, 0, 0, 0, time.UTC) {
		t.Error("Expecting Arrival: ", time.Date(2020, 1, 20, 1, 0, 0, 0, time.UTC), " Result: ", (*travels)[3].Arrival)
	}

	if (*travels)[0].Pending != 1 {
		t.Error("Expecting Pending total = 1. Result:", (*travels)[0].Pending)
	}

	if (*travels)[0].OnBoard != 0 {
		t.Error("Expecting OnBoard total = 0. Result:", (*travels)[0].OnBoard)
	}

	if (*travels)[0].Cancelled != 0 {
		t.Error("Expecting Cancelled total = 0. Result:", (*travels)[0].Cancelled)
	}

	if (*travels)[0].Confirmed != 1 {
		t.Error("Expecting Confirmed total = 1. Result:", (*travels)[0].Confirmed)
	}

	if (*travels)[1].Pending != 0 {
		t.Error("Expecting Pending total = 0. Result:", (*travels)[1].Pending)
	}

	if (*travels)[1].OnBoard != 1 {
		t.Error("Expecting OnBoard total = 1. Result:", (*travels)[1].OnBoard)
	}

	if (*travels)[1].Cancelled != 1 {
		t.Error("Expecting Cancelled total = 1. Result:", (*travels)[1].Cancelled)
	}

	if (*travels)[1].Confirmed != 0 {
		t.Error("Expecting Confirmed total = 0. Result:", (*travels)[1].Confirmed)
	}

	if (*travels)[2].Pending != 0 {
		t.Error("Expecting Pending total = 0. Result:", (*travels)[2].Pending)
	}

	if (*travels)[2].OnBoard != 2 {
		t.Error("Expecting OnBoard total = 2. Result:", (*travels)[2].OnBoard)
	}

	if (*travels)[2].Cancelled != 0 {
		t.Error("Expecting Cancelled total = 0. Result:", (*travels)[2].Cancelled)
	}

	if (*travels)[2].Confirmed != 1 {
		t.Error("Expecting Confirmed total = 1. Result:", (*travels)[2].Confirmed)
	}

	if (*travels)[3].Pending != 0 {
		t.Error("Expecting Pending total = 0. Result:", (*travels)[3].Pending)
	}

	if (*travels)[3].OnBoard != 0 {
		t.Error("Expecting OnBoard total = 0. Result:", (*travels)[3].OnBoard)
	}

	if (*travels)[3].Cancelled != 0 {
		t.Error("Expecting Cancelled total = 0. Result:", (*travels)[3].Cancelled)
	}

	if (*travels)[3].Confirmed != 0 {
		t.Error("Expecting Confirmed total = 0. Result:", (*travels)[3].Confirmed)
	}

	if (*travels)[0].Capacity != 15 {
		t.Error("Expecting Capacity = 15. Result:", (*travels)[0].Capacity)
	}

	if (*travels)[1].Capacity != 15 {
		t.Error("Expecting Capacity = 15. Result:", (*travels)[1].Capacity)
	}

	if (*travels)[2].Capacity != 15 {
		t.Error("Expecting Capacity = 15. Result:", (*travels)[2].Capacity)
	}

	if (*travels)[3].Capacity != 15 {
		t.Error("Expecting Capacity = 15. Result:", (*travels)[3].Capacity)
	}

	if (*travels)[0].Origin != "Origen" {
		t.Error("Expecting Origin = Origen. Result:", (*travels)[0].Origin)
	}

	if (*travels)[1].Origin != "Origen" {
		t.Error("Expecting Origin = Origen. Result:", (*travels)[1].Origin)
	}

	if (*travels)[2].Origin != "Origen" {
		t.Error("Expecting Origin = Origen. Result:", (*travels)[2].Origin)
	}

	if (*travels)[3].Origin != "Origen" {
		t.Error("Expecting Origin = Origen. Result:", (*travels)[3].Origin)
	}

	if (*travels)[0].Destination != "Destino" {
		t.Error("Expecting Destination = Destino. Result:", (*travels)[0].Destination)
	}

	if (*travels)[1].Destination != "Destino" {
		t.Error("Expecting Destination = Destino. Result:", (*travels)[1].Destination)
	}

	if (*travels)[2].Destination != "Destino" {
		t.Error("Expecting Destination = Destino. Result:", (*travels)[2].Destination)
	}

	if (*travels)[3].Destination != "Destino" {
		t.Error("Expecting Destination = Destino. Result:", (*travels)[3].Destination)
	}
}

func TestTravelUsecase_Notify(t *testing.T) {
	mockDBReservation1 := []*models.Reservation{}
	mockDBReservation2 := []*models.Reservation{}
	mockDBReservation3 := []*models.Reservation{}

	reservation1 := models.Reservation{ID: 1, TravelID: 1, Status: helpers.Confirmed}
	reservation2 := models.Reservation{ID: 2, TravelID: 1, Status: helpers.Pending}
	reservation3 := models.Reservation{ID: 3, TravelID: 1, Status: helpers.OnBoard}
	reservation4 := models.Reservation{ID: 4, TravelID: 1, Status: helpers.Confirmed}
	reservation5 := models.Reservation{ID: 5, TravelID: 1, Status: helpers.Confirmed}
	reservation6 := models.Reservation{ID: 6, TravelID: 1, Status: helpers.OnBoard}
	reservation7 := models.Reservation{ID: 7, TravelID: 1, Status: helpers.OnBoard}
	reservation8 := models.Reservation{ID: 8, TravelID: 3, Status: helpers.Confirmed}
	reservation9 := models.Reservation{ID: 9, TravelID: 3, Status: helpers.Pending}
	reservation10 := models.Reservation{ID: 10, TravelID: 3, Status: helpers.OnBoard}
	reservation11 := models.Reservation{ID: 11, TravelID: 2, Status: helpers.Confirmed}
	reservation22 := models.Reservation{ID: 22, TravelID: 2, Status: helpers.Pending}
	reservation33 := models.Reservation{ID: 33, TravelID: 2, Status: helpers.OnBoard}
	reservation44 := models.Reservation{ID: 44, TravelID: 2, Status: helpers.Confirmed}
	reservation55 := models.Reservation{ID: 55, TravelID: 2, Status: helpers.Confirmed}
	reservation66 := models.Reservation{ID: 66, TravelID: 2, Status: helpers.OnBoard}
	reservation77 := models.Reservation{ID: 77, TravelID: 2, Status: helpers.OnBoard}

	mockDBReservation1 = append(mockDBReservation1, &reservation1)
	mockDBReservation1 = append(mockDBReservation1, &reservation2)
	mockDBReservation1 = append(mockDBReservation1, &reservation3)
	mockDBReservation1 = append(mockDBReservation1, &reservation4)
	mockDBReservation1 = append(mockDBReservation1, &reservation5)
	mockDBReservation1 = append(mockDBReservation1, &reservation6)
	mockDBReservation1 = append(mockDBReservation1, &reservation7)
	mockDBReservation3 = append(mockDBReservation3, &reservation8)
	mockDBReservation3 = append(mockDBReservation3, &reservation9)
	mockDBReservation3 = append(mockDBReservation3, &reservation10)
	mockDBReservation2 = append(mockDBReservation2, &reservation11)
	mockDBReservation2 = append(mockDBReservation2, &reservation22)
	mockDBReservation2 = append(mockDBReservation2, &reservation33)
	mockDBReservation2 = append(mockDBReservation2, &reservation44)
	mockDBReservation2 = append(mockDBReservation2, &reservation55)
	mockDBReservation2 = append(mockDBReservation2, &reservation66)
	mockDBReservation2 = append(mockDBReservation2, &reservation77)

	mockDB := []models.Travel{}
	travel1 := models.Travel{ID: 1, Departure: time.Date(2020, 4, 12, 17, 0, 0, 0, time.UTC), Arrival: time.Date(2020, 4, 12, 18, 0, 0, 0, time.UTC), Reservations: mockDBReservation1, Template: false, Capacity: 15, Origin: "Origen", Destination: "Destino"}
	travel2 := models.Travel{ID: 2, Departure: time.Date(2020, 4, 12, 18, 0, 0, 0, time.UTC), Arrival: time.Date(2020, 4, 12, 19, 0, 0, 0, time.UTC), Reservations: mockDBReservation2, Template: false, Capacity: 15, Origin: "Origen", Destination: "Destino"}
	travel3 := models.Travel{ID: 3, Departure: time.Date(2020, 4, 12, 21, 0, 0, 0, time.UTC), Arrival: time.Date(2020, 4, 12, 22, 0, 0, 0, time.UTC), Reservations: mockDBReservation3, Template: false, Capacity: 15, Origin: "Origen", Destination: "Destino"}

	mockDB = append(mockDB, travel1)
	mockDB = append(mockDB, travel2)
	mockDB = append(mockDB, travel3)

	mockRepo := repository.NewTravelRepositoryMock(&mockDB)
	timeoutContext := time.Duration(15 * time.Second)
	useCase := NewTravelUsecase(mockRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()
	travels, err := useCase.Notify(ctx)
	if err != nil {
		t.Error("Expecting Error to be nil", err)
	}
	if len(*travels) == 0 {
		t.Error("Expecting travels > 0", err)
	}
	reser := &[]models.Reservation{}
	for _, t := range *travels {
		for _, r := range t.Reservations {
			if r.Status != "onboard" {
				*reser = append(*reser, *r)
			}
		}
	}
	if len(*reser) != 8 {
		t.Error("Expecting reservations = 8", err)
	}
}
