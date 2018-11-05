package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/radaiming/web-frameworks-demo/beego/controllers"
	"github.com/radaiming/web-frameworks-demo/core/authorize"
	"net/http"
)

func auth(c *context.Context) {
	var accessToken string
	err := c.Input.Bind(&accessToken, "access_token")
	if err != nil {
		c.Abort(http.StatusUnauthorized, "no access token")
		return
	}

	userName, err := authorize.TokenToUserName(accessToken)
	if err != nil {
		c.Abort(http.StatusInternalServerError, "internal error")
		return
	}
	c.Input.SetData("user_name", userName)
}

// actually the code structure is not that beego style
func main() {
	beego.InsertFilter("/user/followers", beego.BeforeRouter, auth)
	beego.InsertFilter("/repos/?:owner/?:repo/git/tags", beego.BeforeRouter, auth)

	beego.Router("/user/?:username/followers", &controllers.FollowersController{}, "get:GetUserFollowers")
	beego.Router("/user/followers", &controllers.FollowersController{}, "get:GetMyFollowers")
	beego.Router("/repos/?:owner/?:repo/git/tags", &controllers.ReposController{}, "post:CreateRepoTag")
	beego.Run("127.0.0.1:4321")
}
