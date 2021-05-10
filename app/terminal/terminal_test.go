package terminal

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil"
)

func TestTerminalCommandStory(t *testing.T) {
	t.Parallel()
	term := NewTerminalManager()
	term.SetCommand(NewCommand(CommandParams{
		Name:     "firstCommand",
		Callback: func(app.App, app.IOContext) error { return nil },
	}))
	term.SetCommand(NewCommand(CommandParams{
		Name:     "secondCommand",
		Callback: func(app.App, app.IOContext) error { return nil },
	}))
	commandNames := term.CommandNames()
	if len(commandNames) != 2 {
		t.Errorf("expected 2 elements and take: %v", commandNames)
		return
	}
	if !varutil.IsArrContainStr(commandNames, "firstCommand") {
		t.Errorf("expected firstCommand")
		return
	}
	if !varutil.IsArrContainStr(commandNames, "secondCommand") {
		t.Errorf("expected secondCommand")
		return
	}
	firstCommand := term.Command("firstCommand")
	if firstCommand.Name() != "firstCommand" {
		t.Errorf(`terminal.Command("firstCommand") expected firstCommand as result`)
		return
	}
	if firstCommand.Callback() == nil {
		t.Errorf(`terminal.Command("firstCommand") callback is expected`)
		return
	}
}

func TestTerminalArgumentsStory(t *testing.T) {
	t.Parallel()
	terminal := NewTerminalManager()
	terminal.SetArgument(NewArgument(ArgumentParams{
		Name:     "shell",
		Help:     "somme command to run in shell",
		Required: true,
		Type:     app.TerminalPIPArgument,
		Commands: NewCommands(
			NewCommand(CommandParams{
				Name: "ls",
				Help: "list directory",
				Arguments: NewArguments(
					NewArgument(ArgumentParams{
						Name: "all",
						Help: "do not ignore entries starting with .",
						Type: app.TerminalBoolArgument,
					}),
				),
				Callback: app.NilCommandCallback,
			}),
		),
	}))
	if terminal.Argument("shell").Command("ls").Argument("all").Name() != "all" {
		t.Errorf("expected all argument into terminal '%v'", terminal)
	}
	terminal.SetCommand(NewCommand(CommandParams{
		Name:     "exe",
		Callback: app.NilCommandCallback,
		Help:     "",
		Arguments: NewArguments(
			NewArgument(ArgumentParams{
				Name: "file",
				Help: "file path to exec",
				Type: app.TerminalTextArgument,
			}),
		),
	}))
	exeCommand := terminal.Command("exe")
	fileArg := exeCommand.Argument("file")
	if fileArg == nil {
		t.Errorf(`terminal.Command("exe").Argument("file") expected argument`)
		return
	}
}
