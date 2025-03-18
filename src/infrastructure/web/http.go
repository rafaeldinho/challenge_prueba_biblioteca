package web

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ServerInstance() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	return e
}
