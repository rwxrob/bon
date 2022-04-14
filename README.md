# 🌳 Go Bonzai™ Utility Branch

*🚧 This is only a placeholder for the moment. 🚧*

[![GoDoc](https://godoc.org/github.com/rwxrob/bon?status.svg)](https://godoc.org/github.com/rwxrob/bon)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

The `bon` branch is for everything to help with development, use, and
discovery of Bonzai branches and leaf commands (`z bon init foo`, etc.)

## Install

This command can be installed as a standalone program or composed into a
Bonzai command tree.

Standalone

```
go install github.com/rwxrob/bon/bon@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/bon"
)

var Cmd = &Z.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{help.Cmd, bon.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C bon bon
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
