package git

import (
	"testing"

	"github.com/goatcms/goatcore/repositories"
)

func TestConnectorAdapterInterface(t *testing.T) {
	t.Parallel()
	var adapter repositories.ConnectorAdapter
	adapter = NewConnector()
	if adapter == nil {
		t.Errorf("NewConnector must return object implements repositories.ConnectorAdapter")
	}
}

func TestIsSupportURL(t *testing.T) {
	t.Parallel()
	adapter := NewConnector()
	if !adapter.IsSupportURL("https://github.com/goatcms/goatcore.git") {
		t.Errorf("https://github.com/goatcms/goatcore.git is supported URL")
	}
}
