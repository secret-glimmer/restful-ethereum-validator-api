package responses

import (
	"validator-api/models"

	"github.com/gofiber/fiber/v2"
)

type BlockReward struct {
	Status string  `json:"status"`
	Reward float64 `json:"rewards"`
}

func ResponseBlockReward(ctx *fiber.Ctx, statusCOde int, reward models.BlockReward) error {
	return ctx.Status(statusCOde).JSON(
		BlockReward{
			Status: reward.Status,
			Reward: reward.Reward,
		},
	)
}
