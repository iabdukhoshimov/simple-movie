package api

import (
	"movie/api/handlres"
	"movie/config"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
)

func SetUpAPI(r *gin.Engine, h handlres.Handler, cfg config.Config) {

	r.POST("/todo", h.CreateMovie)

}
