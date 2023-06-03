package routes

import (
	"prakerja/eventmanagement/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/events", controllers.GetEventController)
	e.POST("/events", controllers.InsertEventController)
	return e
}
