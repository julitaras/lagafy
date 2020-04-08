package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ResponseHandler error response
type ResponseHandler struct {
	Message interface{} `json:"message"`
}

//NewStatushandler retrieves usecase
func NewStatushandler(g *gin.Engine) {
	handler := &ResponseHandler{}

	g.GET("/ping", gin.Logger(), gin.Recovery(), handler.Ping)
}

//Ping handler
func (p *ResponseHandler) Ping(c *gin.Context) {

	c.JSON(http.StatusOK, &ResponseHandler{
		Message: "Pong!",
	})
	return
}
