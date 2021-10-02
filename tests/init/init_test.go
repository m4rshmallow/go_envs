package test_init

import (
	"fmt"
	"testing"

	"github.com/m4rshmallow/go_envs"
)

func TestEnvInit(t *testing.T) {
	go_envs.Init("../.env")

	fmt.Println(go_envs.Map)
	if go_envs.Map.Get("ENV_1") != go_envs.MockMap["ENV_1"] {
		t.Error("ENV doesn't load")
		return
	}

	if len(go_envs.Map) < 3 {
		t.Errorf("Not enough envs was set, expected - %d, got - %d\n", 3, len(go_envs.Map))
		return
	}

	for k, v := range go_envs.MockMap {
		currentEnvValue := go_envs.Map[k]
		if currentEnvValue != v {
			t.Errorf("%s has invalid value - expected = %s; got = %s\n", k, v, currentEnvValue)
			break
		}
	}
}