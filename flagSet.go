package main

import (
    "os"
    "fmt"

    flag "github.com/spf13/pflag"
)

func main() {
    create := flag.NewFlagSet("create", flag.ExitOnError)
    del := flag.NewFlagSet("delete", flag.ExitOnError)

    createFileFlag := create.String("filename", "nothing.yaml", "name of the file")
    createNsFlag := create.String("namespace", "default", "name of the namespace")

    delNsFlag := del.String("namespace", "default", "name of the namespace")

    if len(os.Args) < 2 {
        flag.PrintDefaults()
        os.Exit(1)
    }

    switch os.Args[1] {
        case "create":
            create.Parse(os.Args[2:])
            fmt.Println("Created resource using filename=" + *createFileFlag + ", namespace=" + *createNsFlag)
        case "delete":
            del.Parse(os.Args[2:])
            fmt.Println("Deleted resource in namespace=" + *delNsFlag)
        default:
            fmt.Println("Not a valid subcommand")
            os.Exit(1)
    }
}
