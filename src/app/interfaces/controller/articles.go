package controller

import (
	"context"
	"errors"
	"fmt"

	"flamingo.me/flamingo/v3/framework/web"
	"go.opencensus.io/trace"

	"github.com/tessig/flamingo-readinglist/src/app/domain"
)

type (
	// ArticlesController displays the home page
	ArticlesController struct {
		responder   *web.Responder
		articleRepo domain.ArticleRepository
		listRepo    domain.ReadingListRepository
	}
)

// Inject dependencies
func (a *ArticlesController) Inject(
	responder *web.Responder,
	articleRepository domain.ArticleRepository,
	listRepository domain.ReadingListRepository,
) *ArticlesController {
	a.responder = responder
	a.articleRepo = articleRepository
	a.listRepo = listRepository

	return a
}

// List all articles
func (a *ArticlesController) List(ctx context.Context, _ *web.Request) web.Result {
	ctx, span := trace.StartSpan(ctx, "app/controller/articles/list")
	defer span.End()

	articles, err := a.articleRepo.All(ctx)
	if err != nil {
		return a.responder.ServerError(err)
	}

	return a.responder.Render("articles", articles)
}

// AddToList adds an article to the list
func (a *ArticlesController) AddToList(ctx context.Context, r *web.Request) web.Result {
	ctx, span := trace.StartSpan(ctx, "app/controller/articles/list/add")
	defer span.End()

	listID, ok := r.Params["listID"]
	if !ok {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInvalidArgument,
			Message: "listID is missing",
		})
		return a.responder.ServerError(errors.New("listID is missing"))
	}
	list, err := a.listRepo.Load(ctx, listID)
	if err != nil {
		list = &domain.ReadingList{ID: "main"}
	}

	articleID, ok := r.Params["articleID"]
	if !ok {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInvalidArgument,
			Message: "articleID is missing",
		})
		return a.responder.ServerError(errors.New("articleID is missing"))
	}
	article, err := a.articleRepo.Get(ctx, articleID)
	if err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeNotFound,
			Message: fmt.Sprintf("article %q not found", articleID),
		})
		return a.responder.ServerError(err)
	}
	list.AddArticle(ctx, article)

	err = a.listRepo.Save(ctx, list)
	if err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeInternal,
			Message: err.Error(),
		})
		return a.responder.ServerError(err)
	}

	return a.responder.RouteRedirect("/articles", nil)
}
