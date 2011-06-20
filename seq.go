package main

import (
    "os"
    "fmt"
    "strconv"
)

func extractArg (arg string) int {
    val, err := strconv.Atoi(arg)
    
    if err != nil {
        os.Stdout.WriteString(arg + " is not a valid number\n");
        os.Exit(1)
    }
    
    return val
}

func main() {
    Normal := 0
    Invalid := 2
    
    ln := len(os.Args)
    // out := os.Stdout.WriteString
    if ln == 1 {
        // out.call("World")
        os.Stdout.WriteString("Usage:\nseq\t[n] | [x] [y]\tOutput either 1 to [n] or from [x] to [y].\n\t\t\tAdd --array to count from 0 to [n-1]\n")
    } else {
        
        // First figure out where we are starting. 
        start := 1
        end := 1
        
        if ln == 2 {
            // Default printing
            end = extractArg(os.Args[1])
        } else if ln == 3 {
            // First, is the initial argument an array flag?
            if os.Args[1] == "--array" {
                start = 0
                end = extractArg(os.Args[2]) - 1
            } else {
                start = extractArg(os.Args[1])
                end = extractArg(os.Args[2])
            }
        } else {
            os.Stdout.WriteString("Invalid number of arguments passed\n");
            os.Exit(Invalid)
        }
        
        for ; start <= end; start++ {
            fmt.Printf("%d\n", start);
        }
    }
    
    os.Exit(Normal)
}