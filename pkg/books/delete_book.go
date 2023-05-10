package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olikhosherstov/examples_gin_postgresql_crud/pkg/common/models"
)

func (h handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	c.Status(http.StatusOK)
}
