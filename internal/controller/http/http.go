package http

import (
	"log"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	router *echo.Echo
}

func New() *Controller {
	log.Println("init controller")
	ctrl := &Controller{
		router: echo.New(),
	}
	return ctrl
}

func (ctrl *Controller) Start() error {
	log.Println("starting http server")
	return ctrl.router.Start("0.0.0.0:8080")
}
