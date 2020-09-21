package dto

import (
	"net/url"
	"time"

	"github.com/tessig/flamingo-readinglist/src/app/domain"
)

// Articles is list of article
type Articles []Article

// Article is a domain transfer object
type Article struct {
	ID     string    `json:"id"`
	Title  string    `json:"title"`
	Link   string    `json:"link"`
	Author string    `json:"author"`
	Date   time.Time `json:"date"`
	Tags   []string  `json:"tags"`
}

// MapToDomainModel maps DTOs to the actual domain model
func (a Articles) MapToDomainModel() []domain.Article {
	var articles []domain.Article

	for _, dtoArticle := range a {
		articles = append(articles, dtoArticle.MapToDomainModel())
	}

	return articles
}

// MapToDomainModel maps DTOs to the actual domain model
func (a Article) MapToDomainModel() domain.Article {
	link, err := url.Parse(a.Link)
	if err != nil {
		link = &url.URL{}
	}

	return domain.Article{
		ID:     a.ID,
		Title:  a.Title,
		Author: a.Author,
		Link:   *link,
		Date:   a.Date,
		Tags:   a.Tags,
	}
}
