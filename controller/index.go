package controller

import (
	"github.com/electroma/ad-manager/logic"
	"github.com/kataras/iris"
)

type IndexData struct {
	Users map[string]logic.UserInfo
	Message string
}

func GetIndex(ctx *iris.Context) {

	if users, err := GetUsers(); err != nil {
		ctx.MustRender("index.html", IndexData{map[string]logic.UserInfo{}, err.Error()})
	} else {
		ctx.MustRender("index.html", IndexData{users, ""})
	}
}

func GetUsers() (map[string]logic.UserInfo, error) {
	return logic.GetUsers()
}