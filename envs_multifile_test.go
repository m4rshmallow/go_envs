package envs_test

import (
	"envs"
	"testing"
)

func TestMultifilesEnv(t *testing.T) {
	expectedEnvs := 6
	envs.Init(".env", ".env_extra")

	if len(envs.EnvMap) != expectedEnvs {
		t.Errorf("Not enough envs was set up, expected = %d; got = %d\n", expectedEnvs, len(envs.EnvMap))
		return
	}
}
