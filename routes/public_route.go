package routes

import (
	"go-mongo/controllers"
	"github.com/labstack/echo/v4"
)

func PublicRoute(e *echo.Echo) {
	e.GET("/movies", controllers.GetMovies)
	e.GET("/comments", controllers.GetComments)
	e.GET("/movie/:id",controllers.GetMovieInfo)
}