package responses

import "github.com/gofiber/fiber/v2"

type Error struct {
	Error string `json:"error"`
}

func ErrorResponse(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(Error{
		Error: message,
	})
}
