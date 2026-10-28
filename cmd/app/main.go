package main

import (
	gotenbergcmd "github.com/gotenberg/gotenberg/v8/cmd"
	_ "github.com/gotenberg/gotenberg/v8/pkg/standard"

	// TODO: change the module path to match the module name in go.mod.
	_ "github.com/gotenberg/gotenberg-template-repository/pkg/modules/example"
)

func main() {
	gotenbergcmd.Run()
}
