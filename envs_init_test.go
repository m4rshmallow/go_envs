package envs_test

import (
	"envs"
	"testing"
)

func TestEnvInit(t *testing.T) {
	envs.Init()

	if envs.EnvMap.Get("ENV_1") != shouldBe["ENV_1"] {
		t.Error("ENV doesn't load")
		return
	}

	if len(envs.EnvMap) != 3 {
		t.Error("Not enough envs was set")
		return
	}

	for k, v := range shouldBe {
		currentEnvValue := envs.EnvMap[k]
		if currentEnvValue != v {
			t.Errorf("%s has invalid value - expected = %s; got = %s\n", k, v, currentEnvValue)
			break
		}
	}
}
