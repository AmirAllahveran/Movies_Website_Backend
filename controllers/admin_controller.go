package controllers

import (
	"go-mongo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)


var (
	seq_movie  = 1
)

func CreateMovie(c echo.Context) error {
	movie := &models.Movie{
		Id: seq_movie,
	}

	if err := c.Bind(movie); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(movie); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	session := connect()
	defer session.Close()
	if err := session.DB(database).C("movie").Insert(movie); err != nil {
		return c.JSON(http.StatusInternalServerError,"There is an internal issue.")
	}
	seq_movie++
	return c.NoContent(http.StatusNoContent)
}

func UpdateMovie(c echo.Context) error{
	c_movie := new(models.Movie)
	if err := c.Bind(c_movie); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, _ := strconv.Atoi(c.Param("id"))
	
	session := connect()
	defer session.Close()
	b_movie := new(models.Movie)
	if err := session.DB(database).C("movie").Find(bson.M{"id": &id}).One(&b_movie); err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	b_movie.Name = c_movie.Name
	b_movie.Description = c_movie.Description
	b_movie.Rating = c_movie.Rating
	
	if err := session.DB(database).C("movie").Update(bson.M{"id": &id}, b_movie); err!= nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func RemoveMovie(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	session := connect()
	defer session.Close()
	if err := session.DB(database).C("movie").Remove(bson.M{"id": &id}); err!= nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func UpdateComment(c echo.Context) error{
	c_comment := new(models.Comment)
	if err := c.Bind(c_comment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, _ := strconv.Atoi(c.Param("id"))
	
	session := connect()
	defer session.Close()
	b_comment := new(models.Comment)
	if err := session.DB(database).C("comment").Find(bson.M{"id": &id}).One(&b_comment); err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	b_comment.Approved = c_comment.Approved
	
	if err := session.DB(database).C("movie").Update(bson.M{"id": &id}, b_comment); err!= nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func RemoveComment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	session := connect()
	defer session.Close()
	if err := session.DB(database).C("comment").Remove(bson.M{"id": &id}); err!= nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

