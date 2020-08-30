package main

import (
    "fmt"
//    "strings"
//    "os"
)

type ints []int
func (data ints)Printsomething() {
    fmt.Println(data)
}

type table struct {
    th []string
    content [][]int
}

func main() {
    /*
    a := make([][]int, 3)
    d1 := []int{ 1, 2, 3 }
    d2 := []int{ 4, 5, 6 }
    d3 := []int{ 7, 8, 9 }
    a = append(a, d1)
    a = append(a, d2)
    a = append(a, d3)
    */
    d1 := []int{ 1, 2, 3 }
    for i, v := range d1 {
        fmt.Println("time before:", d1)
        if v <= 99 {
            temp := d1[i+1:]
            d1 = d1[:i:i]
            d1 = append(d1, v+1)
            d1 = append(d1, temp...)
        }
        fmt.Println("time after:", d1)
    }
}
