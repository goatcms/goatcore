package helpc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/terminalm/termcommands"
	"github.com/goatcms/goatcore/app/terminal/termformatter"
)

// printCommandList show commands list
func printCommandList(out app.Output, commands app.TerminalCommands) (err error) {
	var (
		lineWidth = termcommands.LineWidth
	)
	commandNames := commands.CommandNames()
	sort.Strings(commandNames)
	commandMax := maxLength(commandNames)
	commandScpace := commandMax + 4
	formatter := termformatter.NewBlockFormatter(out, lineWidth,
		termformatter.NewBlockDef(commandMax+2, termformatter.ToRight, termformatter.ToRight),
		termformatter.NewBlockDef(2, termformatter.ToRight, termformatter.ToRight),
		termformatter.NewBlockDef(lineWidth-(commandScpace), termformatter.Justify, termformatter.ToLeft),
	)
	for _, commandName := range commandNames {
		// command details
		command := commands.Command(commandName)
		mainArgs := command.MainArguments()
		fargs := make([]string, len(mainArgs))
		for i, name := range mainArgs {
			arg := command.Argument(name)
			if arg.Required() {
				fargs[i] = name
			} else {
				fargs[i] = fmt.Sprintf("[%s]", name)
			}
		}
		formatedMainArgs := ""
		if len(mainArgs) != 0 {
			formatedMainArgs = fmt.Sprintf("[%s] ", strings.Join(fargs, ", "))
		}
		formatter.PrintBlocks(commandName, "", formatedMainArgs+command.Help())
	}
	return nil
}
