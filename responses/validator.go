package responses

import (
	"github.com/gofiber/fiber/v2"
)

type Validators struct {
	Keys []string `json:"keys"`
}

func ResponseValidators(ctx *fiber.Ctx, statusCOde int, keys []string) error {
	return ctx.Status(statusCOde).JSON(
		Validators{
			Keys: keys,
		},
	)
}
