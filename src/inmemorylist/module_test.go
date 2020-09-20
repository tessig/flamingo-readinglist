package inmemorylist_test

import (
	"testing"

	"flamingo.me/flamingo/v3/framework/config"

	"github.com/tessig/flamingo-readinglist/src/inmemorylist"
)

func TestModule_Configure(t *testing.T) {
	if err := config.TryModules(nil, new(inmemorylist.Module)); err != nil {
		t.Error(err)
	}
}
