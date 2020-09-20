package healthcheck_test

import (
	"testing"

	"flamingo.me/flamingo/v3/framework/config"

	"github.com/tessig/flamingo-readinglist/src/healthcheck"
)

func TestModule_Configure(t *testing.T) {
	if err := config.TryModules(nil, new(healthcheck.Module)); err != nil {
		t.Error(err)
	}
}
