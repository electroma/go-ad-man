package main

import (
	"os"
	"github.com/kataras/iris"
	"github.com/kataras/go-template/html"
	"github.com/electroma/go-ad-man/controller"
	"github.com/electroma/go-ad-man/logic"
	"github.com/kataras/go-errors"
	"github.com/Sirupsen/logrus"
)

func main() {

	// configuration
	logic.DCHost = os.Getenv("DC_URL")
	logic.Base = os.Getenv("BASE_DN")

	logrus.Infof("Running with DN=%s URL=%s", logic.Base, logic.DCHost)

	if err := validateConf(); err != nil {
		panic(err)
	}

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
	iris.UseTemplate(html.New(html.Config{Layout:"layout.html"})).Directory("./views", ".html")

	// development version
	iris.Config.IsDevelopment = true
	iris.StaticCacheDuration = 0

	iris.Get("/login", controller.GetLogin)
	iris.Post("/login", controller.PostLogin)

	iris.Get("/", controller.GetIndex)

	iris.Get("/add", controller.GetAdd)
	iris.Post("/add", controller.PostAdd)
	iris.Get("/delete/:name", controller.GetDelete)

	iris.Listen(":6111")
}
func validateConf() error {
	if len(logic.Base) == 0 {
		return errors.New("Base DN is not provided")
	}
	if len(logic.DCHost) == 0 {
		return errors.New("DC host is not provided")
	}
	return nil
}