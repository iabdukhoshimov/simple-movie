package handlres

import (
	"movie/models"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateMovie(ctx *gin.Context) {

	var body models.CreateMovie

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.storage.Movie().Create(ctx, body)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": resp})
}
