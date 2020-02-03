package envs

import (
	"testing"
)

func TestEnvironmentssetStory(t *testing.T) {
	t.Parallel()
	var (
		envs = NewEnvironments()
		err  error
	)
	if err = envs.SetAll(map[string]string{
		"SOME_KEY": "SOME_VALUE",
	}); err != nil {
		t.Error(err)
		return
	}
	result := envs.GetAll()
	if len(result) != 1 {
		t.Errorf("Expected 1 element in result map")
	}
	if result["SOME_KEY"] != "SOME_VALUE" {
		t.Errorf("value for SOME_KEY must be equals to SOME_VALUE")
	}
}
