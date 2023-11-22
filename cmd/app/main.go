package main

import (
	gotenbergcmd "github.com/gotenberg/gotenberg/v7/cmd"
	_ "github.com/gotenberg/gotenberg/v7/pkg/standard"

	// TODO: change namespace.
	_ "github.com/gotenberg/gotenberg-template-repository/pkg/modules/example"
)

func main() {
	gotenbergcmd.Run()
}
