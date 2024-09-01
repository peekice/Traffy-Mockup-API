package gateways

import "github.com/gofiber/fiber/v2"

func GatewayUsers(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api")

	api.Get("/reports", gateway.GetAllReports)
	api.Post("/create_report", gateway.CreateReport)
	api.Delete("/delete_report", gateway.DeleteReport)

	api.Put("/accept_report", gateway.AcceptReport)
	api.Put("/finish_report", gateway.FinishReport)

	api.Put("/comment_report", gateway.CommentReport)

	api.Put("/add_reaction", gateway.AddReaction)
	api.Put("/remove_reaction", gateway.RemoveReaction)
}
