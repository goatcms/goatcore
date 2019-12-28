package namespaces

import "github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"

// Namasepaces is namespace storage for scope
type Namasepaces struct {
	task, lock string
}

// NewNamespaces create new Namasepaces instance
func NewNamespaces(task, lock string) pipservices.Namespaces {
	return pipservices.Namespaces(Namasepaces{
		task: task,
		lock: lock,
	})
}

// Task return task namespace
func (namasepaces Namasepaces) Task() (value string) {
	return namasepaces.task
}

// Lock return lock namespace
func (namasepaces Namasepaces) Lock() (value string) {
	return namasepaces.lock
}
