package internal

import (
	"os"

	"github.com/j2udev/boa"
)

type (
	Config struct {
		Command Command `mapstructure:"command"`
		Module  string  `mapstructure:"module"`
	}

	Command struct {
		PkgName     string    `mapstructure:"pkg"`
		SubCommands []Command `mapstructure:"subCommands"`
		Flags       []Flag    `mapstructure:"flags"`
		Name        string    `mapstructure:"name"`
		Aliases     []string  `mapstructure:"aliases"`
		Long        string    `mapstructure:"long"`
		Short       string    `mapstructure:"short"`
		Example     string    `mapstructure:"example"`
		Version     string    `mapstructure:"version"`
	}

	Flag struct {
		Type         string `mapstructure:"type"`
		Persistent   bool   `mapstructure:"persistent"`
		Name         string `mapstructure:"name"`
		Shorthand    string `mapstructure:"shorthand"`
		Usage        string `mapstructure:"usage"`
		DefaultValue any    `mapstructure:"default"`
	}
)

var Cfg Config

func readConfig() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = boa.NewViperCfg().
		WithConfigName("boa").
		WithConfigPaths(cwd).
		ReadInConfig().
		Build().
		UnmarshalExact(&Cfg)
	// err := boa.NewDefaultViperCfg("boa").Build().UnmarshalExact(&Cfg)
	if err != nil {
		log.Fatal(err)
	}
	if Cfg.Module == "" {
		log.Fatal("module must be specified")
	}
	setDefaults(&Cfg.Command)
}

func validateCmd(cmd Command) {
	if cmd.PkgName == "main" || cmd.PkgName == "init" || cmd.PkgName == "unsafe" {
		log.Fatal("invalid package name for cmd: %s", cmd.Name)
	}
}

func validateFlag(flag Flag) {
	switch flag.Type {
	case "string":
	case "bool":
	default:
		log.Fatalf("invalid flag type %s for flag %s", flag.Type, flag.Name)
	}
}

func setDefaults(cmd *Command) {
	validateCmd(*cmd)
	for i, f := range cmd.Flags {
		validateFlag(f)
		if f.Type == "bool" && f.DefaultValue == nil {
			cmd.Flags[i].DefaultValue = false
		}
	}
	subCommands := cmd.SubCommands
	for i := range subCommands {
		setDefaults(&subCommands[i])
	}
}

func (c Command) HasSubCommands() bool {
	if len(c.SubCommands) > 0 {
		return true
	}
	return false
}

func (c Command) HasFlags() bool {
	if len(c.Flags) > 0 {
		return true
	}
	return false
}
