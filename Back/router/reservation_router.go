package router

import (
	"api-dashboard/reservation"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

//reservationHandler returns handler
type reservationHandler struct {
	reservation.Usecase
}

// @Summary Make Check-In of user for the selected travel.
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param id path int true "Travel"
// @Success 200 {object} models.Reservation
// @Router /reservation/status/{id} [put]
// @Security ApiKeyAuth
func (p *reservationHandler) CheckIn(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	v, err := p.Usecase.CheckIn(c, uint(id))
	if err != nil {
		if err.Error() == "No es posible hacer el check in en este horario" {
			c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}

// @Summary Create reservation on the given travel.
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param travelId path int true "Travel"
// @Success 200 {object} models.Reservation
// @Router /reservation/{travelId} [post]
// @Security ApiKeyAuth
func (p *reservationHandler) Create(c *gin.Context) {
	email := c.Value("userEmail").(string)
	name := c.Value("userName").(string)
	travelId := c.Param("travelId")

	v, err := p.Usecase.Create(c, travelId, email, name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}

// @Summary Get my reservations list.
// @Tags reservations
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Reservation
// @Router /reservation/myreservations [get]
// @Security ApiKeyAuth
func (p *reservationHandler) GetListReservations(c *gin.Context) {
	email := c.Value("userEmail").(string)
	name := c.Value("userName").(string)
	v, err := p.Usecase.GetListReservations(c, email, name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}

// @Summary Delete reservation
// @Tags reservations
// @Accept  json
// @Produce  json
// @Param id path int true "Id reservation"
// @Success 200 integer int
// @Router /reservation/{id} [delete]
// @Security ApiKeyAuth
func (p *reservationHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	v, err := p.Usecase.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}
