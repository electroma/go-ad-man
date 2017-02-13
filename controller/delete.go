package controller

import (
	"github.com/electroma/go-ad-man/logic"
	"github.com/kataras/iris"
	log "github.com/Sirupsen/logrus"
)

func GetDelete(ctx *iris.Context) {
	n := ctx.Param("name")
	if err := logic.DeleteUser(n); err != nil {
		// todo: error reporting
		log.Errorln("Failed to delete user ", n, ": ", err.Error())
	}
	ctx.Redirect("/")
}
