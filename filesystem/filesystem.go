package filesystem

import (
	"os"
	"io/ioutil"
)

func FileExists(p string) bool {
	_, err := os.Stat(p)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func ForAllFiles(path string, predicate func(os.FileInfo, string)) error {
	list, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, dir := range list {
		if dir.Name() == "." || dir.Name() == ".." {
			continue
		}

		newPath := path + "/" + dir.Name()
		if dir.IsDir() {
			ForAllFiles(newPath, predicate)
		} else {
			predicate(dir, newPath)
		}
	}

	return nil
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

/*
func CopyFile(dst, src string) (int64, os.Error) {
	sf, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer sf.Close()
	df, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer df.Close()
	return io.Copy(df, sf)
}*/
