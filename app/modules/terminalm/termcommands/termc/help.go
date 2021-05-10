package termc

import (
	"sort"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/terminalm/termservices"
	"github.com/goatcms/goatcore/app/terminal"
	"github.com/goatcms/goatcore/app/terminal/termformatter"
)

// RunHelp show help
func RunHelp(a app.App, ctx app.IOContext) (err error) {
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
	commands := terminal.MergeCommands(term, Commands()...)
	info := deps.Info
	// header
	io.Out().Printf("%s %s\n", a.Name(), a.Version().String())
	if info.License().Company() != "" {
		license := info.License()
		io.Out().Printf("Develop by @%s", license.Company())
		if !license.Start().IsZero() {
			io.Out().Printf(" %v", license.Start().Year())
			if !license.End().IsZero() {
				io.Out().Printf("-%v", license.End().Year())
			}
		}
		if license.Type() != "" {
			// type is like: all rights reserved
			io.Out().Printf(" %s", license.Type())
		}
		io.Out().Printf("\n")
	}
	if info.PoweredBy() != "" {
		io.Out().Printf("%s\n", info.PoweredBy())
	}
	if info.Description() != "" {
		io.Out().Printf("\n%s\n\n", info.Description())
	}
	// commands
	commandNames := commands.CommandNames()
	sort.Strings(commandNames)
	max := maxLength(commandNames)
	if len(commandNames) > 0 {
		io.Out().Printf("Commands:\n")
		formatter := termformatter.NewBlockFormatter(io.Out(), lineWidth,
			termformatter.NewBlockDef(max+2, termformatter.ToRight, termformatter.ToRight),
			termformatter.NewBlockDef(2, termformatter.ToRight, termformatter.ToRight),
			termformatter.NewBlockDef(lineWidth-(max+4), termformatter.Justify, termformatter.ToLeft),
		)
		for _, commandName := range commandNames {
			command := commands.Command(commandName)
			formatter.PrintBlocks(commandName, "", command.Help())
		}
	}
	// arguments
	argumentNames := term.ArgumentNames()
	max = maxLength(argumentNames)
	if len(argumentNames) > 0 {
		io.Out().Printf("nArguments:\n")
		formatter := termformatter.NewBlockFormatter(io.Out(), lineWidth,
			termformatter.NewBlockDef(max+4, termformatter.ToRight, termformatter.ToRight),
			termformatter.NewBlockDef(2, termformatter.ToRight, termformatter.ToRight),
			termformatter.NewBlockDef(lineWidth-(max+6), termformatter.Justify, termformatter.ToLeft),
		)
		for _, argumentName := range argumentNames {
			argument := term.Argument(argumentName)
			formatter.PrintBlocks("--"+argumentName, "", argument.Help())
		}
	}
	RunHealth(a, ctx)
	return nil
}
