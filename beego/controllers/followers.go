package controllers

import (
	"github.com/astaxie/beego"
	"github.com/radaiming/web-frameworks-demo/core"
)

type FollowersController struct {
	beego.Controller
}

func (c FollowersController) GetUserFollowers() {
	userName := c.Ctx.Input.Param("username")
	followers, err := core.GetFollowers(userName)
	if err != nil {
		c.StopRun()
		return
	}

	c.Data["json"] = followers
	c.ServeJSON()
}

func (c FollowersController) GetMyFollowers() {
	userName := c.Ctx.Input.GetData("user_name").(string)
	followers, err := core.GetFollowers(userName)
	if err != nil {
		c.StopRun()
		return
	}

	c.Data["json"] = followers
	c.ServeJSON()
}
