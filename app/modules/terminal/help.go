package terminal

import (
	"sort"
	"strings"

	"github.com/goatcms/goatcore/app"
)

// PrintHelp show help message
func PrintHelp(a app.App) (err error) {
	var deps struct {
		Name         string    `app:"AppName"`
		Version      string    `app:"AppVersion"`
		Welcome      string    `app:"?AppWelcome"`
		Company      string    `app:"?AppCompany"`
		GoatVersion  string    `engine:"GoatVersion"`
		CommandName  string    `argument:"?$1"`
		CommandScope app.Scope `dependency:"CommandScope"`

		Input  app.Input  `dependency:"InputService"`
		Output app.Output `dependency:"OutputService"`
	}
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	// header
	deps.Output.Printf("%s %s\n", deps.Name, deps.Version)
	if deps.Company != "" {
		deps.Output.Printf("Develop by @%s all rights reserved\n", deps.Company)
	}
	deps.Output.Printf("Powered by GoatCore %s (%s)\n", deps.GoatVersion, "https://github.com/goatcms/goatcore")
	if deps.Welcome != "" {
		deps.Output.Printf("\n%s\n", deps.Welcome)
	}
	// content
	keys, err := deps.CommandScope.Keys()
	if err != nil {
		return err
	}
	isFirstCommand := true
	maxLength := 0
	for _, key := range keys {
		if len(key) > maxLength {
			maxLength = len(key)
		}
	}
	maxLength = maxLength - len(commandPrefix) + 1
	sort.Strings(keys)
	for _, key := range keys {
		if strings.HasPrefix(key, commandPrefix) {
			if isFirstCommand {
				deps.Output.Printf("\nCommands:\n")
				isFirstCommand = false
			}
			helpStr, err := deps.CommandScope.Get(key)
			if err != nil {
				return err
			}
			deps.Output.Printf("%s  %s\n", fixSpace(key[len(commandPrefix):], maxLength), helpStr)
		}
	}
	isFirstArgument := true
	for _, key := range keys {
		if strings.HasPrefix(key, argumentPrefix) {
			if isFirstArgument {
				deps.Output.Printf("\nArguments:\n")
				isFirstArgument = false
			}
			helpStr, err := deps.CommandScope.Get(key)
			if err != nil {
				return err
			}
			deps.Output.Printf("%11s  %s\n", fixSpace(key[len(argumentPrefix):], maxLength), helpStr)
		}
	}
	deps.Output.Printf("\n")
	return nil
}
