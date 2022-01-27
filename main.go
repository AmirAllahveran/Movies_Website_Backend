package main

import (
	"go-mongo/routes"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.AdminRoute(e)
	routes.UserRoute(e)
	routes.PublicRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}
