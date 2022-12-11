package bon

import (
	"fmt"
	"strings"
	"text/template"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `bon`,
	Aliases:     []string{`bonzai`},
	Summary:     help.S(_bon),
	Description: help.D(_bon),
	Version:     `v0.1.2`,
	Copyright:   `Copyright 2022 Rob Muhlestein`,
	Site:        `rwxrob.tv`,
	Source:      `git@github.com:rwxrob/bon.git`,
	Issues:      `https://github.com/rwxrob/bon/issues`,
	License:     `Apache-2.0`,
	Commands:    []*Z.Cmd{help.Cmd, commandCmd},
}

var commandCmd = &Z.Cmd{
	Name:        `command`,
	Aliases:     []string{`cmd`},
	Summary:     help.S(_command),
	Description: help.D(_command),
	Commands:    []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {

		if len(args) == 0 {
			args = append(args, `foo`)
		}

		t, err := template.New("t").Parse(commandtmpl)
		if err != nil {
			return err
		}
		var buf strings.Builder
		if err := t.Execute(&buf, struct{ Name string }{args[0]}); err != nil {
			return err
		}
		fmt.Print(buf.String())
		return nil
	},
}
