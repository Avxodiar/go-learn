package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Println("Welcome to the real world!")
    fmt.Println("Are you ready? (yes/no)")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    ready := scanner.Text()

    fmt.Println()
    if ready == "yes" {
        fmt.Println("Let`s GO!")
    } else {
        fmt.Println("You are not Prepared!")
    }
}
