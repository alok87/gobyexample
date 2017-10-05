package main

import (
    "fmt"
//    flag "github.com/spf13/pflag"
    "github.com/spf13/cobra"
)


func NewCreateCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use: "create",
        Short: "create lucy",
        Long: "create lucy woman",
        Run: func (cmd *cobra.Command, args []string) {
            fmt.Println("creating")
        },
    }
    return cmd
}

func NewDeleteCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use: "delete",
        Short: "delete lucy",
        Long: "delete lucy man",
        Run: func (cmd *cobra.Command, args []string) {
            fmt.Println("deleting")
        },
    }
    return cmd
}

func NewCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use: "lucy",
        Short: "lucy teaches how to use flagsets",
        Long: "lucy teaches how to use flagsets",
        Run: func (cmd *cobra.Command, args []string) {
            cmd.Help()
        },
    }

    cmd.AddCommand(NewCreateCmd())
    cmd.AddCommand(NewDeleteCmd())

    return cmd
}

func main() {
    cmd := NewCmd()
    cmd.Execute()
}
