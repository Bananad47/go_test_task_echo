package app

import (
	"github.com/Bananad47/go_test_task_echo/internal/endpoints"
	"github.com/Bananad47/go_test_task_echo/internal/middlewares"
	"github.com/labstack/echo/v4"
)

type App struct {
	echo *echo.Echo
}

const addr = ":80"

var middlewaresSlice = []echo.MiddlewareFunc{
	middlewares.FindAdmin,
}

func CreateApp() *App {
	app := &App{}
	app.echo = echo.New()
	app.echo.Use(middlewaresSlice...)
	app.echo.GET("/", endpoints.TimeHandler)
	return app
}

func (a *App) Start() error {
	return a.echo.Start(addr)
}
