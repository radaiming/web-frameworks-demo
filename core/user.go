package core

import "github.com/radaiming/web-frameworks-demo/core/models"

func GetFollowers(userName string) ([]models.User, error) {
	// get user from database and return
	return []models.User{{
		AvatarURL: "http://xxx...",
	}}, nil
}
