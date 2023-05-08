package internal

import (
	"github.com/j2udev/boa-cli/templates"
	"github.com/spf13/cobra"
)

func Initialize(cmd *cobra.Command, args []string) {
	generateBoaConfig()
}

func generateBoaConfig() {
	handleErr(templates.NewBoaConfigTemplate().Execute(getWriter("boa.yml"), true))
}
