package helpc

import (
	"sort"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/app/terminal"
)

// runOverview show main help page (commands list and tooltips)
func runOverview(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			Info     commservices.Info     `dependency:"CommonInfo"`
			Terminal termservices.Terminal `dependency:"TerminalService"`
		}
		io   = ctx.IO()
		term = a.Terminal()
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	RunHeader(a, ctx)
	// commands
	commands := terminal.MergeCommands(term, Commands()...)
	commandNames := commands.CommandNames()
	sort.Strings(commandNames)
	io.Out().Printf("\vCommands:\n")
	printCommandList(io.Out(), commands)
	// app lvl arguments
	argumentNames := term.ArgumentNames()
	if len(argumentNames) > 0 {
		io.Out().Printf("\nArguments:\n")
	}
	printArguments(io.Out(), "", term, []string{})
	io.Out().Printf("\nRemember:\n")
	io.Out().Printf(" You can display a command details by 'help [command name]'.\n")
	io.Out().Printf(" Enjoy 'help %s'.\n\n", commandNames[0])
	RunHealth(a, ctx)
	return nil
}
