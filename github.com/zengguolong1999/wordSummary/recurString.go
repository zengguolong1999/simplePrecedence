package main

import "fmt"

type mytype string

func (v mytype)String() string {
    return fmt.Sprintf("%s", v)
}
func main() {
    var sa mytype = mytype("this")
    fmt.Printf("%v", sa)
}
