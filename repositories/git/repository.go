package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Repository represent single git repository
type Repository struct {
	path string
}

// NewRepository create new Repository instance
func NewRepository(path string) *Repository {
	return &Repository{
		path: path,
	}
}

// Pull update repository
func (repo *Repository) Pull() (err error) {
	var (
		out bytes.Buffer
	)
	cmd := exec.Command("git", "pull")
	cmd.Dir = repo.path
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("fail execute command git pull (cwd: %v): %v %v", repo.path, err, string(out.Bytes()))
	}
	return nil
}
