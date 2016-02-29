package scaffolding
/*
import (
	"github.com/goatcms/goat-core/filesystem"
	"os"
	"strings"
)

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

func BuildFiles(src, dest string, exclude []string) error {
	if !strings.HasSuffix(src, "/") {
		src = src + "/"
	}
	if !strings.HasSuffix(dest, "/") {
		dest = dest + "/"
	}

	err := os.MkdirAll(dest, CreateDirMode)
	if err != nil {
		return err
	}

	renderer, err := NewRenderer(src)
	if err != nil {
		return err
	}

	loop := filesystem.DirLoop{
		OnFile: execFileFactory(src, dest, renderer),
		OnDir:  execDirFactory(dest),
		Filter: filterFactory(exclude),
	}
	if err = loop.Run(src); err != nil {
		return err
	}

	if filesystem.FileExists(src + ScaffoldingDir) {
		if err = filesystem.Copy(src+ScaffoldingDir, dest+ScaffoldingDir); err != nil {
			return err
		}
	}

	if err = filesystem.Copy(src+ConfigPath, dest+ConfigPath); err != nil {
		return err
	}

	return nil
}
*/
