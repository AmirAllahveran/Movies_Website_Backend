package routes

import (
	"go-mongo/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	r := e.Group("/user")
	r.POST("/vote", controllers.CreateVote)
	r.POST("/comment", controllers.CreateComment)
}