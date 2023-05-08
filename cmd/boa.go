package cmd

import (
	"github.com/j2udev/boa"
	"github.com/spf13/cobra"
)

func NewBoaCmd() *cobra.Command {
	return boa.NewCobraCmd("boa").
		WithLongDescription("The Boa CLI is used to generate a new go module that leverages the Boa library").
		WithShortDescription("generate a new Boa project").
		WithVersion("0.1.0").
		WithSubCommands(
			NewInitCmd(),
			NewGenerateCmd(),
		).
		Build()
}
