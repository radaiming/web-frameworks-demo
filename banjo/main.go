package main

import (
	"encoding/json"
	"github.com/nsheremet/banjo"
	"github.com/radaiming/web-frameworks-demo/core"
	"github.com/radaiming/web-frameworks-demo/core/authorize"
	"net/http"
	"net/url"
)

func getMyFollowers(c *banjo.Context) {
	u, err := url.Parse(c.Request.URL)
	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return
	}
	userName, err := authorize.TokenToUserName(u.Query().Get("access_token"))
	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		return
	}

	followers, err := core.GetFollowers(userName)
	if err != nil {
		c.Response.Status = http.StatusInternalServerError
		return
	} else {
		c.Response.Status = http.StatusOK
		respJSON, err := json.MarshalIndent(followers, "  ", "  ")
		if err != nil {
			c.Response.Status = http.StatusInternalServerError
		} else {
			c.Response.Status = http.StatusOK
			c.Response.Headers = map[string]string{"Content-Type": "application/json"}
			c.Response.Body = string(respJSON)
		}
	}
}

func main() {
	// banjo doesn't support middleware and path param
	app := banjo.Create(banjo.DefaultConfig())
	app.Get("/user/followers", getMyFollowers)
	app.Run()
}
