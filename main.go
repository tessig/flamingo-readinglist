package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/gotemplate"
	coreHealthcheck "flamingo.me/flamingo/v3/core/healthcheck"
	"flamingo.me/flamingo/v3/core/locale"
	"flamingo.me/flamingo/v3/framework/opencensus"

	"github.com/tessig/flamingo-readinglist/src/app"
	"github.com/tessig/flamingo-readinglist/src/healthcheck"
	"github.com/tessig/flamingo-readinglist/src/inmemorylist"
)

func main() {
	flamingo.App([]dingo.Module{
		new(locale.Module),
		new(gotemplate.Module),
		new(opencensus.Module),
		new(coreHealthcheck.Module),
		new(healthcheck.Module),
		new(inmemorylist.Module),
		new(app.Module),
	})
}
