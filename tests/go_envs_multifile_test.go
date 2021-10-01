package go_envs_tests

import (
	"testing"

	"github.com/m4rshmallow/go_envs"
)

func TestMultifilesEnv(t *testing.T) {
	expectedEnvs := 6
	go_envs.Init(".env", ".env_extra")

	if len(go_envs.Map) != expectedEnvs {
		t.Errorf("Not enough envs was set up, expected = %d; got = %d\n", expectedEnvs, len(go_envs.Map))
		return
	}
}
