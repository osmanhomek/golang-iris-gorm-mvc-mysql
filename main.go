package main

import (
	"gomvc/controllers"
	"gomvc/models"
	"gomvc/repos"
	"gomvc/services"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	//masterpage
	tmpl := iris.HTML("./templates", ".html").Layout("masterpage.html").Reload(true)
	app.RegisterView(tmpl)

	//static files
	app.StaticWeb("/static", "./static")

	//routes
	app.Get("/", homeHandler)

	// **** BOOKS (MySQL)
	dbhost := os.Getenv("MYSQL_ADDON_HOST")
	dbname := os.Getenv("MYSQL_ADDON_DB")
	dbuser := os.Getenv("MYSQL_ADDON_USER")
	dbpassword := os.Getenv("MYSQL_ADDON_PASSWORD")
	dbport := os.Getenv("MYSQL_ADDON_PORT")
	db, err := gorm.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		app.Logger().Fatalf("error while loading the tables: %v", err)
		return
	}
	//for migrate
	db.AutoMigrate(&models.Book{})

	bookRepo := repos.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	books := mvc.New(app.Party("/books"))
	books.Register(bookService)
	books.Handle(new(controllers.BookController))

	//error
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("error.html")
	})

	//start
	app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
