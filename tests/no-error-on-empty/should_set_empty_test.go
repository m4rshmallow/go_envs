package test_no_error_on_empty

import (
	"testing"

	"github.com/m4rshmallow/go_envs"
)

func TesShouldSetEmptyEnv(t *testing.T) {
	go_envs.Init(".env_empty")
}

func TestShouldReturnEmptyValue(t *testing.T) {
	go_envs.Init(".env_empty")

	envName := "EMPTY"
	
	if go_envs.Map.Get(envName) != "" {
		t.Errorf("%s doesn't set\n", envName)
		return
	}
}