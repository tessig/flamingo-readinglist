package domain

import (
	"context"

	"flamingo.me/flamingo/v3/framework/opencensus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

type (
	// ReadingList stores Articles for future reading
	ReadingList struct {
		ID       string
		Articles []Article
	}

	// ReadingListRepository can store and load ReadingLists
	ReadingListRepository interface {
		Load(ctx context.Context, id string) (*ReadingList, error)
		Save(ctx context.Context, l *ReadingList) error
	}
)

var (
	stat         = stats.Int64("readinglist/articles/amount", "Amount of added articles", stats.UnitDimensionless)
	keyListID, _ = tag.NewKey("ListID")
)

func init() {
	_ = opencensus.View("readinglist/articles/amount", stat, view.Count(), keyListID)
}

// AddArticle to the list
func (l *ReadingList) AddArticle(ctx context.Context, a Article) {

	ctx, _ = tag.New(ctx, tag.Upsert(keyListID, l.ID))
	stats.Record(ctx, stat.M(1))
	l.Articles = append(l.Articles, a)
}
