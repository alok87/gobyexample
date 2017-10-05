package main

import (
	"github.com/alok87/gobyexample/cobra/pkg/cmd"
)

func main() {

	cmd := cmd.NewCobraCommand()
	cmd.Execute()

}
