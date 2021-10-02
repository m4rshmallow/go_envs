package test_autoload

import (
	"testing"

	"github.com/m4rshmallow/go_envs"
	_ "github.com/m4rshmallow/go_envs/autoload"
)

func TestAutoloadEnv(t *testing.T) {
	if len(go_envs.Map) != 3 {
		t.Error("Autoload doesn't work")
		return
	}
}
