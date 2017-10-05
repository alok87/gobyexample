package main

import (
    "fmt"
    //flag "github.com/spf13/pflag"
    "flag"
)

/*
Result
$ go run flags.go -i=542 -b=true get pods
def-value 61 false
def-value 542 true
Non flags command line args [get pods]
*/

func main() {
    strFlag := flag.String("sFlag", "def-value", "describing string flag")
    intFlag := flag.Int("iFlag", 61, "describing int flag")
    boolFlag := flag.Bool("bFlag", false, "describing bool flag")

    var s string
    var i int
    var b bool

    flag.StringVar(&s, "s", "def-value", "usage-s")
    flag.IntVar(&i, "i", 62, "usage-i")
    flag.BoolVar(&b, "b", true, "usage-b")

    flag.Parse()
    fmt.Println(*strFlag, *intFlag, *boolFlag)
    fmt.Println(s, i, b)
    fmt.Println("Non flags command line args", flag.Args())
}
