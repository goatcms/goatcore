package disk

import "os"

// IsExist return true if file or directory exists
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// IsDir return true if directory exists
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// IsFile return true if file exists
func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}

// MkdirAll create all path nodes
func MkdirAll(dest string, filemode os.FileMode) error {
	return os.MkdirAll(dest, filemode)
}
