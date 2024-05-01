package http

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/vlad-marlo/pgx-example/internal/repository"
)

type Controller struct {
	router *echo.Echo
	repo   repository.Interface
}

func New(repo repository.Interface) *Controller {
	log.Println("init controller")
	ctrl := &Controller{
		router: echo.New(),
		repo:   repo,
	}
	ctrl.configureRoutes()
	return ctrl
}

func (ctrl *Controller) configureRoutes() {
	log.Println("configuring routes")
	ctrl.router.POST("/", ctrl.HandleCreate)
}

func (ctrl *Controller) Start() error {
	log.Println("starting http server")
	return ctrl.router.Start("0.0.0.0:8080")
}
