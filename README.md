# 🌳 Go Bonzai Utility Branch

[![GoDoc](https://godoc.org/github.com/rwxrob/bon?status.svg)](https://godoc.org/github.com/rwxrob/bon)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

This `bon` [bonzai](https://github.com/rwxrob/bonzai) branch contains helpers for creating and managing stateful command branches in Go. 

## Install

You can just download from the [releases page](https://github.com/rwxrob/bon/releases).

```
curl -L https://github.com/rwxrob/bon/releases/latest/download/bon-linux-amd64 -o ~/.local/bin/bon
curl -L https://github.com/rwxrob/bon/releases/latest/download/bon-darwin-amd64 -o ~/.local/bin/bon
curl -L https://github.com/rwxrob/bon/releases/latest/download/bon-darwin-arm64 -o ~/.local/bin/bon
curl -L https://github.com/rwxrob/bon/releases/latest/download/bon-windows-amd64 -o ~/.local/bin/bon

```

Or with `go`:

```
go install github.com/rwxrob/bon/cmd/bon@latest
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

## Command Line Usage

```
bon help
```

## Build and Release Instructions

Building workflow uses the [`good`](https://github.com/rwxrob/good) Go helper tool (often composited into bonzai personal command trees (`z go`):

```
cd cmd/bon
good build
gh release create
gh release upload TAGVER build/*
```
