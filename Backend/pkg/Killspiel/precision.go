package Killspiel

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type precision struct {
	Precision float64 `json:"precision"`
}

func get(ctx *fiber.Ctx) error {
	return ctx.JSON(precision{Precision: conf.Precision})
}

func post(ctx *fiber.Ctx) error {
	var pre precision
	err := ctx.BodyParser(&pre)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if pre.Precision < 0.01 {
		return ctx.Status(http.StatusBadRequest).SendString("precision must be greater than 0.01")
	}

	conf.Precision = pre.Precision
	err = conf.Save()
	if err != nil {
		log.Printf("Error occurred saving config after updating precision: %v\n", err)
	}

	return ctx.SendStatus(http.StatusNoContent)
}
