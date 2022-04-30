// Copyright 2022 Robert Muhlestein.
// SPDX-License-Identifier: Apache-2.0

// Package bon provides the Bonzai command branch of the same name.
package bon

import (
	"fmt"
	"log"
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/fs/dir"
	"github.com/rwxrob/fs/file"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
	"gopkg.in/yaml.v2"
)

var Cmd = &Z.Cmd{

	Name:      `bon`,
	Summary:   `official Bonzai™ helper utility`,
	Version:   `v0.1.1`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Site:      `rwxrob.tv`,
	Source:    `git@github.com:rwxrob/bon.git`,
	Issues:    `github.com/rwxrob/bon/issues`,

	Commands: []*Z.Cmd{
		help.Cmd, conf.Cmd, vars.Cmd, // common, first is default
		distCmd,
	},

	Description: `
		The **{{.Name}}** branch is for everything to help with development,
		use, distribution, and discovery of Bonzai branches and leaf commands
		({{- template "long" "" | pre -}}).`,
}

func gobuild(args ...string) error {
	a := []string{`go`, `build`}
	a = append(a, args...)
	return Z.Exec(a...)
}

var godist = &Z.Cmd{
	Name:     `dist`,
	Summary:  `go distribution related commands`,
	Commands: []*Z.Cmd{godistbuild},
}

type GoBuildParams struct {
	Targets []GoBuildTarget
	O       map[string]any `yaml:",inline"`
}

type GoBuildTarget struct {
	OS   string
	Arch []string
}

var godistbuild = &Z.Cmd{
	Name:     `build`,
	Summary:  `build for for multiple architectures into dist dir`,
	Commands: []*Z.Cmd{help.Cmd},
	Description: `
			This build looks for a build.yaml file in the current directory
			and runs the build command on each building them all concurrently
			into the *dist* directory where they are ready for upload to
			GitHub as a release. Just add a README.md and run release.
	`,
	Call: func(_ *Z.Cmd, args ...string) error {
		if !file.Exists(`build.yaml`) {
			return gobuild(args...)
		}
		buf, err := os.ReadFile(`build.yaml`)
		if err != nil {
			return err
		}
		p := new(GoBuildParams)
		if err := yaml.Unmarshal(buf, p); err != nil {
			return err
		}
		os.RemoveAll(`dist`)
		dir.Create(`dist`)
		for _, target := range p.Targets {
			for _, arch := range target.Arch {
				log.Printf("Building for %v/%v", target.OS, arch)
				name := fmt.Sprintf("%v_%v_%v", dir.Name(), target.OS, arch)
				os.Setenv(`GOOS`, target.OS)
				os.Setenv(`GOARCH`, arch)
				gobuild(`-o`, `dist/`+name)
			}
		}
		return nil
	},
}
