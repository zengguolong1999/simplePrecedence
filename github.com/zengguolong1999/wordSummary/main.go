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
    data = data[:count]
    words := strings.Fields(string(data))
    mapwords := make(map[string]int)
    for _, word := range words {
        word = strings.Trim(word, ",.:\"“”!?")
        mapwords[word] += 1
    }
    for word, count := range mapwords {
        fmt.Println(word, ": ", count)
    }
}
