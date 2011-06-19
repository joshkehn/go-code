package main
import "fmt"

func main() {
    if (true) {
        fmt.Printf("true")
    } else if (!false) {
        fmt.Printf("false")
    } else {
        fmt.Printf("other")
    }
}