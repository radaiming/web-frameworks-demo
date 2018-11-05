package controllers

import (
	"github.com/astaxie/beego"
	"github.com/radaiming/web-frameworks-demo/core"
	"github.com/radaiming/web-frameworks-demo/core/views"
)

type ReposController struct {
	beego.Controller
}

func (c ReposController) CreateRepoTag() {
	var (
		userName = c.Ctx.Input.GetData("user_name").(string)
		req      views.CreateTagRequest
	)
	err := c.ParseForm(&req)
	if err != nil {
		c.StopRun()
		return
	}

	resp, err := core.CreateTag(userName, c.Ctx.Input.Param(":owner"), c.Ctx.Input.Param(":repo"), &req)
	if err != nil {
		c.StopRun()
		return
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
