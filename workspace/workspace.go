package workspace

import (
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/generator"
	"github.com/goatcms/goat-core/repos"
	"github.com/goatcms/goat-core/varutil"
	"os"
	"os/user"
	"path/filepath"
)

const (
	workspaceUrl        = "http://github.com/goatcms/workspace"
	repositoryCachePath = "/.goat/workspace/"
	goatDataPath        = "/.goat/"
	secretDataPath      = goatDataPath + "secrets"
	valuesDataPath      = goatDataPath + "values"
	defaultSrcPath      = "src/"
	mainFile            = "workspace.goat.json"
)

type Workspace struct {
	path       string
	Src        string                       `json:"src"`
	Packages   repos.PackageManager         `json:"packages"`
	Repository repos.Repository             `json:"-"`
	Secrets    *generator.GeneratedResource `json:"-"`
	Values     *generator.GeneratedResource `json:"-"`
}

func NewWorkspace(path string) (*Workspace, error) {
	w := &Workspace{}
	if err := w.Init(path); err != nil {
		return nil, err
	}
	return w, nil
}

func ReadWorkspace(path string) (*Workspace, error) {
	w := &Workspace{
		path: path,
	}
	if err := w.Read(); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *Workspace) Init(path string) error {
	varutil.FixDirPath(&path)
	w.path = path
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	//sub objects
	err = w.InitGenerators()
	if err != nil {
		return err
	}
	//default values
	w.Repository = repos.NewRepository(path)
	w.Src = defaultSrcPath
	w.FixPaths()
	//recursive
	w.Packages.Init(w.path + w.Src)
	return nil
}

func (w *Workspace) InitGenerators() error {
	var err error
	w.Secrets, err = generator.NewGeneratedResource(w.path + secretDataPath)
	if err != nil {
		return err
	}
	w.Values, err = generator.NewGeneratedResource(w.path + valuesDataPath)
	if err != nil {
		return err
	}
	return nil
}

func (w *Workspace) FixPaths() {
	varutil.FixDirPath(&w.Src)
}

func (w *Workspace) Create() error {
	basePath, err := w.getWorkspaceSrcPath()
	if err != nil {
		return err
	}

	if !filesystem.IsExist(w.path) {
		if err := os.MkdirAll(w.path, 0777); err != nil {
			return err
		}
	}
	filesystem.CopyDirectory(basePath, w.path, filterGit)

	err = w.InitGenerators()
	if err != nil {
		return err
	}

	return w.Write()
}

func (w *Workspace) GetAbsPath() string {
	return w.path
}

func (w *Workspace) getWorkspaceSrcPath() (string, error) {
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

func (w *Workspace) Read() error {
	if err := varutil.ReadJson(w.path+mainFile, w); err != nil {
		return err
	}
	if err := w.Init(w.path); err != nil {
		return err
	}
	return nil
}

func (w *Workspace) Write() error {
	return varutil.WriteJson(w.path+mainFile, w)
}

func CleanWorkspaceCache() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}
	cachePath := usr.HomeDir + repositoryCachePath
	return os.RemoveAll(cachePath)
}

func filterGit(info os.FileInfo, path string) bool {
	return path != ".git"
}
