package main

import (
    "os"
    "strings"
    "log"
    "fmt"
)

const (
    MAX_LEN int = 2000
)

func main() {
    file, err := os.Open("data")
    if err != nil {
        log.Fatal(err)
    }
    data := make([]byte, MAX_LEN)
    count, err := file.Read(data)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("read %d bytes", count)
    words := strings.ToUpper(string(data))
    fmt.Println(words)
}
