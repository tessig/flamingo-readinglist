package fake

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"go.opencensus.io/trace"

	"github.com/tessig/flamingo-readinglist/src/app/domain"
)

type (
	// Articles is a fake repository
	Articles struct{}
)

var _ domain.ArticleRepository = new(Articles)

// All Articles
func (a *Articles) All(ctx context.Context) []domain.Article {
	ctx, span := trace.StartSpan(ctx, "app/fake/articles/all")
	defer span.End()

	result := make([]domain.Article, 8)
	for i := 0; i < 8; i++ {
		result[i], _ = a.Get(ctx, strconv.Itoa(i))
	}

	return result
}

// Get returns a fake article with the given ID and never fails
func (a *Articles) Get(ctx context.Context, id string) (domain.Article, error) {
	ctx, span := trace.StartSpan(ctx, "app/fake/articles/get")
	defer span.End()

	return domain.Article{
		ID:     id,
		Title:  fmt.Sprintf("Article %s", id),
		Author: fmt.Sprintf("Guy Incognito %s", id),
		Link:   url.URL{},
		Date:   time.Now(),
		Tags: []string{
			"fake",
			fmt.Sprintf("tag-%s", id),
		},
	}, nil
}
