package app

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"

	"github.com/tessig/flamingo-readinglist/src/app/application/fake"
	"github.com/tessig/flamingo-readinglist/src/app/domain"
	"github.com/tessig/flamingo-readinglist/src/app/interfaces/controller"
)

type (
	// Module basic struct
	Module struct {
		useFakeArticles bool
	}
	routes struct {
		home     *controller.HomeController
		articles *controller.ArticlesController
	}
)

// Inject dependencies
func (m *Module) Inject(
	cfg *struct {
		UseFakeArticles bool `inject:"config:articles.useFake"`
	},
) *Module {
	if cfg != nil {
		m.useFakeArticles = cfg.UseFakeArticles
	}
	return m
}

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))

	if m.useFakeArticles {
		injector.Bind(new(domain.ArticleRepository)).To(new(fake.Articles))
	}
}

// CueConfig for the module
func (m *Module) CueConfig() string {
	return `
articles: {
    useFake: bool|*false
}
`
}

// Inject dependencies
func (r *routes) Inject(
	home *controller.HomeController,
	articles *controller.ArticlesController,
) *routes {
	r.home = home
	r.articles = articles

	return r
}

// Routes definition for the module
func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.MustRoute("/static/*name", `flamingo.static.file(name,dir?="static")`)

	registry.MustRoute("/", "home")
	registry.HandleAny("home", r.home.Home)

	registry.MustRoute("/articles", "articles")
	registry.HandleAny("articles", r.articles.List)

	registry.MustRoute("/articles/add", `articles.add(articleID,listID?="main")`)
	registry.HandleAny("articles.add", r.articles.AddToList)
}
