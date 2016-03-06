package scaffolding

import (
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/repos"
	"github.com/goatcms/goat-core/workspace"
	"os"
	"strings"
)

type Scaffolding struct {
	workspace Workspace
	Src       string
	Config    Config
}

func NewScaffolding(url string) (*Scaffolding, error) {
	repo := NewRepository(url)
	repoPath, err := repos.Load(url)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(repoPath, "/") {
		repoPath = repoPath + "/"
	}

	config := &Config{}
	if err = config.Read(repoPath + ConfigPath); err != nil {
		return nil, err
	}

	s := &Scaffolding{
		Src:    repoPath,
		Config: *config,
	}
	return s, nil
}

func (s *Scaffolding) BuildSource(dest string) error {
	if !strings.HasSuffix(s.Src, "/") {
		s.Src = s.Src + "/"
	}
	if !strings.HasSuffix(dest, "/") {
		dest = dest + "/"
	}

	srcSecretsPath := s.Src + GenerateSecretsPath
	secretsGenerator := NewGenerator()
	if err := secretsGenerator.LoadDefinitions(srcSecretsPath); err != nil {
		return err
	}
	secretsGenerator.LoadValues(srcSecretsPath)
	secretsGenerator.GenerateValues()

	srcValuesPath := s.Src + GenerateValuesPath
	valuesGenerator := NewGenerator()
	if err := valuesGenerator.LoadDefinitions(srcValuesPath); err != nil {
		return err
	}
	valuesGenerator.LoadValues(srcValuesPath)
	valuesGenerator.GenerateValues()

	rendererData := RendererData{
		Secrets: secretsGenerator.Values,
		Values:  valuesGenerator.Values,
	}

	s.BuildModule(dest, &rendererData)

	destSecretsPath := dest + GenerateSecretsPath
	if err := secretsGenerator.SaveValues(destSecretsPath); err != nil {
		return err
	}

	destValuesPath := dest + GenerateValuesPath
	if err := valuesGenerator.SaveValues(destValuesPath); err != nil {
		return err
	}

	return nil
}

func (s *Scaffolding) BuildModule(dest string, rendererData *RendererData) error {
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

	//for static module (without scaffolding file)
	scaffoldingFile := s.Src + ConfigPath
	if filesystem.IsFile(scaffoldingFile) {
		if err = filesystem.Copy(scaffoldingFile, dest+ConfigPath); err != nil {
			return err
		}
	}

	renderer, err := NewRenderer(s.Src, s.Config.Delimiters, rendererData)
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

	if err = filesystem.Copy(scaffoldingFile, dest+ConfigPath); err != nil {
		return err
	}

	//load modules
	for _, module := range s.Config.Modules {
		moduleScaffolding, err := NewScaffolding(module.Url)
		if err != nil {
			return err
		}
		moduleScaffolding.BuildModule(dest+module.Path, rendererData)
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
