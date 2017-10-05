package cmd

import (
	"github.com/spf13/cobra"
)

//NewCobraCommand is a Root command in the command tree
func NewCobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cobra",
		Short: "cobra root command short information",
		Long:  "cobra root command long information",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	//Adds a new Command to the root
	cmd.AddCommand(NewCreateCommand())

	return cmd
}
