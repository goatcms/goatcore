package helpc

import (
	"fmt"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/terminalm/termcommands"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/app/terminal"
	"github.com/goatcms/goatcore/app/terminal/termformatter"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// runCommandHelp show command help
func runCommandHelp(a app.App, ctx app.IOContext, commandPath string) (err error) {
	var (
		deps struct {
			Info     commservices.Info     `dependency:"CommonInfo"`
			Terminal termservices.Terminal `dependency:"TerminalService"`
		}
		io          = ctx.IO()
		term        = a.Terminal()
		lineWidth   = termcommands.LineWidth
		command     app.TerminalCommand
		commandName string
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	// load command
	commands := terminal.MergeCommands(term, Commands()...)
	if commandName, command, err = loadCommandByPath(commands, commandPath); err != nil {
		return err
	}
	mainArgs := command.MainArguments()
	argumentNames := command.ArgumentNames()
	// formatters
	formatter := termformatter.NewBlockFormatter(io.Out(), lineWidth,
		termformatter.NewBlockDef(2, termformatter.Justify, termformatter.ToLeft),
		termformatter.NewBlockDef(lineWidth-2, termformatter.Justify, termformatter.ToLeft),
	)
	// display
	RunHeader(a, ctx)
	io.Out().Printf("\nSyntax:\n")
	mainArgsStr := ""
	if len(mainArgs) > 0 {
		mainArgsStr = strings.Join(mainArgs, " ") + " "
	}
	io.Out().Printf("  %s %s--arg=value -- ignored\n", commandName, mainArgsStr)
	if len(argumentNames) > 0 {
		io.Out().Printf("\nArguments:\n")
	}
	printArguments(io.Out(), commandName, command, mainArgs)
	io.Out().Printf("\nDescription:\n")
	formatter.PrintBlocks("", command.Help())
	io.Out().Printf("\n")
	return nil
}

func loadCommandByPath(commands app.TerminalCommands, commandPath string) (commandName string, command app.TerminalCommand, err error) {
	path := strings.Split(commandPath, "/")
	if len(path)%2 != 1 {
		return "", nil, goaterr.Errorf("%s it is pointer for argument (expected command pointer)", commandPath)
	}
	i := 1
	current := commands
	for ; i < len(path); i += 2 {
		fmt.Printf("%d", i)
		if command = current.Command(path[i-1]); command == nil {
			return "", nil, goaterr.Errorf("unknow command %s in %s path", path[i-1], commandPath)
		}
		if current = command.Argument(path[i]); current == nil {
			return "", nil, goaterr.Errorf("unknow argument %s in %s path", path[i], commandPath)
		}
	}
	commandName = path[i-1]
	if command = current.Command(commandName); command == nil {
		return "", nil, goaterr.Errorf("unknow command %s in %s path", path[i-1], commandPath)
	}
	return
}
