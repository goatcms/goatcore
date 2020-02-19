package pipelinem

/* Uncomment to test SandboxHealthChecker function manually

import (
	"testing"

	"github.com/goatcms/goatcore/app/mockupapp"
)

func TestSandboxHealthChecker(t *testing.T) {
	t.Parallel()
	var (
		err  error
		mapp *mockupapp.App
		msg  string
	)
	if mapp, _, err = newApp(mockupapp.MockupOptions{}); err != nil {
		t.Error(err)
		return
	}
	if msg, err = SandboxHealthChecker(mapp, mapp.IOContext().Scope()); err != nil {
		t.Errorf("%s, %s", msg, err.Error())
		return
	}
}
*/
