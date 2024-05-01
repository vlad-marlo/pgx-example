package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vlad-marlo/pgx-example/internal/model"
)

func (ctrl *Controller) HandleCreate(ctx echo.Context) error {
	var req *model.TodoCreateRequest
	if err := ctx.Bind(req); err != nil {
		return err
	}
	return nil
}
