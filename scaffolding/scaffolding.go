package scaffolding

import (
	"fmt"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/varutil"
	"github.com/goatcms/goat-core/workspace"
)

const (
	templatesPath   = ".goat/templates"
	scaffoldingPath = "scaffolding.goat.json"
)

type Delimiters struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type Module struct {
	Package string `json:"packag"`
	Dest    string `json:"dest"`
}

type Template struct {
	Repository string `json:"path,omitempty"`
	Src        string `json:"src,omitempty"`
	Dest       string `json:"dest"`
}

type Scaffolding struct {
	workspace  *workspace.Workspace
	path       string
	Delimiters Delimiters            `json:"delimiters"`
	Modules    []Module              `json:"modules"`
	Suffixes   []string              `json:"suffixes"`
	On         map[string][]Template `json:"on"`
}

func NewScaffolding(w *workspace.Workspace, p string) (*Scaffolding, error) {
	s := &Scaffolding{}
	if err := s.Init(w, p); err != nil {
		return nil, err
	}
	return s, nil
}

func ReadScaffolding(w *workspace.Workspace, p string) (*Scaffolding, error) {
	s := &Scaffolding{}
	if err := s.Init(w, p); err != nil {
		return nil, err
	}
	if err := s.Read(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Scaffolding) Init(w *workspace.Workspace, p string) error {
	varutil.FixDirPath(&p)
	s.workspace = w
	s.path = p
	s.Delimiters.Left = "<<<"
	s.Delimiters.Right = ">>>"
	s.Suffixes = []string{".css", ".sass", ".scss", ".html", ".xhtml", ".htm",
		".js", ".jsx", ".php", ".py", ".go", ".c", ".cpp", ".h", ".hpp", ".rb",
		".mk", ".md"}
	return nil
}

func (s *Scaffolding) Read() error {
	return varutil.ReadJson(s.path+scaffoldingPath, s)
}

func (s *Scaffolding) IsScaffolding() bool {
	return filesystem.IsFile(s.path + scaffoldingPath)
}

func (s *Scaffolding) Write() error {
	return varutil.WriteJson(s.path+scaffoldingPath, s)
}

func (s *Scaffolding) AddModule(packageId, dest string) (Module, error) {
	m := Module{
		Package: packageId,
		Dest:    dest,
	}
	s.Modules = append(s.Modules, m)
	return m, nil
}

func (s *Scaffolding) BuildModule(src string) error {
	varutil.FixDirPath(&src)
	varutil.FixDirPath(&s.path)
	if filesystem.IsExist(s.path) {
		return fmt.Errorf("A directory '" + s.path + "' exists")
	}
	builder := Builder{
		Src:          src,
		Suffixes:     s.Suffixes,
		FileRenderer: FileRenderer{},
	}
	if err := builder.FileRenderer.Init(s.Delimiters); err != nil {
		return err
	}
	if err := builder.Build(s.path); err != nil {
		return err
	}

	if !s.IsScaffolding() {
		//no process static templates
		//scaffolding.goat.json is not required
		return nil
	}

	if err := s.Read(); err != nil {
		return err
	}

	for _, module := range s.Modules {
		if err := s.BuildSubModule(module); err != nil {
			return err
		}
	}

	return nil
}

func (s *Scaffolding) BuildSubModule(module Module) error {
	varutil.FixDirPath(&s.path)
	dest := s.path + module.Dest
	src, err := s.workspace.Packages.Get(module.Package)
	if err != nil {
		return err
	}

	subScaffolding, err := NewScaffolding(s.workspace, dest)
	if err != nil {
		return err
	}
	if err := subScaffolding.BuildModule(src); err != nil {
		return err
	}

	return nil
}

func (s *Scaffolding) BuildResource(r Resource) error {
	varutil.FixDirPath(&s.path)
	if templates, exist := s.On[r.Type]; exist {
		for _, template := range templates {
			if err := s.renderTemplate(r, template); err != nil {
				return err
			}
		}
	}
	for _, module := range s.Modules {
		subModuleScaffolding, err := NewScaffolding(s.workspace, module.Dest)
		if err != nil {
			return err
		}
		subModuleScaffolding.BuildResource(r)
	}
	return nil
}

func (s *Scaffolding) renderTemplate(r Resource, t Template) error {
	if t.Repository == "" && (t.Src == "") {
		return fmt.Errorf("You must defined source Repository or file/directory")
	}
	if t.Repository == "" && (t.Src == "." || t.Src == "/") {
		return fmt.Errorf("Source file/directory has illegal value")
	}
	basePath := s.path
	if t.Repository != "" {
		var err error
		basePath, err = s.workspace.Packages.Get(t.Repository)
		if err != nil {
			return err
		}
	}
	varutil.FixDirPath(&basePath)
	basePath = basePath + t.Src
	builder := Builder{
		Src:      basePath,
		Suffixes: s.Suffixes,
		FileRenderer: FileRenderer{
			Data: r,
		},
	}
	if err := builder.FileRenderer.Init(s.Delimiters); err != nil {
		return err
	}
	if err := builder.Build(s.path); err != nil {
		return err
	}
	return nil
}
