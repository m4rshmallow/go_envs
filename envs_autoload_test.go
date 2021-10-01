package envs_test

import (
	"envs"
	"testing"

	_ "envs/autoload"
)

func TestAutoloadEnv(t *testing.T) {
	if len(envs.EnvMap) != 3 {
		t.Error("Autoload doesn't work")
		return
	}
}
