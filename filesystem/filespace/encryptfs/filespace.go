package encryptfs

import (
	"os"

	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/encryptfs/cipherfs"
	"github.com/goatcms/goatcore/varutil/idutil"
)

// Settings is set of params for EncryptFS
type Settings struct {
	Salt     []byte
	Secret   []byte
	HostOnly bool
	Cipher   cipherfs.Cipher
}

// EncryptFS is an encrypted filespace
type EncryptFS struct {
	baseFS filesystem.Filespace
	hash   []byte
	Cipher cipherfs.Cipher
}

// NewEncryptFS create new EncryptFS instance
func NewEncryptFS(baseFS filesystem.Filespace, settings Settings) (_ filesystem.Filespace, err error) {
	var hash []byte
	hash = append(hash, settings.Secret...)
	if settings.HostOnly {
		var hostID string
		if hostID, err = idutil.HostID(); err != nil {
			return nil, err
		}
		hash = append(hash, hostID...)
	}
	hash = append(hash, settings.Salt...)
	return &EncryptFS{
		baseFS: baseFS,
		hash:   hash,
		Cipher: settings.Cipher,
	}, nil
}

func (fs *EncryptFS) Copy(src, dest string) error {
	return fs.baseFS.Copy(src, dest)
}

func (fs *EncryptFS) CopyDirectory(src, dest string) error {
	return fs.baseFS.CopyDirectory(src, dest)
}

func (fs *EncryptFS) CopyFile(src, dest string) error {
	return fs.baseFS.CopyFile(src, dest)
}

func (fs *EncryptFS) ReadDir(path string) ([]os.FileInfo, error) {
	return fs.baseFS.ReadDir(path)
}

func (fs *EncryptFS) IsExist(path string) bool {
	return fs.baseFS.IsExist(path)
}

func (fs *EncryptFS) IsFile(path string) bool {
	return fs.baseFS.IsFile(path)
}

func (fs *EncryptFS) IsDir(path string) bool {
	return fs.baseFS.IsDir(path)
}

func (fs *EncryptFS) MkdirAll(path string, filemode os.FileMode) error {
	return fs.baseFS.MkdirAll(path, filemode)
}

func (fs *EncryptFS) ReadFile(path string) (data []byte, err error) {
	if data, err = fs.baseFS.ReadFile(path); err != nil {
		return nil, err
	}
	return fs.Cipher.Decrypt(fs.hash, data)
}

func (fs *EncryptFS) WriteFile(path string, data []byte, perm os.FileMode) (err error) {
	if data, err = fs.Cipher.Encrypt(fs.hash, data); err != nil {
		return err
	}
	return fs.baseFS.WriteFile(path, data, perm)
}

func (fs *EncryptFS) Filespace(path string) (childFS filesystem.Filespace, err error) {
	if childFS, err = fs.baseFS.Filespace(path); err != nil {
		return nil, err
	}
	return &EncryptFS{
		Cipher: fs.Cipher,
		hash:   fs.hash,
		baseFS: childFS,
	}, nil
}

func (fs *EncryptFS) Reader(path string) (reader filesystem.Reader, err error) {
	if reader, err = fs.baseFS.Reader(path); err != nil {
		return nil, err
	}
	return fs.Cipher.DecryptReader(fs.hash, reader)
}

func (fs *EncryptFS) Writer(path string) (writer filesystem.Writer, err error) {
	if writer, err = fs.baseFS.Writer(path); err != nil {
		return nil, err
	}
	return fs.Cipher.EncryptWriter(fs.hash, writer)
}

func (fs *EncryptFS) Remove(path string) error {
	return fs.baseFS.Remove(path)
}

func (fs *EncryptFS) RemoveAll(path string) error {
	return fs.baseFS.RemoveAll(path)
}

func (fs *EncryptFS) Lstat(path string) (os.FileInfo, error) {
	return fs.baseFS.Lstat(path)
}
