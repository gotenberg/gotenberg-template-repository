package main

import (
	gotenbergcmd "github.com/gotenberg/gotenberg/v7/cmd"

	// Gotenberg modules. You may also cherry-pick the standard modules.
	_ "github.com/gotenberg/gotenberg/v7/pkg/standard"
	// Custom modules.
	// TODO: change namespace.
	_ "github.com/gotenberg/gotenberg-template-repository/pkg/modules/example"
)

func main() {
	gotenbergcmd.Run()
}
