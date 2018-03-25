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
	cmd := exec.Command("git", "-C", repo.path, "pull")
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("fail execute command git -C %v pull: %v %v", repo.path, err, string(out.Bytes()))
	}
	return nil
}
