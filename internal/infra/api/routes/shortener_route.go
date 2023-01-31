package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/julianojj/desafio_encurtador_url/internal/infra/api/controllers"
)

type ShortenerRoute struct {
	App                      *gin.Engine
	MakeShortenerController  *controllers.MakeShortenerURLController
	GetOriginalURLController *controllers.GetOriginalURLController
}

func NewShortenerRoute(app *gin.Engine, makeShortenerController *controllers.MakeShortenerURLController, getOriginalURLController *controllers.GetOriginalURLController) *ShortenerRoute {
	return &ShortenerRoute{
		App:                      app,
		MakeShortenerController:  makeShortenerController,
		GetOriginalURLController: getOriginalURLController,
	}
}

func (s *ShortenerRoute) Init() {
	s.App.POST("/cut", s.MakeShortenerController.Handle)
	s.App.GET("/uncut/:id", s.GetOriginalURLController.Handle)
}
