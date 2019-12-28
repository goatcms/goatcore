package pipc

import (
	"regexp"

	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices/namespaces"
)

const (
	cutset = "\n\t "
)

var (
	// namePattern define correct name
	namePattern = regexp.MustCompile("^[a-zA-Z_]+[a-zA-Z0-9_]*$")
	// defaultNamespace is default namespace for main task
	defaultNamespace = namespaces.NewNamespaces("main", "")
)
