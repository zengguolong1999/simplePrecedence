package myutils

func ExistInStrings(strs []string, str string) bool {
    for _, s := range strs {
        if s == str {
            return true
        }
    }
    return false
}

func TransStringByte(s string) byte {
    if len(s) != 0 {
        return s[0]
    }else {
        var b byte
        return b
    }
}

