package client

import (
	"github.com/agambondan/web-go-blog-grpc-rest/app/lib"
	"github.com/gofiber/fiber/v2"
)

// APIInfoGet func
// @Summary show info response
// @Description show info response
// @Accept  application/json
// @Produce  application/json
// @Success 200 {object} map[string]interface{} "success"
// @Failure 400 {object} lib.Response "bad request"
// @Failure 404 {object} lib.Response "not found"
// @Failure 409 {object} lib.Response "conflict"
// @Failure 500 {object} lib.Response "internal error"
// @Router /info.json [get]
// @Tags Index
func APIInfoGet(c *fiber.Ctx) error {
	info := fiber.Map{
		"id":           "app_id",
		"version":      "v1.0.0",
		"name":         "app name ..",
		"description":  "app description ..",
		"dependencies": fiber.Map{},
		"agent_id":     c.Get("X-Agent-ID"),
		"user_id":      c.Get("X-User-ID"),
	}

	return lib.OK(c, info)
}
