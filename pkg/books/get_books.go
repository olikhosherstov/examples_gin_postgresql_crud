package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olikhosherstov/examples_gin_postgresql_crud/pkg/common/models"
)

func (h handler) GetBooks(c *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &books)
}
