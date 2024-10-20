package tasker_test

import (
	"testing"

	"github.com/avila-r/env"
	"github.com/avila-r/tasker"
)

var (
	// path defines the location of the .env file used to load environment variables.
	// You can modify this variable to point to a different path if needed.
	path = tasker.RootPath
)

var (
	required_envs = []string{
		"SERVER_URL",
		"POSTGRES_DSN",
		"POSTGRES_DB_NAME",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
	}
)

func Test_Environment(t *testing.T) {
	verify := func(key string) {
		v := env.Get(key)

		if v == "" {
			t.Errorf("%v variable is missing", key)
		}
	}

	if err := env.Load(path); err != nil {
		t.Errorf("error while loading .env file: %v", err.Error())
	}

	for _, e := range required_envs {
		verify(e)
	}
}
