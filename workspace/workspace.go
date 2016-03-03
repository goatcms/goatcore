package workspace

import (
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/repos"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	workspaceUrl        = "github.com/goatcms/go-project"
	repositoryCachePath = "/.goat/workspace/"
	srcPath             = "src/"
)

type Workspace struct {
	path       string
	repository repos.Repository
}

func NewWorkspace(path string) (*Workspace, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return &Workspace{
		path:       path,
		repository: repos.NewRepository(path),
	}, nil
}

func (w *Workspace) Create() error {
	basePath, err := w.getWorkspaceBasePath()
	if err != nil {
		return err
	}

	if !filesystem.IsExist(w.path) {
		if err := os.MkdirAll(w.path, 0777); err != nil {
			return err
		}
	}
	filesystem.CopyDirectory(basePath, w.path, filterGit)
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

func (w *Workspace) GetAbsPath() string {
	return w.path
}

func (w *Workspace) getRepositoryPath(url string) string {
	//TODO: Remove http:// and https:// from url prefix
	p := w.path + srcPath + url
	if !strings.HasSuffix(p, "/") {
		p = p + "/"
	}
	return p
}

func (w *Workspace) getWorkspaceBasePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	cachePath := usr.HomeDir + repositoryCachePath
	if filesystem.IsExist(cachePath) {
		return cachePath, nil
	}
	workspaceRepo := repos.NewRepository(cachePath)
	if err := workspaceRepo.Clone(workspaceUrl); err != nil {
		return "", err
	}
	return cachePath, nil
}

func filterGit(info os.FileInfo, path string) bool {
	return path != ".git"
}
