package healthcheck_test

import (
	"context"
	"testing"

	"github.com/tessig/flamingo-readinglist/src/app/domain"
	"github.com/tessig/flamingo-readinglist/src/healthcheck"
)

type (
	repoMock struct {
		articles []domain.Article
	}
)

func (r *repoMock) All(context.Context) []domain.Article {
	return r.articles
}

func (r *repoMock) Get(context.Context, string) (domain.Article, error) {
	return domain.Article{}, nil
}

func TestArticles_Status(t *testing.T) {
	type fields struct {
		repo domain.ArticleRepository
	}
	tests := []struct {
		name        string
		fields      fields
		wantStatus  bool
		wantMessage string
	}{
		{
			name: "ok",
			fields: fields{
				repo: &repoMock{articles: []domain.Article{
					{
						ID: "1",
					},
				}},
			},
			wantStatus:  true,
			wantMessage: "Articles can be fetched",
		},
		{
			name: "not ok",
			fields: fields{
				repo: &repoMock{},
			},
			wantStatus:  false,
			wantMessage: "No articles found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := new(healthcheck.Articles).Inject(tt.fields.repo)
			status, message := a.Status()
			if status != tt.wantStatus {
				t.Errorf("Status() status = %v, want %v", status, tt.wantStatus)
			}
			if message != tt.wantMessage {
				t.Errorf("Status() message = %v, want %v", message, tt.wantMessage)
			}
		})
	}
}
