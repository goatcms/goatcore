// Package repos is experimental. You must be careful. It is still development
package repos

// NewRepository create repository by path
func NewRepository(path string) Repository {
	return NewGitRepository(path)
}
