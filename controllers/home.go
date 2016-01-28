package controllers

import (
	"github.com/astaxie/beego"
	"github.com/canghai908/beeblog/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}
