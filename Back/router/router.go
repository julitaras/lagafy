package router

import (
	"api-dashboard/middleware"
	"api-dashboard/reservation"
	"net/http"

	"api-dashboard/travel"
	"time"

	_passengerRepo "api-dashboard/passenger/repository"

	_travelRepo "api-dashboard/travel/repository"
	_travelUsecase "api-dashboard/travel/usecase"

	_reservationRepo "api-dashboard/reservation/repository"
	_reservationUsecase "api-dashboard/reservation/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//SetRoutes
func SetRoutes(r *gin.Engine, client *gorm.DB) {
	passengerRepo := _passengerRepo.NewPassengerRepository(client)
	travelRepo := _travelRepo.NewTravelRepository(client)
	reservationRepo := _reservationRepo.NewReservationRepository(client)

	timeoutContext := time.Duration(15 * time.Second)

	tuc := _travelUsecase.NewTravelUsecase(travelRepo, timeoutContext)
	ruc := _reservationUsecase.NewReservationUsecase(reservationRepo, passengerRepo, travelRepo, timeoutContext)
	tucWebjob := _travelUsecase.NewTravelUsecase(travelRepo, timeoutContext)

	group := r.Group("/api")
	group.Use(middleware.Authorize())

	// SET HANDLERS
	NewTravelHandler(group, tuc)
	NewReservationHandler(group, ruc)
	NewStatusHandler(group)

	webjobGroup := r.Group("/apigroup")
	webjobGroup.Use(middleware.JwtAuth())

	// SET HANDLERS
	NewTravelHandlerWebjob(webjobGroup, tucWebjob)

}

//ResponseError error response
type ResponseError struct {
	Message string `json:"message"`
}

//NewTravelHandlerWebjob retrieves usecase
func NewTravelHandlerWebjob(g *gin.RouterGroup, uc travel.Usecase) {
	travelHandler := &travelHandler{
		uc,
	}

	g.POST("/travel", travelHandler.Insert)
	g.GET("/travels/templates", travelHandler.GetTemplates)
}

//NewTravelHandler retrieves usecase
func NewTravelHandler(g *gin.RouterGroup, uc travel.Usecase) {
	travelHandler := &travelHandler{
		uc,
	}

	g.POST("/travel", travelHandler.Insert)
	g.GET("/travel/:id", travelHandler.GetByID)
	g.GET("/travels/current", travelHandler.GetCurrentsTravels)
	g.GET("/travels/templates", travelHandler.GetTemplates)
	g.GET("/travels/information", travelHandler.GetTravelInfo)
	g.DELETE("/travel/:id", travelHandler.Delete)
	g.PUT("/travel/:id", travelHandler.UpdateTravel)
	g.GET("/travels/notifications", travelHandler.Notify)
}

//NewReservationHandler retrieves usecase
func NewReservationHandler(g *gin.RouterGroup, uc reservation.Usecase) {
	reservationHandler := &reservationHandler{
		uc,
	}
	g.POST("/reservation/:travelId", reservationHandler.Create)
	g.PUT("/reservation/status/:id", reservationHandler.CheckIn)
	g.DELETE("/reservation/:id", reservationHandler.Delete)
	g.GET("/reservation/myreservations", reservationHandler.GetListReservations)
}

//NewStatusHandler todo
func NewStatusHandler(g *gin.RouterGroup) {

	g.GET("/checkClaims", func(c *gin.Context) {
		email := c.Value("userEmail").(string)
		name := c.Value("userName").(string)
		userData := "Name: " + name + ", Email: " + email
		c.JSON(http.StatusOK, &ResponseHandler{
			Message: userData,
		})
		return
	})

	g.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, &ResponseHandler{
			Message: "Pong!",
		})
		return

	})
}
