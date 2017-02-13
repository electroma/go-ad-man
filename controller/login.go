package controller

import (
	"github.com/kataras/iris"
	"github.com/electroma/go-ad-man/logic"
)

const USER_VAR = "user"

type LoginData struct {
	Name    string
	Pass    string
	Message string
}

func GetLogin(ctx *iris.Context) {
	ctx.Render("login.html", LoginData{})
}

func PostLogin(ctx *iris.Context) {
	loginData := LoginData{}
	err := ctx.ReadForm(&loginData)

	if err != nil {
		loginData.Message = "Unexpected error: " + err.Error()
		ctx.MustRender("login.html", loginData)
	} else if loginErr := login(loginData.Name, loginData.Pass); loginErr != nil {
		loginData.Message = "Failed to login: " + loginErr.Error()
		ctx.MustRender("login.html", loginData)
	} else {
		ctx.Session().Set(USER_VAR, loginData.Name)
		ctx.Session().Set("pass", loginData.Pass)
		ctx.Redirect("/")
	}
}

func login(user, pass string) error {
	b := "CN=Users,DC=corp,DC=riglet,DC=io"
	return logic.LoginToRemoteAd("CN="+user+"," + b, pass)
}
