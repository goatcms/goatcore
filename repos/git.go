package repos

import (
	"github.com/goatcms/goat-core/filesystem"
	"os"
	"os/exec"
	"strings"
)

type GitRepository struct {
	path string
}

func NewGitRepository(path string) Repository {
	if strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return &GitRepository{
		path: path,
	}
}

func (r *GitRepository) Clone(url string) error {
	if !filesystem.IsDir(r.path) {
		os.MkdirAll(r.path, 0777)
	}

	cmd := exec.Command("git", "clone", "http://"+url, r.path)
	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}

	return nil
}

func (r *GitRepository) Uninit() error {
	return os.RemoveAll(r.path + ".git")
}
