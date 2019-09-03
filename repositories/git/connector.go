package git

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/goatcms/goatcore/filesystem/disk"
	"github.com/goatcms/goatcore/repositories"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Connector is wrapper for many repositories connector adapters
type Connector struct{}

// NewConnector create new Connector instance
func NewConnector() *Connector {
	return &Connector{}
}

// IsSupportURL check if repository URL is supported
func (connector *Connector) IsSupportURL(url string) bool {
	return strings.HasPrefix(url, "git+") || strings.HasPrefix(url, "git://") || strings.HasSuffix(url, ".git")
}

// IsSupportRepo check if local repository is supported
func (connector *Connector) IsSupportRepo(path string) bool {
	return disk.IsExist(path + "/.git")
}

// Clone clone repository to local directory
func (connector *Connector) Clone(url string, version repositories.Version, destPath string) (repo repositories.Repository, err error) {
	var (
		out  bytes.Buffer
		args []string
	)
	if version.Branch == "" {
		version.Branch = "master"
	}
	if strings.HasPrefix(url, "git+") {
		url = url[len("git+"):]
	}
	// clone
	args = []string{"clone", "--branch", version.Branch, url, destPath}
	cmd := exec.Command("git", args...)
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err = cmd.Run(); err != nil {
		return nil, goaterr.Errorf("external git app fail %v: %v %v", args, err, string(out.Bytes()))
	}
	// checkout
	if version.Revision == "" {
		return &Repository{
			path: destPath,
		}, nil
	}
	args = []string{"checkout", version.Branch}
	cmd = exec.Command("git", args...)
	cmd.Dir = destPath
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err = cmd.Run(); err != nil {
		return nil, goaterr.Errorf("external git app fail %v: %v %v", args, err, string(out.Bytes()))
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
