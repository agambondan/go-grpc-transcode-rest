package user

import (
	"github.com/agambondan/web-go-blog-grpc-rest/app/repo"
)

type Controller struct {
	userRepository repo.UserRepository
}

func NewUserController(repo repo.UserRepository) *Controller {
	return &Controller{userRepository: repo}
}
