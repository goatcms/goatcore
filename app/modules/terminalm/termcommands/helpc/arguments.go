package helpc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/terminalm/termcommands"
	"github.com/goatcms/goatcore/app/terminal/termformatter"
	"github.com/goatcms/goatcore/varutil"
)

// printArguments show command help
func printArguments(out app.Output, commandName string, arguments app.TerminalArguments, mains []string) (err error) {
	var (
		lastInnerPipeline string
	)
	lineWidth := termcommands.LineWidth
	args := arguments.ArgumentNames()
	argumentNames := arguments.ArgumentNames()
	sort.Strings(argumentNames)
	argsMax := maxLength(args)
	argsSpace := argsMax + 5
	argsFormatter := termformatter.NewBlockFormatter(
		out,
		lineWidth,
		termformatter.NewBlockDef(argsSpace, termformatter.ToRight, termformatter.ToRight),
		termformatter.NewBlockDef(1, termformatter.ToRight, termformatter.ToRight),
		termformatter.NewBlockDef(lineWidth-argsSpace-1, termformatter.Justify, termformatter.ToLeft),
	)
	attrFormatter := termformatter.NewBlockFormatter(
		out,
		lineWidth,
		termformatter.NewBlockDef(argsSpace+1, termformatter.ToRight, termformatter.ToRight),
		termformatter.NewBlockDef(lineWidth-argsSpace-1, termformatter.Justify, termformatter.ToLeft),
	)
	// display
	for _, name := range mains {
		argument := arguments.Argument(name)
		if argument == nil {
			panic(fmt.Errorf("unknow argument %s", name))
		}
		argsFormatter.PrintBlocks(name, "", argument.Help())
		printArgumentOptions(attrFormatter, argument)
	}
	for _, name := range argumentNames {
		if varutil.IsArrContainStr(mains, name) {
			continue
		}
		argument := arguments.Argument(name)
		if argument == nil {
			panic(fmt.Errorf("unknow argument %s", name))
		}
		argsFormatter.PrintBlocks(fmt.Sprintf("--%s", name), "", argument.Help())
		printArgumentOptions(attrFormatter, argument)
	}
	for _, name := range argumentNames {
		argument := arguments.Argument(name)
		commands := argument.CommandNames()
		if len(commands) == 0 {
			continue
		}
		out.Printf("\nList of commands for %s argument pipeline:\n", name)
		printCommandList(out, argument)
		lastInnerPipeline = name
	}
	if lastInnerPipeline != "" {
		out.Printf("\nRemember:\n")
		out.Printf(" You can display a pipeline command details by\n")
		out.Printf("   help [command name]/[pipeline argument]/[command name]\n")
		out.Printf(" like:\n")
		commands := arguments.Argument(lastInnerPipeline).CommandNames()
		sort.Strings(commands)
		out.Printf("   help %s/%s/%s\n", commandName, lastInnerPipeline, commands[0])
	}
	return nil
}

func printArgumentOptions(formatter termformatter.BlockFormatter, arg app.TerminalArgument) {
	flags := []string{}
	if arg.Required() {
		flags = append(flags, "Required")
	} else {
		flags = append(flags, "Optional")
	}
	switch arg.Type() {
	case app.TerminalBoolArgument:
		flags = append(flags, "Bool")
	case app.TerminalEmailArgument:
		flags = append(flags, "Email")
	case app.TerminalFloatArgument:
		flags = append(flags, "Float")
	case app.TerminalIntArgument:
		flags = append(flags, "Int")
	case app.TerminalOtherArgument:
		flags = append(flags, "Other")
	case app.TerminalPathArgument:
		flags = append(flags, "Path")
	case app.TerminalPIPArgument:
		flags = append(flags, "PIP")
	case app.TerminalTextArgument:
		flags = append(flags, "Text")
	case app.TerminalURLArgument:
		flags = append(flags, "URL")
	case app.TerminalUndefinedArgument:
		flags = append(flags, "MIX")
	}
	formatter.PrintBlocks("", fmt.Sprintf("(%s)", strings.Join(flags, " ")))
}
