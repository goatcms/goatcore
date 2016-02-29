package scaffolding

import (
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/repos"
	"os"
	"strings"
	"fmt"
)

type Scaffolding struct {
	Src    string
	Config Config
}

func NewScaffolding(url string) (*Scaffolding, error) {
	repoPath, err := repos.Load(url)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(repoPath, "/") {
		repoPath = repoPath + "/"
	}

	config, err := readConfig(repoPath + ConfigPath)
	if err != nil {
		return nil, err
	}

	s := &Scaffolding{
		Src:    repoPath,
		Config: *config,
	}
	return s, nil
}

func (s *Scaffolding) Build(dest string) error {
	if !strings.HasSuffix(s.Src, "/") {
		s.Src = s.Src + "/"
	}
	if !strings.HasSuffix(dest, "/") {
		dest = dest + "/"
	}

	err := os.MkdirAll(dest, CreateDirMode)
	if err != nil {
		return err
	}

	renderer, err := NewRenderer(s.Src, s.Config.Delimiters)
	if err != nil {
		return err
	}

	loop := filesystem.DirLoop{
		OnFile: execFileFactory(s.Src, dest, renderer),
		OnDir:  execDirFactory(dest),
		Filter: filterFactory([]string{
			".git",
		}),
	}
	if err = loop.Run(s.Src); err != nil {
		return err
	}

	if filesystem.FileExists(s.Src + ScaffoldingDir) {
		if err = filesystem.Copy(s.Src+ScaffoldingDir, dest+ScaffoldingDir); err != nil {
			return err
		}
	}

	if err = filesystem.Copy(s.Src+ConfigPath, dest+ConfigPath); err != nil {
		return err
	}

	//load sub repositories
	fmt.Printf("%v", s.Config)
	for _, sub := range s.Config.Subs {
		subScaffolding, err := NewScaffolding(sub.Url)
		if err != nil {
			return err
		}
		subScaffolding.Build(dest + sub.Path)
	}

	return nil
}

func execFileFactory(src, dest string, renderer *Renderer) func(os.FileInfo, string) error {
	return func(file os.FileInfo, path string) error {
		err := renderer.Render(src+path, dest+path)
		if err != nil {
			return err
		}
		return nil
	}
}

func execDirFactory(dest string) func(os.FileInfo, string) error {
	return func(file os.FileInfo, path string) error {
		err := os.MkdirAll(dest+path, CreateDirMode)
		if err != nil {
			return err
		}
		return nil
	}
}

func filterFactory(exclude []string) func(os.FileInfo, string) bool {
	return func(file os.FileInfo, path string) bool {
		if path == ".git" || path == ConfigPath || path == ScaffoldingDir {
			return false
		}
		for _, element := range exclude {
			if path == element {
				return false
			}
		}
		return true
	}
}
