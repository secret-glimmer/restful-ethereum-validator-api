package handlers

import (
	"strconv"
	"validator-api/pkgs/consts"
	fn "validator-api/pkgs/functions"
	"validator-api/responses"
	s "validator-api/server"
	"validator-api/services"

	"github.com/gofiber/fiber/v2"
)

type HandlerBlockReward struct {
	Server *s.Server
}

func NewHandlerBlockReward(server *s.Server) *HandlerBlockReward {
	return &HandlerBlockReward{
		Server: server,
	}
}

// @Summary Get block reward
// @Description Get block reward by slot
// @Tags blockreward
// @Accept json
// @Produce json
// @Param slot path integer true "Slot"
// @Success 200 {object} responses.BlockReward
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /blockreward/{slot} [get]
func (h *HandlerBlockReward) GetBlockReward(ctx *fiber.Ctx) error {
	slot, _ := strconv.Atoi(ctx.Params("slot"))

	// Check if slot is in future
	if fn.IsSlotInFuture(slot) {
		return responses.ErrorResponse(ctx,
			fiber.StatusBadRequest,
			consts.SlotInFuture,
		)
	}

	reward, isEmpty, isError := services.GetBlockReward(slot, h.Server.Config)

	// Check internal server error
	if isError {
		return responses.ErrorResponse(ctx,
			fiber.StatusInternalServerError,
			consts.InternelServerError,
		)
	}

	// Check if slot does exist
	if isEmpty {
		return responses.ErrorResponse(ctx,
			fiber.StatusNotFound,
			consts.NotFound,
		)
	}

	return responses.ResponseBlockReward(ctx, fiber.StatusOK, reward)
}
