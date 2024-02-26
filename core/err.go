package core

import (
	"backend/api"

	"github.com/gofiber/fiber/v2"
)

var (
	log = GetLogger()
)

func SendErr(ctx *fiber.Ctx, err error, logErr bool) error {
	if logErr {
		log.Error(err)
	}
	return ctx.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
}
func SendErrWithMsg(ctx *fiber.Ctx, err error, msg string) error {
	if err != nil {
		log.Error(err)
	}
	return ctx.Status(400).JSON(fiber.Map{"success": false, "message": msg})

}

func LogErrWithMsg(ctx *fiber.Ctx, err error, msg string) error {
	if err != nil {
		log.Error(err)
	}
	return ctx.Status(400).JSON(fiber.Map{"success": false, "message": msg})

}
func SendErrString(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(400).JSON(fiber.Map{"success": false, "message": msg})
}

func SendValidError(ctx *fiber.Ctx, errs []*api.ErrorResponse, logErr bool) error {
	if logErr {
		log.Error(errs)
	}
	resp := api.FiberValidErr{Success: false, Errors: errs}
	return ctx.Status(fiber.StatusBadRequest).JSON(resp)

}
