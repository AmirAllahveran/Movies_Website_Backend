package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

// database name are statically declared
const database = "test"
// DB connection
func connect() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("session err:", err)
		os.Exit(1)
	}
	return session
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}