package routes

import (
	s "validator-api/server"
	"validator-api/server/handlers"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"
)

func ConfigureRoutes(server *s.Server) {
	server.App.Get("/docs/*", swagger.WrapHandler)

	apiV1 := server.App.Group("/api/v1")

	groupBlockReward := apiV1.Group("/blockreward")
	configureGroupBlockReward(server, groupBlockReward)

	groupSyncDuties := apiV1.Group("/syncduties")
	configureGroupSyncDuties(server, groupSyncDuties)
}

func configureGroupBlockReward(server *s.Server, router fiber.Router) {
	handler := handlers.NewHandlerBlockReward(server)

	router.Get("/:slot", handler.GetBlockReward)
}

func configureGroupSyncDuties(server *s.Server, router fiber.Router) {
	handler := handlers.NewHandlerSyncDuties(server)

	router.Get("/:slot", handler.GetSyncDuties)
}
