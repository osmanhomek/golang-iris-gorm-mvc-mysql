package main

import (
	"github.com/kataras/iris"
)

func homeHandler(ctx iris.Context) {
	ctx.ViewData("Title", "Hi Page")
	ctx.ViewData("Name", "iris")
	ctx.View("home.html")
}
