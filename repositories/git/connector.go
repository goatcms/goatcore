package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/goatcms/goatcore/filesystem/disk"
	"github.com/goatcms/goatcore/repositories"
)

// Connector is wrapper for many repositories connector adapters
type Connector struct{}

// NewConnector create new Connector instance
func NewConnector() *Connector {
	return &Connector{}
}

// IsSupportURL check if repository URL is supported
func (connector *Connector) IsSupportURL(url string) bool {
	return strings.HasPrefix(url, "git://") || strings.HasSuffix(url, ".git")
}

// IsSupportRepo check if local repository is supported
func (connector *Connector) IsSupportRepo(path string) bool {
	return disk.IsExist(path + "/.git")
}

// Clone clone repository to local directory
func (connector *Connector) Clone(url, version, destPath string) (repo repositories.Repository, err error) {
	var (
		out  bytes.Buffer
		args []string
	)
	args = []string{"clone", url, destPath}
	cmd := exec.Command("git", args...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err = cmd.Run(); err != nil {
		return nil, fmt.Errorf("external git app fail %v: %v %v", args, err, string(out.Bytes()))
	}
	return &Repository{
		path: destPath,
	}, nil
}

// Open open repository from local filesystem
func (connector *Connector) Open(path string) (repo repositories.Repository, err error) {
	return &Repository{
		path: path,
	}, nil
}
