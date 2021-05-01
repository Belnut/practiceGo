package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    params := os.Args
    var total int64
    total = 0

    for i := 1; i < len(params); i++ {
        num, err := strconv.ParseInt(params[i], 0, 64)
        if err != nil {
            fmt.Println("parsing error!")
            os.Exit(1)
        }
        total += num
    }
    fmt.Printf("total : %d\n", total)
}
