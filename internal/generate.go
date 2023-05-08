package internal

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/j2udev/boa-cli/templates"
	"github.com/spf13/cobra"
)

type templateData struct {
	Module string
	Command
}

var (
	dryRun     bool
	initialize bool
	overwrite  bool
)

func Generate(cmd *cobra.Command, args []string) {
	// read in boa config file
	// this set the Cfg variable defined in config.go
	readConfig()
	// set global vars based on cobra cmd flagset
	parseFlags(cmd)
	// generate main.go
	generateMain()
	// generate commands (and respective packages) recursively
	generateCommand(Cfg.Command)
	// go mod commands only called if the init flag is set
	if initialize {
		goModInit()
		goModTidy()
	}
}

func generateMain() {
	handleErr(templates.NewMainTemplate().Execute(getWriter("main.go"), Cfg))
}

func generateCommand(cmd Command) {
	if !dryRun {
		handleErrDebug(os.Mkdir("cmd", 0755))
	}
	filePath := fmt.Sprintf("cmd/%s.go", cmd.Name)
	handleErr(templates.NewCobraCmdTemplate().Execute(getWriter(filePath), tmplData(cmd)))
	generatePkg(cmd)
	if !cmd.HasSubCommands() {
		return
	}
	for _, c := range cmd.SubCommands {
		generateCommand(c)
	}
}

func generatePkg(cmd Command) {
	pkgName := "internal"
	if !dryRun && !cmd.HasSubCommands() {
		if cmd.PkgName == "" {
			handleErrDebug(os.Mkdir(pkgName, 0755))
		} else {
			pkgName = fmt.Sprintf("internal/%s", cmd.PkgName)
			handleErrDebug(os.MkdirAll(pkgName, 0755))
		}
	}
	if !cmd.HasSubCommands() {
		filePath := fmt.Sprintf("%s/%s.go", pkgName, cmd.Name)
		handleErr(templates.NewPkgTemplate().Execute(getWriter(filePath), tmplData(cmd)))
	}
}

func goModInit() {
	var stdErr bytes.Buffer
	cmd := exec.Command("go", "mod", "init", Cfg.Module)
	cmd.Stderr = &stdErr
	err := cmd.Run()
	handleErr(err)
}

func goModTidy() {
	var stdErr bytes.Buffer
	cmd := exec.Command("go", "get", "github.com/j2udev/boa@main")
	cmd.Stderr = &stdErr
	err := cmd.Run()
	handleErr(err)

	stdErr.Reset()
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Stderr = &stdErr
	err = cmd.Run()
	handleErr(err)
}

func getWriter(filePath string) io.Writer {
	if dryRun {
		return os.Stdout
	}
	if overwrite {
		return createFile(filePath)
	}
	if !exists(filePath) {
		return createFile(filePath)
	}
	log.Fatalf("%s already exists; if you wish to overwrite it use the --overwrite flag", filePath)
	return nil
}

func parseFlags(cmd *cobra.Command) {
	var err error
	flags := cmd.Flags()
	dryRun, err = flags.GetBool("dry-run")
	handleErr(err)
	initialize, err = flags.GetBool("init")
	handleErr(err)
	overwrite, err = flags.GetBool("overwrite")
	handleErr(err)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func createFile(filePath string) *os.File {
	file, err := os.Create(filePath)
	handleErr(err)
	return file
}

func tmplData(cmd Command) templateData {
	return templateData{
		Module:  Cfg.Module,
		Command: cmd,
	}
}
