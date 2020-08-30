package myutils

func IsExistInStrings(strs []string, str string) bool {
    if v := GetIndexInStrings(strs, str); v == -1 {
        return false
    }
    return true
}

func GetIndexInStrings(strs []string, s string) int {
    for i, v := range strs {
        if s == v {
            return i
        }
    }
    return -1
}

func NoRepeatAppend(src []string, s string) ([]string, bool) {
    notExist := true
    for _, v := range src {
        if v == s {
            notExist = false
        }
    }
    if notExist {
        return append(src, s), true
    }
    return src, false
}

func NoRepeatAppends(src []string, s []string) ([]string, bool) {
    inserted := false
    for _, v1 := range s {
        notExist := true
        for _, v2 := range src {
            if v1 == v2 {
                notExist = false
                break
            }
        }
        if notExist {
            src = append(src, v1)
            inserted = true
        }
    }
    return src, inserted
}

