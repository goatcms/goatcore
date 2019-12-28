package pipservices

// Runner run command pipeline
type Runner interface {
	Run(pip Pip) (err error)
}
