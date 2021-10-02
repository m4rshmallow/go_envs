package go_envs_tests

import (
	"testing"

	"github.com/m4rshmallow/go_envs"
)

func TestSkipEmpty(t *testing.T) {
	go_envs.Init(".env_empty")
}