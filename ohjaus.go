package main

import (
    "os"
    "flag"
)

var help = flag.Bool("help", false, "Don't print final newline")
var version = flag.Bool("version", false, "Print version string")

const (
    Space = " "
    Newline = "\n"
)

func main() {
    flag.Parse();
    
    if *help {
        os.Stdout.WriteString("HELP\n")
    } else if *version {
        os.Stdout.WriteString("VERSION\n")
    } else {
        flag.PrintDefaults();
    }
}