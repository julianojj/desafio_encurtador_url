package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/desafio_encurtador_url/internal/usecases"
)

type GetOriginalURLController struct {
	GetOriginalURL *usecases.GetOriginalURL
}

func NewGetOriginalURLController(getOriginalURL *usecases.GetOriginalURL) *GetOriginalURLController {
	return &GetOriginalURLController{
		GetOriginalURL: getOriginalURL,
	}
}

func (m *GetOriginalURLController) Handle(c *gin.Context) {
	shortURL := c.Param("id")
	output, err := m.GetOriginalURL.Execute(shortURL)
	if err == nil {
		c.JSON(http.StatusOK, output)
		return
	}
	if err.Error() == "short not found" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"status":  http.StatusNotFound,
		})
		return
	}
	if err.Error() == "expired short" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
			"status":  http.StatusUnprocessableEntity,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
		"status":  http.StatusInternalServerError,
	})
}
