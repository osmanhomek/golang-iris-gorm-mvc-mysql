package controllers

import (
	"gomvc/models"
	"gomvc/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type BookController struct {
	Ctx     iris.Context
	Service services.BookService
	Session *sessions.Session
}

var registerBookStaticView = mvc.View{
	Name: "book_home.html",
	Data: iris.Map{"Title": "Book Home Page"},
}

//http://localhost:8080/books
func (c *BookController) Get() (results []models.Book) {
	return c.Service.GetAll()
}

//http://localhost:8080/books/42
func (c *BookController) GetBy(id int64) (results models.Book) {
	return c.Service.GetByID(id)
}

//http://localhost:8080/books/profile/followers/osman
func (c *BookController) GetProfileFollowersBy(name string) (results models.Book) {
	return c.Service.GetByName(name)
}

//http://localhost:8080/books/profiles/osman/41
func (c *BookController) GetProfilesBy(name string, id int64) (results models.Book) {
	return c.Service.GetByName(name)
}
