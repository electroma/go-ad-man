package controller

import (
	"github.com/electroma/go-ad-man/logic"
	"github.com/kataras/iris"
	"strings"
	"errors"
)

type AddData struct {
	Name        string
	DisplayName string
	Password    string
	Enabled     bool
	Message     string
}

func GetAdd(ctx *iris.Context) {
	ctx.MustRender("add.html", AddData{Enabled: true})
}

func PostAdd(ctx *iris.Context) {
	addData := AddData{}
	ctx.ReadForm(&addData)

	if err := CreateUser(&addData); err != nil {
		addData.Message = err.Error()
		ctx.MustRender("add.html", addData)
	} else {
		ctx.Redirect("/")
	}
}

func CreateUser(user *AddData) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Password = strings.TrimSpace(user.Password)
	if len(user.Name) < 1 {
		return errors.New("Name is not provided")
	}
	if len(user.Password) < 1 {
		return errors.New("Password is not provided")
	}
	return logic.CreateUser(logic.UserInfo{Name: user.Name, DisplayName: user.DisplayName, Password: user.Password, Enabled: user.Enabled})
}
