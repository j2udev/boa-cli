package cmd

import (
	"github.com/j2udev/boa"
	"github.com/j2udev/boa-cli/internal"
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	return boa.NewCobraCmd("initialize").
		WithAliases([]string{"init"}).
		WithRunFunc(internal.Initialize).
		Build()
}
