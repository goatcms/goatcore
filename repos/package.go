package repos

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/goatcms/goatcore/filesystem/disk"
	"github.com/goatcms/goatcore/varutil"
)

type RepositoryPackage struct {
	Url string `json:"url"`
	Rev string `json:"rev"`
}

type PackageManager struct {
	path         string
	Repositories map[string]RepositoryPackage `json:"repositories"`
}

func NewPackageManager(srcPath string) *PackageManager {
	p := &PackageManager{}
	p.Init(srcPath)
	return p
}

func (p *PackageManager) Init(path string) {
	varutil.FixDirPath(&path)
	if p.Repositories == nil {
		p.Repositories = map[string]RepositoryPackage{}
	}
	p.path = path
}

func (p *PackageManager) CreteRecord(depId string) (string, error) {
	localPath, url, rev := p.decodeDepStr(depId)

	if _, exists := p.Repositories[localPath]; exists {
		return "", fmt.Errorf("Record " + localPath + " exist")
	}

	p.Repositories[localPath] = RepositoryPackage{
		Url: url,
		Rev: rev,
	}

	return localPath, nil
}

func (p *PackageManager) Get(localPath string) (string, error) {
	fullPath := p.path + localPath
	record, exists := p.Repositories[localPath]
	if !exists {
		return "", fmt.Errorf("Record for \""+localPath+"\" path not exist %v", p.Repositories)
	}
	if !disk.IsExist(fullPath) {
		if err := p.load(fullPath, record.Url, record.Rev); err != nil {
			return "", err
		}
	}
	return fullPath, nil
}

func (p *PackageManager) UpdateAll() error {
	for localPath, record := range p.Repositories {
		if err := p.update(localPath, record); err != nil {
			return err
		}
	}
	return nil
}

func (p *PackageManager) UpdateRepo(localPath string) error {
	record, exists := p.Repositories[localPath]
	if exists {
		return fmt.Errorf("Package " + localPath + " no found")
	}
	return p.update(localPath, record)
}

func (p *PackageManager) update(localPath string, record RepositoryPackage) error {
	fullPath := p.path + localPath
	if !disk.IsExist(fullPath) {
		if err := p.load(fullPath, record.Url, record.Rev); err != nil {
			return err
		}
	} else {
		repo := NewRepository(fullPath)
		if err := repo.Pull(); err != nil {
			return err
		}
		if record.Rev != "" {
			if err := repo.Checkout(record.Rev); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *PackageManager) load(fullPath, url, rev string) error {
	if disk.IsExist(fullPath) {
		return fmt.Errorf("Directory " + fullPath + " exist")
	}
	repo := NewRepository(fullPath)
	if err := repo.Clone(url); err != nil {
		return err
	}
	if rev != "" {
		if err := repo.Checkout(rev); err != nil {
			return err
		}
	}
	return nil
}

func (p *PackageManager) decodeDepStr(url string) (string, string, string) {
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	rev := filepath.Ext(url)
	url = strings.TrimSuffix(url, rev)
	localPath := url
	if rev != "" {
		rev = rev[1:]
	}
	url = "http://" + localPath
	return localPath, url, rev
}
