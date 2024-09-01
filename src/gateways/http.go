package gateways

import (
	service "traffy-mock-crud/src/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	ReportSevice service.IReportsService
}

func NewHTTPGateway(app *fiber.App, reports service.IReportsService) {
	gateway := &HTTPGateway{
		ReportSevice: reports,
	}

	GatewayUsers(*gateway, app)
}
