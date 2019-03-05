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

//http://localhost:8080/books/id/1
func (c *BookController) GetIdBy(id int64) (results models.Book) {
	return c.Service.GetByID(id)
}

//http://localhost:8080/books/name/osman
func (c *BookController) GetNameBy(name string) (results models.Book) {
	return c.Service.GetByName(name)
}
