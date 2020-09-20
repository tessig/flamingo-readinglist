package domain

import (
	"context"
	"net/url"
	"time"
)

type (
	// Article represents a editorial piece to read
	Article struct {
		ID     string
		Title  string
		Author string
		Link   url.URL
		Date   time.Time
		Tags   []string
	}

	// ArticleRepository can fetch Articles
	ArticleRepository interface {
		All(ctx context.Context) []Article
		Get(ctx context.Context, id string) (Article, error)
	}
)
