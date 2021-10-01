package envs_test

import (
	"envs"
	"os"
	"testing"
)

func TestEnvSetInSystem(t *testing.T) {
	envs.Init()
	for k, v := range shouldBe {
		osEnv := os.Getenv(k)
		if osEnv != v {
			t.Errorf("os.env.%s is invalid; expected = %s; got = %s\n", k, v, osEnv)
		}
	}
}
