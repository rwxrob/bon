package bon

import (
	_ "embed"
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

//go:embed desc/help.md
var helpdoc string

var Cmd = &Z.Cmd{
	Name:        `bon`,
	Aliases:     []string{`bonzai`},
	Summary:     `bonzai helper utility`,
	Description: helpdoc,
	Version:     `v0.1.1`,
	Copyright:   `Copyright 2022 Rob Muhlestein`,
	Site:        `rwxrob.tv`,
	Source:      `git@github.com:rwxrob/bon.git`,
	Issues:      `https://github.com/rwxrob/bon/issues`,
	License:     `Apache-2.0`,
	Commands:    []*Z.Cmd{help.Cmd, commandCmd},
}

//go:embed desc/command.md
var commanddoc string

//go:embed tmpl/command.tmpl
var commandtmpl string

var commandCmd = &Z.Cmd{
	Name:        `command`,
	Aliases:     []string{`cmd`},
	Summary:     `print code for new bonzai.Cmd`,
	Description: commanddoc,
	Commands:    []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, _ ...string) error {
		fmt.Print(commandtmpl)
		return nil
	},
}
