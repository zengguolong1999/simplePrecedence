package main

//import "os"
//import "log"
import "fmt"
import "strings"
//import "unicode"

func main() {
    //if err := os.Chmod("index.html", 0400); err != nil {
    //    log.Fatal(err)
    //}
    //datas := os.Environ()
    //for _, data := range datas {
    //    fmt.Println(data)
    //}
    var reader strings.Reader = strings.Reader{ "It is zgl", 0, -1 }
    a, err := reader.Seek(2, 3)
    fmt.Println(a, err)
}

func ReplaceAllWithOverlap(s, old, new string) string {
    str := []byte(s)
    oldstr := []byte(old)
    newstr := []byte(new)
    for i:=0; i<len(str)-len(oldstr)+1; i++{
        if str[i] == oldstr[0] {
            var j int
            for j=1; j<len(oldstr) && str[i+j]==oldstr[j]; j++ {
                ;
            }
            if j==len(oldstr) {
                temp := str[i+len(oldstr):]
                str = append(str[0:i:i], newstr...)
                str = append(str, temp...)
            }
        }
    }
    return string(str)
}

//func myReplaceAll(s, old, new string) string {
//
//}
