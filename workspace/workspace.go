package workspace

import (
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/repos"
	"strings"
)

const (
	workspaceUrl = "github.com/goatcms/go-project"
	srcPath      = "src/"
)

type Workspace struct {
	path       string
	repository repos.Repository
}

func NewWorkspace(path string) *Workspace {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return &Workspace{
		path:       path,
		repository: repos.NewRepository(path),
	}
}

func (r *Workspace) Create() error {
	if err := r.repository.Clone(workspaceUrl); err != nil {
		return err
	}
	if err := r.repository.Uninit(); err != nil {
		return err
	}
	return nil
}

func (w *Workspace) LoadRepository(url string) (string, error) {
	repoPath := w.getRepositoryPath(url)
	if filesystem.IsDir(repoPath) {
		return repoPath, nil
	}

	repo := repos.NewGitRepository(repoPath)
	repo.Clone(url)
	return repoPath, nil
}

func (w *Workspace) getRepositoryPath(url string) string {
	//TODO: Remove http:// and https:// from url prefix
	p := w.path + srcPath + url
	if !strings.HasSuffix(p, "/") {
		p = p + "/"
	}
	return p
}
