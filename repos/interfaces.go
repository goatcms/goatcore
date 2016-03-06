package repos

type Repository interface {
	Clone(string) error
	Checkout(string) error
	Pull() error
	Uninit() error
}
	
