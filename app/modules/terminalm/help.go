package terminalm

import (
	"sort"
	"strings"

	"github.com/goatcms/goatcore/app"
)

// HelpComamnd show help message
func HelpComamnd(a app.App, ctx app.IOContext) (err error) {
	var (
		deps struct {
			Name         string    `app:"AppName"`
			Version      string    `app:"AppVersion"`
			Welcome      string    `app:"?AppWelcome"`
			Company      string    `app:"?AppCompany"`
			GoatVersion  string    `engine:"GoatVersion"`
			CommandName  string    `argument:"?$1"`
			CommandScope app.Scope `dependency:"CommandScope"`
		}
		io = ctx.IO()
	)
	if err = a.DependencyProvider().InjectTo(&deps); err != nil {
		return err
	}
	// header
	io.Out().Printf("%s %s\n", deps.Name, deps.Version)
	if deps.Company != "" {
		io.Out().Printf("Develop by @%s all rights reserved\n", deps.Company)
	}
	io.Out().Printf("Powered by GoatCore %s (%s)\n", deps.GoatVersion, "https://github.com/goatcms/goatcore")
	if deps.Welcome != "" {
		io.Out().Printf("\n%s\n", deps.Welcome)
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
				io.Out().Printf("\nCommands:\n")
				isFirstCommand = false
			}
			helpStr, err := deps.CommandScope.Get(key)
			if err != nil {
				return err
			}
			io.Out().Printf("%s  %s\n", fixSpace(key[len(commandPrefix):], maxLength), helpStr)
		}
	}
	isFirstArgument := true
	for _, key := range keys {
		if strings.HasPrefix(key, argumentPrefix) {
			if isFirstArgument {
				io.Out().Printf("\nArguments:\n")
				isFirstArgument = false
			}
			helpStr, err := deps.CommandScope.Get(key)
			if err != nil {
				return err
			}
			io.Out().Printf("%11s  %s\n", fixSpace(key[len(argumentPrefix):], maxLength), helpStr)
		}
	}
	HealthComamnd(a, ctx)
	return nil
}
