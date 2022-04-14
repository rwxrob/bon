// Copyright 2022 Robert Muhlestein.
// SPDX-License-Identifier: Apache-2.0

// Package bon provides the Bonzai command branch of the same name.
package bon

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

// main branch
var Cmd = &Z.Cmd{

	Name:      `bon`,
	Summary:   `official Bonzai™ helper utility`,
	Version:   `v0.0.0`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Site:      `rwxrob.tv`,
	Source:    `git@github.com:rwxrob/bon.git`,
	Issues:    `github.com/rwxrob/bon/issues`,

	Commands: []*Z.Cmd{
		help.Cmd, conf.Cmd, vars.Cmd, // common, first is default
	},

	Description: `
	  {{define "long" -}} 
		Here is something
		that spans multiple
		lines that would otherwise be too long for a single action.
		{{- end}}

		The **{{.Name}}** branch is for everything to help with development,
		use, and discovery of Bonzai branches and leaf commands
		({{- template "long" "" | pre -}}).`,
}
