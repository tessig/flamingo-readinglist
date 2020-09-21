package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"go.opencensus.io/trace"

	"github.com/tessig/flamingo-readinglist/src/app/domain"
	"github.com/tessig/flamingo-readinglist/src/app/infrastructure/dto"
)

var _ domain.ArticleRepository = &BlogService{}

// BlogService is an implementation for a e external service
type BlogService struct {
	logger  flamingo.Logger
	client  *http.Client
	baseURL string
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
func (b *BlogService) All(ctx context.Context) []domain.Article {
	ctx, span := trace.StartSpan(ctx, "app/infrastructure/blog-service/all")
	defer span.End()
	resp, err := b.client.Get(b.baseURL + "/articles")
	if err != nil {
		b.logger.WithContext(ctx).Error(err)
		return nil
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			b.logger.Warn(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		b.logger.WithContext(ctx).Error("Service responded with http status code: " + strconv.Itoa(resp.StatusCode))
		return nil
	}

	var articles dto.Articles

	err = json.NewDecoder(resp.Body).Decode(&articles)
	if err != nil {
		b.logger.WithContext(ctx).Error(err)
	}

	return articles.MapToDomainModel()
}

// Get single article
func (b *BlogService) Get(ctx context.Context, id string) (domain.Article, error) {
	ctx, span := trace.StartSpan(ctx, "app/infrastructure/blog-service/single")
	defer span.End()
	resp, err := b.client.Get(b.baseURL + "/articles/id/" + id)
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
