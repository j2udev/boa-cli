package cmd

import (
	"github.com/j2udev/boa"
	"github.com/j2udev/boa-cli/internal"
	"github.com/spf13/cobra"
)

func NewGenerateCmd() *cobra.Command {
	return boa.NewCobraCmd("generate").
		WithAliases([]string{"gen"}).
		WithLongDescription("Generate a new Boa CLI").
		WithShortDescription("generate a new Boa CLI").
		WithRunFunc(internal.Generate).
		WithBoolPFlag("init", "i", false, "run a go mod init/tidy upon generation").
		WithBoolPFlag("overwrite", "o", false, "overwrite any file collisions upon generation").
		WithBoolPFlag("dry-run", "d", false, "generate project to stdOut instead of writing to the filesystem").
		Build()
}
