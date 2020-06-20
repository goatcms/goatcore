package app

// CommandCallback is function call to run user command
type CommandCallback func(App, IOContext) (err error)

// HealthCheckerCallback is function to check application health
type HealthCheckerCallback func(App, Scope) (msg string, err error)
