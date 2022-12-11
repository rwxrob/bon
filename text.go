package bon

import _ "embed"

//go:embed text/en/bon.md
var _bon string

//go:embed text/en/command.md
var _command string

// ----------------------------- templates ----------------------------

//go:embed tmpl/command.tmpl
var commandtmpl string
