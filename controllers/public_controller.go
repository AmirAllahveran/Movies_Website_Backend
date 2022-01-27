package controllers

import (
	"go-mongo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

type(
	MoviesRes struct{
		Movies  []models.Movie  `json:"movies"`
	}

	CommentRes struct{
		Id 		int 	`json:"id"`
		Author  string	`json:"author"`
		Body 	string 	`json:"body"`
	}

	CommentsRes struct{
		Movie	string 	`json:"movie"`
		Comments []CommentRes `json:"comments"`
	}

	MovieInfoRes struct{
		Id       	int 			`json:"id"`
		Name     	string          `json:"name,omitempty" validate:"required"`
		Description string          `json:"description,omitempty" validate:"required"`
		Rating    	float64       	`json:"rating,omitempty" validate:"max=1,min=0"`
	}

)

func GetMovies(c echo.Context) error{

	result := new(MoviesRes)
	session := connect()
	defer session.Close()
	if err := session.DB(database).C("movie").Find(nil).All(&result); err != nil{
		return c.JSON(http.StatusInternalServerError,"There is an internal issue.")
	}

	return c.JSON(http.StatusOK,result)

}


func GetComments(c echo.Context) error{
	name := c.QueryParam("name")
	session := connect()
	defer session.Close()
	movie := new(models.Movie)
	if err := session.DB(database).C("movie").Find(bson.M{"name": &name}).One(&movie); err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var comments []models.Comment
	if err := session.DB(database).C("comment").Find(bson.M{"MovieID": &movie.Id}).All(&comments); err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	result := new(CommentsRes)
	result.Movie = movie.Name
	for i := 0; i < len(comments); i++ {
		var c CommentRes
		c.Id = comments[i].Id
		c.Author = comments[i].User.Username
		c.Body = comments[i].Comment
		result.Comments = append(result.Comments, c) 
	}
	return c.JSON(http.StatusOK,result)
}


func GetMovieInfo(c echo.Context) error{
	id, _ := strconv.Atoi(c.Param("id"))
	session := connect()
	defer session.Close()
	movie := new(models.Movie)
	if err := session.DB(database).C("movie").Find(bson.M{"id": &id}).One(&movie); err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	result := new(MovieInfoRes)
	result.Id = movie.Id
	result.Name = movie.Name
	result.Description = movie.Description
	result.Rating = movie.Rating
	return c.JSON(http.StatusOK,result)
}