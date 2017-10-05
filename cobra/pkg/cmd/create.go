package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Short description for creating a cobra",
		Long:  "Long description for creating a cobra",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("created cobra")
		},
	}
}
