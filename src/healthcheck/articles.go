package healthcheck

import (
	"context"

	"flamingo.me/flamingo/v3/core/healthcheck/domain/healthcheck"

	"github.com/tessig/flamingo-readinglist/src/app/domain"
)

type (
	// Articles checks the article source
	Articles struct {
		repo domain.ArticleRepository
	}
)

var _ healthcheck.Status = new(Articles)

// Inject dependencies
func (a *Articles) Inject(
	repo domain.ArticleRepository,
) *Articles {
	a.repo = repo

	return a
}

// Status of the remote article source
func (a *Articles) Status() (bool, string) {
	articles, err := a.repo.All(context.Background())
	if err != nil {
		return false, err.Error()
	}
	if len(articles) == 0 {
		return false, "No articles found"
	}

	return true, "Articles can be fetched"
}
