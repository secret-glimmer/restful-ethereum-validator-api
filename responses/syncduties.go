package responses

import (
	"github.com/gofiber/fiber/v2"
)

type SyncDuties struct {
	Keys []string `json:"keys"`
}

func ResponseSyncDuties(ctx *fiber.Ctx, statusCOde int, keys []string) error {
	return ctx.Status(statusCOde).JSON(
		SyncDuties{
			Keys: keys,
		},
	)
}
