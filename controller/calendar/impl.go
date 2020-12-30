package calendar

import (
	"net/http"

	"github.com/acm-uiuc/core/context"
	"github.com/acm-uiuc/core/service"
)

type CalendarController struct {
	svc *service.Service
}

func New(svc *service.Service) *CalendarController {
	return &CalendarController{
		svc: svc,
	}
}

func (controller *CalendarController) GetEvents(ctx *context.Context) error {
	events, err := controller.svc.Calendar.GetEvents()
	if err != nil {
		return ctx.JSONError(
      http.StatusBadRequest,
			"Failed Events List",
			"could not retrieve events list",
			err,
		)
	}

	return ctx.JSON(http.StatusOK, events)
}
