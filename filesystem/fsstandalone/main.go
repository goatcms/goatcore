package file

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil"
)

// StandaloneFile is a standalone file (it cointains its filesystem handler)
type StandaloneFile struct {
	filespace  filesystem.Filespace
	path       string
	mime       string
	name       string
	createTime time.Time
}

// NewStandaloneFile create new standalone file instance
func NewStandaloneFile(filespace filesystem.Filespace, path, mime string) (filesystem.File, error) {
	now := time.Now()
	base := filepath.Base(path)
	return &StandaloneFile{
		filespace:  filespace,
		path:       path,
		mime:       mime,
		createTime: now,
		name:       base,
	}, nil
}

// NewTMPStandaloneFile create new temporary standalone file instance
func NewTMPStandaloneFile(filespace filesystem.Filespace, name, mime string) (filesystem.File, error) {
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
	return &StandaloneFile{
		filespace:  filespace,
		path:       unixTimeStr + "." + safeMIME + "." + safeBase + "." + randStr + "." + safeExt,
		mime:       mime,
		createTime: now,
		name:       base,
	}, nil
}

// OpenStandaloneFile open file and create its standalone file instance
func OpenStandaloneFile(filespace filesystem.Filespace, path string) (filesystem.File, error) {
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
	return &StandaloneFile{
		filespace:  filespace,
		path:       path,
		createTime: createTime,
		mime:       mime,
		name:       parts[2],
	}, nil
}

// Filespace return filespace of the StandaloneFile
func (f *StandaloneFile) Filespace() filesystem.Filespace {
	return f.filespace
}

// Path return a file path
func (f *StandaloneFile) Path() string {
	return f.path
}

// IsExist return true if file exists
func (f *StandaloneFile) IsExist() bool {
	return f.filespace.IsExist(f.path)
}

// IsFile return true if node is a file
func (f *StandaloneFile) IsFile() bool {
	return f.filespace.IsExist(f.path)
}

// IsDir return true if node is a directory
func (f *StandaloneFile) IsDir() bool {
	return false
}

// ReadFile read file by path
func (f *StandaloneFile) ReadFile() ([]byte, error) {
	return f.filespace.ReadFile(f.path)
}

// WriteFile write file by path
func (f *StandaloneFile) WriteFile(data []byte, perm os.FileMode) error {
	return f.filespace.WriteFile(f.path, data, perm)
}

// Reader return a file node reader
func (f *StandaloneFile) Reader() (filesystem.Reader, error) {
	return f.filespace.Reader(f.path)
}

// Writer return a file node writer
func (f *StandaloneFile) Writer() (filesystem.Writer, error) {
	return f.filespace.Writer(f.path)
}

// Remove delete node by path
func (f *StandaloneFile) Remove() error {
	return f.filespace.Remove(f.path)
}

// MIME return file MIME type
func (f *StandaloneFile) MIME() string {
	return f.mime
}

// CreateTime return file create time
func (f *StandaloneFile) CreateTime() time.Time {
	return f.createTime
}

// Name return file name
func (f *StandaloneFile) Name() string {
	return f.name
}

// Is handler to remove file
func (f *StandaloneFile) removeSignal(interface{}) error {
	return f.Remove()
}

// DeferOn connect file with scope. Remove it on event by eventID
func (f *StandaloneFile) DeferOn(scope app.EventScope, eventID int) {
	scope.On(eventID, f.removeSignal)
}
