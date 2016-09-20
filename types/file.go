package types

// File is default file interface (use for many filesources)
type File interface {
	Filespace() string
	Path() string
}
