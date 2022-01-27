package controllers

import (
	"go-mongo/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)


var (
	seq_vote  = 1
	seq_comment = 1
)

func CreateVote(c echo.Context) error {
	vote := &models.Vote{
		Id: seq_vote,
	}
	if err := c.Bind(vote); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(vote); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	session := connect()
	defer session.Close()
	if err := session.DB(database).C("vote").Insert(vote); err != nil {
		return c.JSON(http.StatusInternalServerError,"There is an internal issue.")
	}
	seq_vote++
	return c.NoContent(http.StatusNoContent)
}

func CreateComment(c echo.Context) error{
	comment := &models.Comment{
		Id: seq_comment,
	}
	if err := c.Bind(comment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(comment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	comment.CreatedAt = time.Now()

	session := connect()
	defer session.Close()
	if err := session.DB(database).C("comment").Insert(comment); err != nil {
		return c.JSON(http.StatusInternalServerError,"There is an internal issue.")
	}
	seq_comment++
	return c.NoContent(http.StatusNoContent)
}