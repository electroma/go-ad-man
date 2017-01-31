package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/go-template/html"
	"github.com/electroma/ad-manager/controller"
)

func main() {

	// naive auth middleware
	iris.UseFunc(func(ctx *iris.Context) {
		if ctx.Path() == "/login" {
			ctx.Next()
		}

		if ctx.Session().Get(controller.USER_VAR) == nil {
			ctx.Redirect("/login")
		}
		ctx.Next()

	})

	// web layer configuration
	iris.OptionSessionsCookie("adman_session")
	iris.UseTemplate(html.New()).Directory("./views", ".html")

	// development version
	iris.Config.IsDevelopment = true
	iris.StaticCacheDuration = 0



	iris.Get("/login", controller.GetLogin)
	iris.Post("/login", controller.PostLogin)

	iris.Get("/", controller.GetIndex)

	iris.Get("/add", controller.GetAdd)
	iris.Post("/add", controller.PostAdd)

	iris.Listen(":6111")
}