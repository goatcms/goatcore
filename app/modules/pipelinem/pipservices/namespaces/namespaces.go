package namespaces

import "github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"

// Namasepaces is namespace storage for scope
type Namasepaces struct {
	task, lock string
}

// NewNamespaces create new Namasepaces instance
func NewNamespaces(params pipservices.NamasepacesParams) pipservices.Namespaces {
	return pipservices.Namespaces(Namasepaces{
		task: params.Task,
		lock: params.Lock,
	})
}

// NewSubNamespaces create new Namasepaces instance into any other
func NewSubNamespaces(parent pipservices.Namespaces, params pipservices.NamasepacesParams) pipservices.Namespaces {
	if parent.Lock() != "" {
		if params.Lock != "" {
			params.Lock = parent.Lock() + ":" + params.Lock
		} else {
			params.Lock = parent.Lock()
		}
	}
	if parent.Task() != "" {
		if params.Task != "" {
			params.Task = parent.Task() + ":" + params.Task
		} else {
			params.Task = parent.Task()
		}
	}
	return pipservices.Namespaces(Namasepaces{
		task: params.Task,
		lock: params.Lock,
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
