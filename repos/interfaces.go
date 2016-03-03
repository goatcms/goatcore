package repos

type Repository interface {
	Clone(string) error
	Uninit() error
}
