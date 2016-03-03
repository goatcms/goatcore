package repos

func NewRepository(path string) Repository {
	return NewGitRepository(path)
}
