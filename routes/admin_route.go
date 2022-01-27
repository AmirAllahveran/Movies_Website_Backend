package routes

import (
	"go-mongo/controllers"
	"github.com/labstack/echo/v4"
)

func AdminRoute(e *echo.Echo) {
	r := e.Group("/admin")
	r.POST("/movie", controllers.CreateMovie)
	r.PUT("/movie/:id", controllers.UpdateMovie)
	r.DELETE("/movie/:Id", controllers.RemoveMovie)
	r.PUT("/comment/:id", controllers.UpdateComment)
	r.DELETE("/comment/:id", controllers.RemoveComment)
}