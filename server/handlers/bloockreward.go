package handlers

import (
	"strconv"
	"validator-api/models"
	f "validator-api/pkgs/functions"
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
	if f.IsSlotInFuture(slot) {
		return responses.ErrorResponse(ctx,
			fiber.StatusBadRequest,
			"Requested slot is in the future.",
		)
	}

	response := models.ResponseBlock{}

	err := services.BlockByTimeStamp(slot, h.Server.Config.QuckNode.Http, &response)

	if err != nil {
		return responses.ErrorResponse(ctx,
			fiber.StatusNotFound,
			err.Error(),
		)
	}

	block := response.Result

	blockNumer := f.IntFromHex(block.Number)
	staticBlockReward := services.StaticBlockReward(blockNumer)
	transactionFee := services.TransactionFee(block, h.Server.Config.QuckNode.Http)
	burntFee := services.BurntFee(blockNumer, &block)

	result := staticBlockReward.Add(staticBlockReward, transactionFee.Sub(transactionFee, burntFee))

	reward, _ := result.Float64()

	return responses.ResponseBlockReward(ctx, fiber.StatusOK, models.BlockReward{
		Status: "",
		Reward: reward,
	})
}
