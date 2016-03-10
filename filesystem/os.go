package filesystem

import (
	"os"
)

func MkdirAll(dest string) error  {
	return os.MkdirAll(dest, DefaultMode)
}
