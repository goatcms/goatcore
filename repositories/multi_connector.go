package repositories

import "github.com/goatcms/goatcore/varutil/goaterr"

// MultiConnector is wrapper for many repositories connector adapters
type MultiConnector struct {
	conns []ConnectorAdapter
}

// NewMultiConnector create new MultiConnector instance
func NewMultiConnector(conns []ConnectorAdapter) *MultiConnector {
	return &MultiConnector{
		conns: conns,
	}
}

// Clone clone repository to local directory
func (connector *MultiConnector) Clone(url string, version Version, destPath string) (repo Repository, err error) {
	for _, adapter := range connector.conns {
		if adapter.IsSupportURL(url) {
			return adapter.Clone(url, version, destPath)
		}
	}
	return nil, goaterr.Errorf("Unsupported url %v (no match any ConnectorAdapter for it)", url)
}

// Open open repository from local filesystem
func (connector *MultiConnector) Open(path string) (repo Repository, err error) {
	for _, adapter := range connector.conns {
		if adapter.IsSupportRepo(path) {
			return adapter.Open(path)
		}
	}
	return nil, goaterr.Errorf("Unsupported repository in path %v (no match any ConnectorAdapter for it)", path)
}
