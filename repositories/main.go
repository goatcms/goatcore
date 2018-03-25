package repositories

// Connector is repositories controll access interface
type Connector interface {
	// Clone clone repository to local directory
	Clone(url, version, destPath string) (repo Repository, err error)
	// Open open repository from local filesystem
	Open(path string) (repo Repository, err error)
}

// Repository is repository access interface
type Repository interface {
	// Pull update repository
	Pull() (err error)
}

// ConnectorAdapter is object support a connector for specified repository system. Like git, maven, svn etc
type ConnectorAdapter interface {
	Connector
	// IsSupportURL check if repository URL is supported
	IsSupportURL(url string) bool
	// IsSupportRepo check if local repository is supported
	IsSupportRepo(path string) bool
}
