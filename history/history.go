package history

import (
	"encoding/json"
	"github.com/goatcms/goat-core/filesystem"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

type History struct {
	path     string
	Timeline HistoryTimeline `json:"timeline"`
}

type HistoryRecord struct {
	Time string      `json:"time"`
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type HistoryTimeline []HistoryRecord

func (a HistoryTimeline) Len() int           { return len(a) }
func (a HistoryTimeline) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a HistoryTimeline) Less(i, j int) bool { return a[i].Time < a[j].Time }

func NewHistory(path string) *History {
	h := &History{
		path:     path,
		Timeline: []HistoryRecord{},
	}

	if !strings.HasSuffix(h.path, "/") {
		h.path = h.path + "/"
	}

	return h
}

func NewHistoryRecord(name string, data *interface{}) *HistoryRecord {
	r := &HistoryRecord{
		Time: strconv.FormatInt(time.Now().Unix(), 10),
		Name: name,
		Data: *data,
	}
	return r
}

func (h *History) Load() error {
	loop := filesystem.DirLoop{
		OnFile: h.loadFileFactory(),
		OnDir:  nil,
		Filter: jsonFilter,
	}
	if err := loop.Run(h.path); err != nil {
		return err
	}
	sort.Sort(HistoryTimeline(h.Timeline))
	return nil
}

func (h *History) Add(r *HistoryRecord) error {
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	path := h.path + r.Time + "." + r.Name + ".json"
	err = ioutil.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	h.Timeline = append(h.Timeline, *r)
	return nil
}

func (h *History) loadFileFactory() func(os.FileInfo, string) error {
	return func(f os.FileInfo, p string) error {
		record := HistoryRecord{}
		elements := strings.Split(path.Base(p), ".")
		if len(elements) < 3 {
			record.Time = "0"
			record.Name = elements[0]
		} else {
			record.Time = elements[0]
			record.Name = elements[1]
		}

		file, err := os.Open(h.path + p)
		if err != nil {
			return err
		}
		defer file.Close()

		jsonParser := json.NewDecoder(file)
		if err = jsonParser.Decode(&record.Data); err != nil {
			return err
		}

		h.Timeline = append(h.Timeline, record)
		return nil
	}
}

func jsonFilter(file os.FileInfo, path string) bool {
	return strings.HasSuffix(path, ".json")
}
