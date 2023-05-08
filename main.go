package main

import (
	"github.com/j2udev/boa-cli/cmd"
)

func main() {
	err := cmd.NewBoaCmd().Execute()
	if err != nil {
		panic(err)
	}
}
