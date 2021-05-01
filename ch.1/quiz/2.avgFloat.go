package main

import (
    "os"
    "strconv"
    "fmt"
)

func main() {
    params := os.Args
    var total float64
    total = 0.0

    for i := 1; i < len(params); i++ {
        num, err := strconv.ParseFloat(params[i], 64)
        if err != nil {
            fmt.Println("parsing error!")
            os.Exit(1)
        }
        total += num;
    }
    total /= (float64)(len(params) - 1)
    fmt.Printf("avg : %f\n", total)
}
