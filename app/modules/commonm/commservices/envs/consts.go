package envs

import "regexp"

var (
	envKey         = "sandboxes_enviroment"
	envNamePattern = regexp.MustCompile("^[a-zA-Z]+([_a-zA-Z]+)?$")
)
