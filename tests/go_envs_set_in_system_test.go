package go_envs_tests

import (
	"os"
	"testing"

	"github.com/m4rshmallow/go_envs"
)

func TestEnvSetInSystem(t *testing.T) {
	go_envs.Init()
	for k, v := range shouldBe {
		osEnv := os.Getenv(k)
		if osEnv != v {
			t.Errorf("os.env.%s is invalid; expected = %s; got = %s\n", k, v, osEnv)
		}
	}
}
