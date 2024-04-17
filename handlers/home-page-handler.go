package handlers

import (
	"github.com/jl-calda/geth-starter/utils/render"
	"github.com/jl-calda/geth-starter/views/pages"
	"github.com/labstack/echo/v4"
)

func HomePageHandler(ctx echo.Context) error {
	return render.Render(ctx, pages.Homepage())
}
