package main

import (
    "log"
    "net/http"
    "os"
    "io"
)

func main() {
    pageHandler := func(w http.ResponseWriter, req *http.Request) {
        file, err := os.Open("index.html")
        if err != nil {
            log.Fatal(err)
        }
        data := make([]byte, 100)
        _, err = file.Read(data)
        if err != nil {
            log.Fatal(err)
        }
        io.WriteString(w, data)
    }

    http.HandleFunc("/index", pageHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
