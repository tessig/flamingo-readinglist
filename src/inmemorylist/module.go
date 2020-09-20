package inmemorylist

import (
	"flamingo.me/dingo"

	"github.com/tessig/flamingo-readinglist/src/app"
	"github.com/tessig/flamingo-readinglist/src/app/domain"
	"github.com/tessig/flamingo-readinglist/src/inmemorylist/application"
)

type (
	// Module to implement in memory storage for ReadingList
	Module struct{}
)

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(domain.ReadingListRepository)).To(new(application.ListRepository)).In(dingo.Singleton)
}

// Depends on other modules
func (m *Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(app.Module),
	}
}
