package controller

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"
	"go.opencensus.io/trace"

	"github.com/tessig/flamingo-readinglist/src/app/domain"
)

type (
	// HomeController displays the home page
	HomeController struct {
		responder *web.Responder
		listRepo  domain.ReadingListRepository
	}
)

// Inject dependencies
func (h *HomeController) Inject(
	responder *web.Responder,
	listRepo domain.ReadingListRepository,
) *HomeController {
	h.responder = responder
	h.listRepo = listRepo

	return h
}

// Home shows the start page
func (h *HomeController) Home(ctx context.Context, _ *web.Request) web.Result {
	ctx, span := trace.StartSpan(ctx, "app/controller/home")
	defer span.End()

	list, err := h.listRepo.Load(ctx, "main")
	if err != nil {
		list = &domain.ReadingList{ID: "main"}
		err = h.listRepo.Save(ctx, list)
		if err != nil {
			span.SetStatus(trace.Status{
				Code:    trace.StatusCodeInternal,
				Message: err.Error(),
			})
			return h.responder.ServerError(err)
		}
	}

	return h.responder.Render("index", list)
}
