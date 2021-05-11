package pipc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/terminal"
)

func ClearCommand() app.TerminalCommand {
	return terminal.NewCommand(terminal.CommandParams{
		Callback: Clear,
		Help:     "clear current pipeline context",
		Name:     "pip:clear",
	})
}

func RunCommand() app.TerminalCommand {
	return terminal.NewCommand(terminal.CommandParams{
		Callback: Run,
		Help:     "Run code pipeline",
		Name:     "pip:run",
		Arguments: terminal.NewArguments([]app.TerminalArgument{
			terminal.NewArgument(terminal.ArgumentParams{
				Help:     "unique execution block name",
				Name:     "name",
				Required: true,
				Type:     app.TerminalTextArgument,
			}),
			terminal.NewArgument(terminal.ArgumentParams{
				Help: "identyfier of sandbox to run body within. For example terminal or container:image (container:alpine:3)",
				Name: "sandbox",
				Type: app.TerminalTextArgument,
			}),
			terminal.NewArgument(terminal.ArgumentParams{
				Help: "contains commands/content to run",
				Name: "body",
				Type: app.TerminalPIPArgument,
			}),
			terminal.NewArgument(terminal.ArgumentParams{
				Help: "contains list of execution block names to wait for (names are separate with comma ',')",
				Name: "wait",
				Type: app.TerminalOtherArgument,
			}),
			terminal.NewArgument(terminal.ArgumentParams{
				Help: "contains list of resources to read and write lock for this task. Resource's names are separate with comma ','. It lock current namespace resources by default. Add '@' commat at the begin of resource name to lock global resources. Global pool is create per application.",
				Name: "wlock",
				Type: app.TerminalOtherArgument,
			}),
			terminal.NewArgument(terminal.ArgumentParams{
				Help: "contains list of resources to lock only for read. See lock attribute.",
				Name: "rlock",
				Type: app.TerminalOtherArgument,
			}),
		}...),
	})
}

func TryCommand() app.TerminalCommand {
	return terminal.NewCommand(terminal.CommandParams{
		Callback: Try,
		Help:     "Run code pipelines. Work similar to 'try ... catch' used by many programming languages ",
		Name:     "pip:try",
		Arguments: terminal.NewArguments([]app.TerminalArgument{
			terminal.NewArgument(terminal.ArgumentParams{
				Help:     "unique execution block name",
				Name:     "name",
				Required: true,
				Type:     app.TerminalTextArgument,
			}),
			terminal.NewArgument(terminal.ArgumentParams{
				Help: "main pipeline to run",
				Name: "body",
				Type: app.TerminalPIPArgument,
			}),
			terminal.NewArgument(terminal.ArgumentParams{
				Help: "run pipeline when body status is success (use self sandbox)",
				Name: "success",
				Type: app.TerminalPIPArgument,
			}),
			terminal.NewArgument(terminal.ArgumentParams{
				Help: "run pipeline when body status is fail (use self sandbox)",
				Name: "fail",
				Type: app.TerminalPIPArgument,
			}),
			terminal.NewArgument(terminal.ArgumentParams{
				Help: "run after success/fail pipeline (use self sandbox)",
				Name: "finally",
				Type: app.TerminalPIPArgument,
			}),
		}...),
	})
}

func LogsCommand() app.TerminalCommand {
	return terminal.NewCommand(terminal.CommandParams{
		Callback: Logs,
		Help:     "show pipeline's execution logs",
		Name:     "pip:logs",
	})
}

func SummaryCommand() app.TerminalCommand {
	return terminal.NewCommand(terminal.CommandParams{
		Callback: Summary,
		Help:     "show pipeline's execution summary logs",
		Name:     "pip:summary",
	})
}

func Commands() []app.TerminalCommand {
	return []app.TerminalCommand{
		ClearCommand(),
		LogsCommand(),
		RunCommand(),
		SummaryCommand(),
		TryCommand(),
	}
}
