package core

import "github.com/radaiming/web-frameworks-demo/core/views"

func CreateTag(userName, owner, repo string, req *views.CreateTagRequest) (*views.CreateTagResponse, error) {
	return &views.CreateTagResponse{
		Message: req.Message,
	}, nil
}
