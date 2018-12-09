package fshelper

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestModulesFromFile(t *testing.T) {
	var (
		err    error
		srcFS  filesystem.Filespace
		destFS filesystem.Filespace
		tmp    []byte
		tmps   string
	)
	t.Parallel()
	// prepare data
	if srcFS, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if destFS, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	if err = srcFS.WriteFile("test.txt", []byte("ContentOfTestTXT"), filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}
	if err = srcFS.WriteFile("path/to/dep/test/file.txt", []byte("ContentOfDepFile"), filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}
	if err = destFS.WriteFile("doesnt/delete/exist/files.txt", []byte("somecontent"), filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}
	if err = Copy(srcFS, destFS, nil); err != nil {
		t.Error(err)
		return
	}
	// test if copy top lvl file
	if tmp, err = destFS.ReadFile("test.txt"); err != nil {
		t.Error(err)
		return
	}
	tmps = string(tmp)
	if tmps != "ContentOfTestTXT" {
		t.Errorf("Incorrect content of 'test.txt' file. Expected 'ContentOfTestTXT' and take %s", tmps)
		return
	}
	// test if copy dep file
	if tmp, err = destFS.ReadFile("path/to/dep/test/file.txt"); err != nil {
		t.Error(err)
		return
	}
	tmps = string(tmp)
	if tmps != "ContentOfDepFile" {
		t.Errorf("Incorrect content of 'test.txt' file. Expected 'ContentOfDepFile' and take %s", tmps)
		return
	}
	// test if no modyfied dest files
	if tmp, err = destFS.ReadFile("doesnt/delete/exist/files.txt"); err != nil {
		t.Error(err)
		return
	}
	tmps = string(tmp)
	if tmps != "somecontent" {
		t.Errorf("Incorrect content of 'doesnt/delete/exist/files.txt' file. Expected 'somecontent' and take %s", tmps)
		return
	}
}
