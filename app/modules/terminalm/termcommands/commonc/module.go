package commonc

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/terminal"
)

func CWDArgument() app.TerminalArgument {
	return terminal.NewArgument(terminal.ArgumentParams{
		Name: `cwd`,
		Help: `The argument set path to current working directory. The CWD word is abbreviation of "Current Working Directory" phrase.`,
		Type: app.TerminalPathArgument,
	})
}

func Arguments() []app.TerminalArgument {
	return []app.TerminalArgument{
		CWDArgument(),
	}
}
