package modules

const (
	// TerminalService is a key for Terminal service
	TerminalService = "TerminalService"
)

// Terminal is global terminal interface
type Terminal interface {
	RunLoop() (err error)
	RunString(s string) (err error)
	RunCommand(args []string) (err error)
}
