package main

import "fmt"

func main() {
    str := "this is me"
    for i:=0; i<len(str); i++ {
        fmt.Println(str[i])
        str = "THAT IS NOT ME"
    }
}
