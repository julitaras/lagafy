package usecase

import (
	"api-dashboard/helpers"
	"api-dashboard/models"
	repository3 "api-dashboard/passenger/repository"
	"api-dashboard/reservation/repository"
	_ "api-dashboard/reservation/repository"
	repository2 "api-dashboard/travel/repository"
	"context"
	"testing"
	"time"
)

func TestReservationUsecase_CheckIn(t *testing.T) {
	mockDB := []models.Reservation{}
	mockDBTravel := []models.Travel{}
	mockDBPassenger := []models.Passenger{}

	travel := &models.Travel{Departure: time.Now().UTC().Add(-time.Minute * time.Duration(1)), ID: 1}
	travel2 := &models.Travel{Departure: time.Now().UTC().Add(time.Minute * time.Duration(45)), ID: 2}
	travel3 := &models.Travel{Departure: time.Now().UTC().Add(-time.Minute * time.Duration(45)), ID: 3}

	mockDB = append(mockDB, models.Reservation{ID: 1, TravelID: 1, Travel: travel, Status: helpers.Confirmed, PassengerID: 1, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})
	mockDB = append(mockDB, models.Reservation{ID: 2, TravelID: 1, Travel: travel, Status: helpers.Pending, PassengerID: 2, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})
	mockDB = append(mockDB, models.Reservation{ID: 3, TravelID: 2, Travel: travel2, Status: helpers.Confirmed, PassengerID: 1, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})
	mockDB = append(mockDB, models.Reservation{ID: 4, TravelID: 3, Travel: travel3, Status: helpers.Confirmed, PassengerID: 1, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})

	mockRepo := repository.NewReservationRepositoryMock(&mockDB)
	mockPassengerRepo := repository3.NewPassengerRepositoryMock(&mockDBPassenger)
	mockTravelRepo := repository2.NewTravelRepositoryMock(&mockDBTravel)
	timeoutContext := time.Duration(15 * time.Second)

	useCase := NewReservationUsecase(mockRepo, mockPassengerRepo, mockTravelRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	existingReservation, errExiting := useCase.CheckIn(ctx, 1)
	pendingReservation, errPending := useCase.CheckIn(ctx, 2)
	lateCheckIn, errLate := useCase.CheckIn(ctx, 3)
	earlyCheckIn, errEarly := useCase.CheckIn(ctx, 4)
	notExistingReservation, errNotExiting := useCase.CheckIn(ctx, 5)

	//existing reservation
	if errExiting != nil {
		t.Error("Expecting error to be nil", errExiting)
	}
	if existingReservation == nil {
		t.Fatal("Expecting reservation not to be nil", existingReservation)
	}
	if existingReservation.Status != helpers.OnBoard {
		t.Error("Expecting status to be "+helpers.OnBoard, existingReservation.Status)
	}
	if existingReservation.ID != 1 {
		t.Error("Expecting id to be 1", existingReservation.ID)
	}

	//pending reservation
	if errPending == nil {
		t.Error("Expecting error not to be nil", errPending)
	}
	if pendingReservation != nil {
		t.Error("Expecting reservation to be nil", pendingReservation.ID)
	}

	//not existing reservation
	if errNotExiting == nil {
		t.Error("Expecting error not to be nil", errNotExiting)
	}
	if notExistingReservation != nil {
		t.Error("Expecting reservation to be nil", notExistingReservation)
	}

	//Late Reservation
	if errLate == nil {
		t.Error("Expecting error not to be nil", errLate)
	}
	if lateCheckIn != nil {
		t.Error("Expecting reservation to be nil", lateCheckIn)
	}

	//Early Reservation
	if errEarly == nil {
		t.Error("Expecting error not to be", errEarly)
	}
	if earlyCheckIn != nil {
		t.Error("Expecting reservation to be nil", earlyCheckIn)
	}
}

func TestReservationUsecase_Create(t *testing.T) {
	mockDB := []models.Reservation{}
	mockDBTravel := []models.Travel{}
	mockDBPassenger := []models.Passenger{}

	reserv := models.Reservation{ID: 1, TravelID: 1, PassengerID: 4, Status: "pending"}

	travel := models.Travel{Departure: time.Now().UTC().Add(time.Minute * time.Duration(45)), ID: 1, Capacity: 2, Reservations: []*models.Reservation{}}
	invalidTravel := models.Travel{Departure: time.Now().UTC().Add(-time.Minute * time.Duration(45)), ID: 2, Capacity: 2, Reservations: []*models.Reservation{&reserv}}
	travelwithReservation := models.Travel{Departure: time.Now().UTC().Add(-time.Minute * time.Duration(45)), ID: 3, Capacity: 2, Reservations: []*models.Reservation{}}
	noCapacityTravel := models.Travel{Departure: time.Now().UTC().Add(-time.Minute * time.Duration(45)), ID: 4, Capacity: 0, Reservations: []*models.Reservation{}}
	mockDBTravel = append(mockDBTravel, travel)
	mockDBTravel = append(mockDBTravel, invalidTravel)
	mockDBTravel = append(mockDBTravel, travelwithReservation)
	mockDBTravel = append(mockDBTravel, noCapacityTravel)

	mockDBPassenger = append(mockDBPassenger, models.Passenger{ID: 1, Name: "Gonzalo", Email: "gonzalogr@lagash.com", CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})

	mockRepo := repository.NewReservationRepositoryMock(&mockDB)
	mockPassengerRepo := repository3.NewPassengerRepositoryMock(&mockDBPassenger)
	mockTravelRepo := repository2.NewTravelRepositoryMock(&mockDBTravel)
	timeoutContext := time.Duration(15 * time.Second)

	useCase := NewReservationUsecase(mockRepo, mockPassengerRepo, mockTravelRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	newReservation, errNewReservation := useCase.Create(ctx, "1", "gonzalogr@lagash.com", "Gonzalo")
	invalidTravelReservation, errInvalidTravel := useCase.Create(ctx, "10", "gonzalogr@lagash.com", "Gonzalo")
	unexistingPassenger, errUnexistingPassenger := useCase.Create(ctx, "1", "gastonb@lagash.com", "Gaston")
	notOnTime, errNotOnTime := useCase.Create(ctx, "2", "lailaw@lagash.com", "Laila")
	alreadyExist, errAlreadyExist := useCase.Create(ctx, "3", "hugof@lagash.com", "Hugo Ali")
	fullCapacity, errFullCapacity := useCase.Create(ctx, "4", "estefanias@lagash.com", "estefania")

	// NEW
	if errNewReservation != nil {
		t.Error("Expecting error to be nil", errNewReservation)
	}
	if newReservation == nil {
		t.Error("Expecting reservation not to be nil")
	}

	// INVALID TRAVEL
	if errInvalidTravel == nil {
		t.Error("Expecting invalid travel error", errInvalidTravel)
	}
	if invalidTravelReservation != nil {
		t.Error("Expecting reservation to be nil")
	}

	// UNEXISTING PASSENGER
	if errUnexistingPassenger != nil {
		t.Error("Expecting error to be nil", errUnexistingPassenger)
	}
	if unexistingPassenger == nil {
		t.Error("Expecting reservation not to be nil")
	}

	// NOT ON TIME
	if errNotOnTime == nil {
		t.Error("Expecting not on time error", errNotOnTime)
	}
	if notOnTime != nil {
		t.Error("Expecting reservation to be nil")
	}

	// EXISTING PREV RESERVATION
	if errAlreadyExist == nil {
		t.Error("Expecting already existing reservation error", errAlreadyExist)
	}
	if alreadyExist != nil {
		t.Error("Expecting reservation to be nil")
	}

	// FULL CAPACITY
	if errFullCapacity == nil {
		t.Error("Expecting full capacity error", errAlreadyExist)
	}
	if fullCapacity != nil {
		t.Error("Expecting reservation to be nil")
	}
}

func TestReservationUsecase_GetListReservations(t *testing.T) {
	mockDB := []models.Reservation{}
	mockDBTravel := []models.Travel{}
	mockDBPassenger := []models.Passenger{}

	passenger := &models.Passenger{ID: 2, Name: "Pablo", Email: "pablo_1@gmail.com"}
	passenger2 := &models.Passenger{ID: 3, Name: "Melisa", Email: "melisa_2@gmail.com"}

	travel := &models.Travel{Departure: time.Now().UTC(), ID: 1}
	travel2 := &models.Travel{Departure: time.Now().UTC().Add(time.Minute * time.Duration(45)), ID: 2}
	travel3 := &models.Travel{Departure: time.Now().UTC().Add(-time.Minute * time.Duration(100)), ID: 3}

	mockDB = append(mockDB, models.Reservation{ID: 1, TravelID: 1, Travel: travel, Status: helpers.Confirmed, PassengerID: 2, Passenger: *passenger, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})
	mockDB = append(mockDB, models.Reservation{ID: 2, TravelID: 3, Travel: travel3, Status: helpers.Pending, PassengerID: 2, Passenger: *passenger, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})
	mockDB = append(mockDB, models.Reservation{ID: 3, TravelID: 2, Travel: travel2, Status: helpers.Confirmed, PassengerID: 2, Passenger: *passenger, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})
	mockDB = append(mockDB, models.Reservation{ID: 4, TravelID: 1, Travel: travel3, Status: helpers.Confirmed, PassengerID: 3, Passenger: *passenger2, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})

	mockRepo := repository.NewReservationRepositoryMock(&mockDB)
	mockPassengerRepo := repository3.NewPassengerRepositoryMock(&mockDBPassenger)
	mockTravelRepo := repository2.NewTravelRepositoryMock(&mockDBTravel)
	timeoutContext := time.Duration(15 * time.Second)

	useCase := NewReservationUsecase(mockRepo, mockPassengerRepo, mockTravelRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	listReservation, errListReservation := useCase.GetListReservations(ctx, "pablo_1@gmail.com", "Pablo")

	if errListReservation != nil {
		t.Error("Expecting error to be nil", errListReservation)
	}
	if len(*listReservation) != 2 {
		t.Error("Expecting count equal 2")
	}

	for _, r := range *listReservation {
		if r.Passenger.Name != "Pablo" {
			t.Error("Wrong passenger")
		}
	}
}

func TestReservationUsecase_Delete(t *testing.T) {
	mockDB := []models.Reservation{}
	mockDBTravel := []models.Travel{}
	mockDBPassenger := []models.Passenger{}

	passenger := &models.Passenger{ID: 2, Name: "Pablo", Email: "pablo_1@gmail.com"}

	travel := &models.Travel{Departure: time.Now().UTC(), ID: 1}

	mockDB = append(mockDB, models.Reservation{ID: 1, TravelID: 1, Travel: travel, Status: helpers.Confirmed, PassengerID: 2, Passenger: *passenger, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()})

	mockRepo := repository.NewReservationRepositoryMock(&mockDB)
	mockPassengerRepo := repository3.NewPassengerRepositoryMock(&mockDBPassenger)
	mockTravelRepo := repository2.NewTravelRepositoryMock(&mockDBTravel)
	timeoutContext := time.Duration(15 * time.Second)

	useCase := NewReservationUsecase(mockRepo, mockPassengerRepo, mockTravelRepo, timeoutContext)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutContext)
	defer cancel()

	deleteWrongID, errDeleteWrongID := useCase.Delete(ctx, 10)
	deleteReservation, errDeleteReservation := useCase.Delete(ctx, 1)
	deleteNoReservation, errDeleteNoReservation := useCase.Delete(ctx, 1)

	// DELETE WRONG ID
	if errDeleteWrongID == nil {
		t.Error("Expecting error not to be nil", errDeleteReservation)
	}
	if deleteWrongID != -0 {
		t.Error("Expecting reservation not to be nil")
	}

	// DELETE CORRECT ID
	if errDeleteReservation != nil {
		t.Error("Expecting error to be nil", errDeleteReservation)
	}
	if deleteReservation != 1 {
		t.Error("Expecting reservation not to be nil")
	}

	// DELETE WITH NO RESERVATIONS
	if errDeleteNoReservation == nil {
		t.Error("Expecting error not to be nil", errDeleteReservation)
	}
	if deleteNoReservation != 0 {
		t.Error("Expecting reservation not to be nil")
	}
}