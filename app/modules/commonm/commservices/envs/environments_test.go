package envs

import (
	"testing"

	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
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
	result := envs.All()
	if len(result) != 1 {
		t.Errorf("Expected 1 element in result map")
	}
	if result["SOME_KEY"] != "SOME_VALUE" {
		t.Errorf("value for SOME_KEY must be equals to SOME_VALUE")
	}
}

func TestCertStory(t *testing.T) {
	t.Parallel()
	var (
		envs = NewEnvironments()
	)
	envs.SetSSHCert(commservices.SSHCert{
		Secret: "Secret",
		Public: "Public",
	})
	result := envs.SSHCert()
	if result.Public != "Public" {
		t.Errorf("Expected 'Public'")
	}
	if result.Secret != "Secret" {
		t.Errorf("Expected 'Secret'")
	}
}
