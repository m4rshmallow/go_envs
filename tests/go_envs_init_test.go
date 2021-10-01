package go_envs_tests

import (
	"testing"

	"github.com/m4rshmallow/go_envs"
)

func TestEnvInit(t *testing.T) {
	go_envs.Init()

	if go_envs.EnvMap.Get("ENV_1") != shouldBe["ENV_1"] {
		t.Error("ENV doesn't load")
		return
	}

	if len(go_envs.EnvMap) != 3 {
		t.Error("Not enough envs was set")
		return
	}

	for k, v := range shouldBe {
		currentEnvValue := go_envs.EnvMap[k]
		if currentEnvValue != v {
			t.Errorf("%s has invalid value - expected = %s; got = %s\n", k, v, currentEnvValue)
			break
		}
	}
}