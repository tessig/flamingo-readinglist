package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/opencensus"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"

	"github.com/tessig/flamingo-readinglist/src/app/domain"
	"github.com/tessig/flamingo-readinglist/src/app/infrastructure/dto"
)

var (
	_              domain.ArticleRepository = &BlogService{}
	stat                                    = stats.Int64("readinglist/blog/response_time", "response time of article fetching", stats.UnitMilliseconds)
	keyEndpoint, _                          = tag.NewKey("endpoint")
)

// BlogService is an implementation for a e external service
type BlogService struct {
	logger  flamingo.Logger
	client  *http.Client
	baseURL string
}

func init() {
	_ = opencensus.View(
		"readinglist/blog/response_time",
		stat,
		view.Distribution(100, 500, 1000, 2500, 5000, 10000),
		keyEndpoint,
	)
}

// Inject dependencies
func (b *BlogService) Inject(
	logger flamingo.Logger,
	client *http.Client,
	cfg *struct {
		BaseURL string `inject:"config:articles.api.baseURL"`
	},
) *BlogService {
	b.logger = logger.WithField(flamingo.LogKeyCategory, "blog-api-client")
	b.client = client
	if cfg != nil {
		b.baseURL = cfg.BaseURL
	}
	return b
}

// All the articles
func (b *BlogService) All(ctx context.Context) ([]domain.Article, error) {
	ctx, span := trace.StartSpan(ctx, "app/infrastructure/blog-service/all")
	defer span.End()

	start := time.Now()
	req, err := http.NewRequest("GET", b.baseURL+"/articles", nil)
	if err != nil {
		return nil, err
	}
	resp, err := b.client.Do(req.WithContext(ctx))
	ctx, _ = tag.New(ctx, tag.Upsert(keyEndpoint, "articles"))
	stats.Record(ctx, stat.M(time.Since(start).Nanoseconds()/1000000))
	if err != nil {
		b.logger.WithContext(ctx).Error(err)
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			b.logger.Warn(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		b.logger.WithContext(ctx).Error("Service responded with http status code: " + strconv.Itoa(resp.StatusCode))
		return nil, err
	}

	var articles dto.Articles

	err = json.NewDecoder(resp.Body).Decode(&articles)
	if err != nil {
		b.logger.WithContext(ctx).Error(err)
		return nil, err
	}

	return articles.MapToDomainModel(), nil
}

// Get single article
func (b *BlogService) Get(ctx context.Context, id string) (domain.Article, error) {
	ctx, span := trace.StartSpan(ctx, "app/infrastructure/blog-service/single")
	defer span.End()

	start := time.Now()

	req, err := http.NewRequest("GET", "/articles/id/"+id, nil)
	if err != nil {
		return domain.Article{}, err
	}
	resp, err := b.client.Do(req.WithContext(ctx))
	ctx, _ = tag.New(ctx, tag.Upsert(keyEndpoint, "articles/id"))
	stats.Record(ctx, stat.M(time.Since(start).Nanoseconds()/1000000))
	if err != nil {
		b.logger.WithContext(ctx).Error(err)
		return domain.Article{}, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			b.logger.Warn(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		b.logger.WithContext(ctx).Error("Service responded with http status code: " + strconv.Itoa(resp.StatusCode))
		return domain.Article{}, errors.New("not found")
	}

	var article dto.Article

	err = json.NewDecoder(resp.Body).Decode(&article)
	if err != nil {
		b.logger.WithContext(ctx).Error(err)
		return domain.Article{}, err
	}

	return article.MapToDomainModel(), nil
}
