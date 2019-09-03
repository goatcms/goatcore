package scopedefer

import (
	"strconv"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// FileDefer provide defer api to remove multi files on remvoe
type FileDefer struct {
	files []filesystem.File
}

// Remove delete files
func (d *FileDefer) Remove(interface{}) error {
	var errors []error
	for _, file := range d.files {
		if err := file.Remove(); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		// TODO: Add suport to containing multiple errors
		return goaterr.Errorf("%v", errors)
	}
	return nil
}

// Add append file to remove on error
func (d *FileDefer) Add(file filesystem.File) {
	d.files = append(d.files, file)
}

// RemoveOn bind scope event  with file to remove on error
func RemoveOn(scope app.Scope, eventID int, file filesystem.File) error {
	var def *FileDefer
	insKey := "_scopedefer.FileDefer:" + strconv.Itoa(eventID)
	fileDeferIns, err := scope.Get(insKey)
	if err != nil || fileDeferIns == nil {
		def = &FileDefer{
			files: []filesystem.File{file},
		}
		scope.On(eventID, def.Remove)
	} else {
		def = fileDeferIns.(*FileDefer)
		def.Add(file)
	}
	return nil
}
