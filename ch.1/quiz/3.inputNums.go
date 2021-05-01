package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        text:= scanner.Text()
        if text == "STOP" {
            return
        }
        num, err := strconv.ParseInt(text, 0, 64)
        if err != nil {
            fmt.Println("Parsing error")
            continue
        }
        fmt.Println("value: ", num)
    }
}
