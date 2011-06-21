package main

import (
    "os"
    "strconv"
)

func extractArg (arg string) int {
    val, err := strconv.Atoi(arg)
    
    if err != nil {
        os.Stderr.WriteString(arg + " is not a valid number\n")
        os.Exit(1)
    }
    
    return val
}

func main() {
    Normal := 0
    Invalid := 2
    sep := "\n"
    out := os.Stdout
    err := os.Stderr
    
    help := "Usage: seq [options] START INC END\n" +
            "-s, --separator\t\tChanges the separator. Default is \\n\n" +
            "-a, --array\t\tSequences from 0 to END - 1. Cannot be used with START or INC\n" + 
            "-h, --help\t\tPrints this help\n"

    ln := len(os.Args)

    if ln == 1 {
        // out.WriteString("Usage:\nseq\t[n] | [x] [y]\tOutput either 1 to [n] or from [x] to [y].\n\t\t\tAdd --array to count from 0 to [n-1]\n")
        out.WriteString(help)
    } else {
        
        // Assume default starting
        start := 1
        end := 1
        inc := 1
        
        // Process arguments
        idx := 1
        for ; idx < ln; idx++ {
            if (os.Args[idx] == "--array" || os.Args[idx] == "-a") {
                // Array style printing
                start = 0
            } else if (os.Args[idx] == "--separator" || os.Args[idx] == "-s") {
                
                if (idx + 1 < len(os.Args)) {
                    sep = os.Args[idx + 1]
                    idx++
                } else {
                    err.WriteString("No separator given\n")
                    os.Exit(Invalid)
                }
            } else if (os.Args[idx] == "--help" || os.Args[idx] == "-h") {
                out.WriteString(help);
                os.Exit(Normal);
            } else { 
                // Processed all options, break
                break
            }
        }
        
        if idx >= ln {
            err.WriteString("No parameters given\n");
            os.Exit(Invalid)
        }
        
        // Figure out where we start and end
        
        if idx + 3 == ln {
            start = extractArg(os.Args[idx]); idx++
            inc = extractArg(os.Args[idx]); idx++
            end = extractArg(os.Args[idx]); idx++
        } else if idx + 2 == ln {
            start = extractArg(os.Args[idx]); idx++
            end = extractArg(os.Args[idx]); idx++
        } else if idx + 1 == ln {
            end = extractArg(os.Args[idx]); idx++
        }
        
        out.WriteString(strconv.Itoa(start))
        start += inc
        for ; start <= end; start += inc {
            out.WriteString(sep + strconv.Itoa(start))
        }
        out.WriteString("\n")
    }
    
    os.Exit(Normal)
}