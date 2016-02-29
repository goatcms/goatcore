package repos

import (
	"fmt"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/system"
	"os"
)

func Load(url string) (string, error) {
	repoPath, err := GetRepoPath(url)
	if err != nil {
		return "", err
	}

  if filesystem.IsDir(repoPath) {
    return repoPath, nil
  }

  _, err = system.RunCommand("git clone", url, repoPath)
	if err != nil {
		return "", err
	}

  return repoPath, nil
}

func getGoPath() (string, error) {
	path := os.Getenv("GOPATH")
	if path == "" {
		return "", fmt.Errorf("$GOPATH must be set")
	}
	return path, nil
}

func GetRepoPath(url string) (string, error) {
	path, err := getGoPath()
	if err != nil {
		return "", err
	}
	return path + "/src/" + url, nil
}
