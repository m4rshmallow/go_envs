package test_exported

import (
	"os"
	"testing"

	"github.com/m4rshmallow/go_envs"
)

func TestEnvSetInSystem(t *testing.T) {
	go_envs.Init("../.env")
	
	for k, v := range go_envs.MockMap {
		osEnv := os.Getenv(k)
		if osEnv != v {
			t.Errorf("os.env.%s is invalid; expected = %s; got = %s\n", k, v, osEnv)
		}
	}
}
