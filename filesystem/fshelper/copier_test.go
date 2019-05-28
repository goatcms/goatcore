package fshelper

import (
	"testing"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestCopierCopyFileStory(t *testing.T) {
	var (
		err      error
		srcFS    filesystem.Filespace
		destFS   filesystem.Filespace
		tmp      []byte
		result   string
		expected = "ContentOfTestTXT"
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
	if err = srcFS.WriteFile("dir/test.txt", []byte(expected), filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}

	if err = (Copier{
		SrcFS:    srcFS,
		SrcPath:  "dir/test.txt",
		DestFS:   destFS,
		DestPath: "dest/test.txt",
	}).Do(); err != nil {
		t.Error(err)
		return
	}
	// test if copy top lvl file
	if tmp, err = destFS.ReadFile("dest/test.txt"); err != nil {
		t.Error(err)
		return
	}
	result = string(tmp)
	if result != expected {
		t.Errorf("Incorrect content of 'test.txt' file. Expected '%s' and take %s", expected, result)
		return
	}
}

func TestCopierCopyDirectoryStory(t *testing.T) {
	var (
		err      error
		srcFS    filesystem.Filespace
		destFS   filesystem.Filespace
		tmp      []byte
		result   string
		expected = "ContentOfTestTXT"
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
	if err = srcFS.WriteFile("dir/test.txt", []byte(expected), filesystem.DefaultUnixFileMode); err != nil {
		t.Error(err)
		return
	}

	if err = (Copier{
		SrcFS:    srcFS,
		SrcPath:  "dir",
		DestFS:   destFS,
		DestPath: "dest",
	}).Do(); err != nil {
		t.Error(err)
		return
	}
	// test if copy top lvl file
	if tmp, err = destFS.ReadFile("dest/test.txt"); err != nil {
		t.Error(err)
		return
	}
	result = string(tmp)
	if result != expected {
		t.Errorf("Incorrect content of 'test.txt' file. Expected '%s' and take %s", expected, result)
		return
	}
}
