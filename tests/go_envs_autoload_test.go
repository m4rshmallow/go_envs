package go_envs_tests

import (
	"testing"

	"github.com/m4rshmallow/go_envs"
	_ "github.com/m4rshmallow/go_envs/autoload"
)

func TestAutoloadEnv(t *testing.T) {
	if len(go_envs.EnvMap) != 3 {
		t.Error("Autoload doesn't work")
		return
	}
}
