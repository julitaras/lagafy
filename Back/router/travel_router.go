package router

import (
	"api-dashboard/models"
	"api-dashboard/travel"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//travelHandler returns handler
type travelHandler struct {
	travel.Usecase
}

// @Summary Create new travel
// @Tags travels
// @Accept  json
// @Produce  json
// @Param travel body models.Travel true "Create a travel"
// @Success 200 {object} models.Travel
// @Router /travel/ [post]
// @Security ApiKeyAuth
func (p *travelHandler) Insert(c *gin.Context) {
	userMail := c.Request.Header.Get("userMail")
	if userMail == "" {
		userMail = c.Value("userEmail").(string)
	}
	if !valideAdministrator(userMail) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized user",
		})
		return
	}
	nv := &[]models.Travel{}
	err := c.Bind(nv)
	if err != nil {
		return
	}
	v, err := p.Usecase.Insert(c, nv)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}

// @Summary Get travel by ID
// @Tags travels
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Travel
// @Param id path int true "travel ID"
// @Router /travel/{id} [get]
// @Security ApiKeyAuth
func (p *travelHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	v, err := p.Usecase.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}

// @Summary Get currents travels
// @Tags travels
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Travel
// @Router /travels/current [get]
// @Security ApiKeyAuth
func (p *travelHandler) GetCurrentsTravels(c *gin.Context) {

	v, err := p.Usecase.GetCurrentsTravels(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}

// @Summary Get templates for travels
// @Tags travels
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Travel
// @Router /travels/templates [get]
// @Security ApiKeyAuth
func (p *travelHandler) GetTemplates(c *gin.Context) {

	v, err := p.Usecase.GetTemplates(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}

// @Summary Update travel
// @Tags travels
// @Accept  json
// @Produce  json
// @Param id path int true "Id travel"
// @Param travel body models.Travel true "Update a travel"
// @Success 200 {object} models.Travel
// @Router /travel/{id} [put]
// @Security ApiKeyAuth
func (p *travelHandler) UpdateTravel(c *gin.Context) {
	email := c.Value("userEmail").(string)
	if !valideAdministrator(email) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized user",
		})
		return
	}
	t := &models.Travel{}
	err := c.Bind(t)
	if err != nil {
		return
	}
	result, err := p.Usecase.UpdateTravel(c, t)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
	return
}

// @Summary Delete travel
// @Tags travels
// @Accept  json
// @Produce  json
// @Success 200 integer int
// @Param id path int true "travel ID"
// @Router /travel/{id} [delete]
// @Security ApiKeyAuth
func (p *travelHandler) Delete(c *gin.Context) {
	email := c.Value("userEmail").(string)
	if !valideAdministrator(email) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized user",
		})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	v, err := p.Usecase.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}

// @Summary Get information from travels
// @Tags travels
// @Accept  json
// @Produce  json
// @Success 200 {array} models.TravelInformation
// @Param start query string true "Start date. Expample: 2020-02-20"
// @Param end query string true "End date. Expample: 2020-03-10"
// @Router /travels/information [get]
// @Security ApiKeyAuth
func (p *travelHandler) GetTravelInfo(c *gin.Context) {
	email := c.Value("userEmail").(string)
	if !valideAdministrator(email) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized user",
		})
		return
	}
	start := c.Query("start")
	end := c.Query("end")

	v, err := p.Usecase.GetTravelInfo(c, start, end)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}

func (p *travelHandler) Notify(c *gin.Context) {

	v, err := p.Usecase.Notify(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, v)
	return
}

func valideAdministrator(userMail string) bool {
	if userMail == "johannac@lagash.com" || userMail == "vanessam@lagash.com" {
		return true
	}
	return false
}
