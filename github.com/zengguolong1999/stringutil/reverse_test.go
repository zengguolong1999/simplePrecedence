package stringutil

import "testing"

func TestReverse(t *testing.T) {
    cases := []struct{
        in, wants string
    }{
        { "Hello, world!", "!dlrow ,olleH" },
        { "hello", "olleh" },
        { "hola", "aloh" },
    }
    for _, v := range cases {
        got := Reverse(v.in)
        if got != v.wants {
            t.Errorf("some")
        }
    }
}
