package history

import (
	"encoding/json"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/varutil"
	"os"
	"sort"
	"strings"
	"time"
)

type History struct {
	path     string
	Timeline Timeline `json:"timeline"`
}

type Record struct {
	Time int64       `json:"time"`
	Type string      `json:"type"`
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Timeline []Record

func (a Timeline) Len() int           { return len(a) }
func (a Timeline) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Timeline) Less(i, j int) bool { return a[i].Time < a[j].Time }

func NewHistory(path string) *History {
	varutil.FixDirPath(&path)
	filesystem.MkdirAll(path)
	h := &History{
		path:     path,
		Timeline: []Record{},
	}
	return h
}

func NewRecord(recordType, name string, data interface{}) *Record {
	return &Record{
		Time: time.Now().Unix(),
		Type: recordType,
		Name: name,
		Data: data,
	}
}

func (h *History) Read() error {
	loop := filesystem.DirLoop{
		OnFile: h.loadFileFactory(),
		OnDir:  nil,
		Filter: jsonFilter,
	}
	if err := loop.Run(h.path); err != nil {
		return err
	}
	sort.Sort(h.Timeline)
	return nil
}

func (h *History) Add(r *Record) error {
	path := h.path + r.Name + "." + r.Type + ".json"
	if err := varutil.WriteJson(path, r); err != nil {
		return err
	}
	h.Timeline = append(h.Timeline, *r)
	return nil
}

func (h *History) loadFileFactory() func(os.FileInfo, string) error {
	return func(f os.FileInfo, p string) error {
		record := Record{}
		file, err := os.Open(h.path + p)
		if err != nil {
			return err
		}
		defer file.Close()

		jsonParser := json.NewDecoder(file)
		if err = jsonParser.Decode(&record); err != nil {
			return err
		}

		h.Timeline = append(h.Timeline, record)
		return nil
	}
}

func jsonFilter(file os.FileInfo, path string) bool {
	return strings.HasSuffix(path, ".json")
}
