package gateways

import (
	"encoding/json"
	"traffy-mock-crud/domain/entities"

	"github.com/gofiber/fiber/v2"
)

func (h *HTTPGateway) GetAllReports(ctx *fiber.Ctx) error {

	params := ctx.Queries()
	district := params["district"]
	status := params["status"]
	reportID := params["report_id"]

	data, err := h.ReportSevice.GetAllReports(district, status, reportID)

	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h *HTTPGateway) CreateReport(ctx *fiber.Ctx) error {

	jsonData := ctx.FormValue("data")

	var report entities.ReportUserModel
	if err := json.Unmarshal([]byte(jsonData), &report); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	file, err := ctx.FormFile("img")
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "cannot upload image"})
	}

	imgFile, err := file.Open()
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "cannot upload image"})
	}
	defer imgFile.Close()

	if err := h.ReportSevice.InsertNewReport(report, imgFile); err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})

}

func (h *HTTPGateway) DeleteReport(ctx *fiber.Ctx) error {

	params := ctx.Queries()
	reportID := params["report_id"]

	if err := h.ReportSevice.DeleteReport(reportID); err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h *HTTPGateway) AcceptReport(ctx *fiber.Ctx) error {

	params := ctx.Queries()
	reportID := params["report_id"]

	var data entities.ReportOrganizeModel
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	if err := h.ReportSevice.AcceptReport(reportID, data); err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})

}

func (h *HTTPGateway) FinishReport(ctx *fiber.Ctx) error {

	params := ctx.Queries()
	reportID := params["report_id"]

	solvedDetail := ctx.FormValue("solved_detail")

	file, err := ctx.FormFile("img")
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "cannot upload image"})
	}

	imgFile, err := file.Open()
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "cannot upload image"})
	}
	defer imgFile.Close()

	if err := h.ReportSevice.FinishReport(reportID, solvedDetail, imgFile); err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h *HTTPGateway) CommentReport(ctx *fiber.Ctx) error {

	params := ctx.Queries()
	reportID := params["report_id"]

	var data entities.ReportUserCommentModel
	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	if err := h.ReportSevice.CommentReport(reportID, data); err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})

}

func (h *HTTPGateway) AddReaction(ctx *fiber.Ctx) error {

	params := ctx.Queries()
	reportID := params["report_id"]

	var reaction entities.ReactionReport
	if err := ctx.BodyParser(&reaction); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	if err := h.ReportSevice.AddReaction(reportID, reaction); err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}

func (h *HTTPGateway) RemoveReaction(ctx *fiber.Ctx) error {

	params := ctx.Queries()
	reportID := params["report_id"]

	var reaction entities.ReactionReport
	if err := ctx.BodyParser(&reaction); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(entities.ResponseMessage{Message: "invalid json body"})
	}

	if err := h.ReportSevice.RemoveReaction(reportID, reaction); err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success"})
}
