package repositories

import (
	"testing"
)

func TestMultiConnector(t *testing.T) {
	t.Parallel()
	var connector Connector
	connector = NewMultiConnector([]ConnectorAdapter{})
	if connector == nil {
		t.Errorf("NewMultiConnector must return object implements repositories.Connector interface")
	}
}
