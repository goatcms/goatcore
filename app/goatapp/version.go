package goatapp

import (
	"fmt"

	"github.com/goatcms/goatcore/app"
)

var (
	NilVersion = NewVersion(0, 0, 0, "-dev")
)

type Version struct {
	major  int
	minor  int
	path   int
	suffix string
}

func NewVersion(major int, minor int, path int, suffix string) app.Version {
	return Version{
		major:  major,
		minor:  minor,
		path:   path,
		suffix: suffix,
	}
}

func (v Version) Major() int {
	return v.major
}

func (v Version) Minor() int {
	return v.minor
}

func (v Version) Path() int {
	return v.path
}

func (v Version) Suffix() string {
	return v.suffix
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d%s", v.major, v.minor, v.path, v.suffix)
}
