package myutils

import (
    "os"
    "log"
    "errors"
)

//This struct is used to parse the data read from file into lines.
type lineReader struct {
    data []byte
    pos int
}

func newLineReader(data []byte) *lineReader {
    return &lineReader{ data, 0 }
}

func (l *lineReader) readLine() ([]byte, error) {
    i := l.pos
    for ; i < len(l.data) && l.data[i] != '\n'; i++ {
    }
    if i == len(l.data) {
        oldPos := l.pos
        l.pos = i
        err := errors.New("Reach end of file.")
        return l.data[oldPos:i:i], err
    }
    oldPos := l.pos
    l.pos = i+1
    return l.data[oldPos:i+1:i+1], nil
}

func ReadFileIntoLine(filename string) [][]byte {
    file, err := os.Open(filename)
    defer file.Close()
    if err != nil {
        log.Fatal(err)
    }
    data := make([]byte, 10000)
    count, err := file.Read(data)
    if err != nil {
        log.Fatal(err)
    }
    data = data[:count]
    l := newLineReader(data)
    lines := make([][]byte, 0)
    count = 0
    for {
        line, err := l.readLine()
        if len(line) != 0 {
            lines = append(lines, line)
            count++
        }
        if err != nil {
            break
        }
    }
    return lines
}

