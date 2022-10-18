package user

import (
	"github.com/agambondan/web-go-blog-grpc-rest/app/lib"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) FindAll(ctx *fiber.Ctx) error {

	return lib.OK(ctx)
}
