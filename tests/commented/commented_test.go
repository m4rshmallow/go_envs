package test_commented_not_export

import (
	"testing"

	"github.com/m4rshmallow/go_envs"
)

func TestCommentedNotExport(t *testing.T) {
	defer func() {
		err := recover()

		if err == nil {
			t.Error("Exported hidden env")
		}
	}()

	go_envs.Init("../.env")
	go_envs.Map.Get("#ENV_4")
}