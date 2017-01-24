package file

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/goatcms/goat-core/app"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/varutil"
)

type File struct {
	filespace  filesystem.Filespace
	path       string
	mime       string
	name       string
	createTime time.Time
}

func NewTMPFile(filespace filesystem.Filespace, name, mime string) (filesystem.File, error) {
	now := time.Now()
	unixTimeStr := strconv.FormatInt(now.Unix(), 36)
	reg, err := regexp.Compile("[^A-Za-z0-9_]+")
	if err != nil {
		return nil, err
	}
	ext := filepath.Ext(name)
	base := name[:len(name)-(len(ext)+1)]
	safeMIME := strings.Replace(mime, "/", "_", -1)
	safeMIME = reg.ReplaceAllString(safeMIME, "")
	safeBase := reg.ReplaceAllString(base, "")
	safeExt := reg.ReplaceAllString(ext, "")
	randStr := varutil.RandString(10, varutil.AlphaNumericBytes)
	return File{
		filespace:  filespace,
		path:       unixTimeStr + "." + safeMIME + "." + safeBase + "." + randStr + "." + safeExt,
		mime:       mime,
		createTime: now,
		name:       base,
	}, nil
}

func OpenTMPFile(filespace filesystem.Filespace, path string) (filesystem.File, error) {
	if !filespace.IsFile(path) {
		return nil, fmt.Errorf("%v is not a file", path)
	}
	parts := strings.Split(path, ".")
	createTimeUnix, err := strconv.ParseInt(parts[0], 36, 64)
	if err != nil {
		return nil, err
	}
	createTime := time.Unix(createTimeUnix, 0)
	mime := strings.Replace(parts[1], "_", "/", -1)
	return File{
		filespace:  filespace,
		path:       path,
		createTime: createTime,
		mime:       mime,
		name:       parts[2],
	}, nil
}

func (f File) Filespace() filesystem.Filespace {
	return f.filespace
}

func (f File) Path() string {
	return f.path
}

func (f File) IsExist() bool {
	return f.filespace.IsExist(f.path)
}

func (f File) IsFile() bool {
	return f.filespace.IsExist(f.path)
}

func (f File) ReadFile() ([]byte, error) {
	return f.filespace.ReadFile(f.path)
}

func (f File) WriteFile(data []byte, perm os.FileMode) error {
	return f.filespace.WriteFile(f.path, data, perm)
}

func (f File) Reader() (filesystem.Reader, error) {
	return f.filespace.Reader(f.path)
}

func (f File) Writer() (filesystem.Writer, error) {
	return f.filespace.Writer(f.path)
}

func (f File) Remove() error {
	return f.filespace.Remove(f.path)
}

func (f File) MIME() string {
	return f.mime
}

func (f File) CreateTime() time.Time {
	return f.createTime
}

func (f File) Name() string {
	return f.name
}

func (f File) Defer(scope app.EventScope) {
	scope.On(app.RollbackEvent, func(interface{}) error {
		return f.Remove()
	})
}
