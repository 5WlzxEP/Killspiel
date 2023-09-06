package ResultCollector

import (
	"Killspiel/pkg/UserCollector"
	"github.com/gofiber/fiber/v2"
	"math"
	"net/http"
)

type postRequest struct {
	States `json:"state"`
	Result float64 `json:"result"`
}

func post(ctx *fiber.Ctx) error {
	req := postRequest{Result: math.NaN()}

	err := ctx.BodyParser(&req)
	if err != nil {
		return err
	}

	if req.States != State.States {
		return ctx.Status(http.StatusBadRequest).SendString("wrong state")
	}

	switch State.States {
	case running:
		if math.IsNaN(req.Result) {
			return ctx.Status(http.StatusBadRequest).SendString("result is not valid")
		}
		resultChan <- req.Result
		return ctx.SendStatus(http.StatusNoContent)
	case none:
		if beginCancelFunc != nil {
			beginCancelFunc()
			return ctx.SendStatus(http.StatusNoContent)
		}
		return ctx.SendStatus(http.StatusFailedDependency)
	case begin:
		UserCollector.EndCollect()
		return ctx.SendStatus(http.StatusNoContent)
	case result:
		return ctx.Status(http.StatusBadRequest).SendString("wrong state")
	}

	return nil
}
