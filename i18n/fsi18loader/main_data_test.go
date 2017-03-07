package fsi18loader

import (
	"strconv"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

const (
	fileContent = `{
		"k1": {
			"k2": {
				"k3": {
					"min_length": "Minimalna długość pola to %v znaków",
					"some1": "ttttttttttttttttttttttttttttttttttttt",
					"some2": "ttttttttttttttttttttttttttttttttttttt",
					"some3": "ttttttttttttttttttttttttttttttttttttt",
					"some4": "ttttttttttttttttttttttttttttttttttttt",
					"some5": "ttttttttttttttttttttttttttttttttttttt",
					"some6": "ttttttttttttttttttttttttttttttttttttt",
					"some7": "ttttttttttttttttttttttttttttttttttttt",
					"some8": "ttttttttttttttttttttttttttttttttttttt",
					"some9": "ttttttttttttttttttttttttttttttttttttt",
					"some10": "ttttttttttttttttttttttttttttttttttttt",
					"some11": "aaaaaaaaaaaaaaaaaaaaa",
					"some12": "aaaaaaaaaaaaaaaaaaaaa",
					"some13": "aaaaaaaaaaaaaaaaaaaaa",
					"some14": "aaaaaaaaaaaaaaaaaaaaa",
					"some15": "aaaaaaaaaaaaaaaaaaaaa",
					"some16": "aaaaaaaaaaaaaaaaaaaaa",
					"some17": "aaaaaaaaaaaaaaaaaaaaa",
					"some18": "aaaaaaaaaaaaaaaaaaaaa",
					"some19": "aaaaaaaaaaaaaaaaaaaaa",
					"some20": "aaaaaaaaaaaaaaaaaaaaa",
				}
			}
		}
	}`
)

func createBenchmarkFilespace(fileperdir, subdirs, maxdeph int) (filesystem.Filespace, error) {
	fs, err := memfs.NewFilespace()
	if err != nil {
		return nil, err
	}
	if err = createDepth(fs, "root/", fileperdir, subdirs, 0, maxdeph); err != nil {
		return nil, err
	}
	return fs, nil
}

func createDepth(fs filesystem.Filespace, path string, n, subdepths, depth, max int) error {
	if depth == max {
		return nil
	}
	if err := createFiles(fs, path+"/", n); err != nil {
		return err
	}
	for i := 0; i < subdepths; i++ {
		istr := strconv.Itoa(i)
		if err := createDepth(fs, path+"/"+istr, n, subdepths, depth+1, max); err != nil {
			return err
		}
	}
	return nil
}

func createFiles(fs filesystem.Filespace, path string, n int) error {
	for i := 0; i < n; i++ {
		istr := strconv.Itoa(i)
		if err := fs.WriteFile(path+istr+".json", []byte(fileContent), 0777); err != nil {
			return err
		}
	}
	return nil
}
