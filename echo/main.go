package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/radaiming/web-frameworks-demo/core"
	"github.com/radaiming/web-frameworks-demo/core/authorize"
	"github.com/radaiming/web-frameworks-demo/core/views"
	"net/http"
)

func getUserFollowers(c echo.Context) error {
	followers, err := core.GetFollowers(c.Param("username"))
	if err != nil {
		return err
	} else {
		return c.JSONPretty(http.StatusOK, followers, "  ")
	}
}

func getMyFollowers(c echo.Context) error {
	followers, err := core.GetFollowers(c.Get("username").(string))
	if err != nil {
		return err
	} else {
		return c.JSONPretty(http.StatusOK, followers, "  ")
	}
}

func createRepoTag(c echo.Context) error {
	var req views.CreateTagRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	tag, err := core.CreateTag(c.Param("owner"), c.Param("repo"), &req)
	if err != nil {
		return err
	} else {
		return c.JSONPretty(http.StatusOK, tag, "  ")
	}
}

// middleware for authentication
func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userName, err := authorize.TokenToUserName(c.QueryParam("access_token"))
		if err != nil {
			return err
		} else {
			c.Set("user_name", userName)
			return next(c)
		}
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	user := e.Group("/user")
	user.GET("/:username/followers", getUserFollowers)
	user.GET("/followers", getMyFollowers, auth)

	repos := e.Group("/repos/:owner/:repo")
	repos.POST("/git/tags", createRepoTag, auth)

	e.Start("127.0.0.1:4321")
}
