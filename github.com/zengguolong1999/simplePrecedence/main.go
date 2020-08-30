package main

import (
    "fmt"
    "strings"
    "log"
    "errors"
    "github.com/zengguolong1999/simplePrecedence/myutils"
)

const (
    NONTERMINATE int = 1
    TERMINATE int = 2
    MAXLEN int = 10000
    UNDEFINED int = 0
    LT int = 1
    EQ int = 2
    GT int = 3
)

var (
    ErrTableNotFoundName = errors.New("Name not found in table")
    ErrStackOverflow = errors.New("stack overflow")
    ErrStackEmpty = errors.New("stack is empty")
    ErrNotFoundRightpartOfProd = errors.New("rightpart not found in prods")
)

type elem struct {
    elemType int
    name string
}

type prod struct {
    leftPart elem
    rightParts [][]elem
}

func getIndexOfProd(prods []prod, name string) int {
    for i, v := range prods {
        if v.leftPart.name == name {
            return i
        }
    }
    return -1
}

//prodPack is used just to contain all information of productions.
type prodPack struct {
    startSym string
    nonTerminals []string
    terminals []string
    prods []prod
}

//This struct is used to represent the table used in simple precedence grammar.
type table struct {
    th []string
    content [][]int
}

//initTable is used to allocate memory to table for the later insert.
func initTable(th []string) table {
    var t table
    l := len(th)
    t.th = append(t.th, th...)
    t.content = make([][]int, l)
    for i:=0; i<l; i++ {
        t.content[i] = make([]int, l)
    }
    return t
}

func (t *table)insert(x, y string, relationship int) error {
    xi := myutils.GetIndexInStrings(t.th, x)
    yi := myutils.GetIndexInStrings(t.th, y)
    if xi == -1 || yi == -1 {
        return ErrTableNotFoundName
    }
    t.content[xi][yi] = relationship
    return nil
}

func (t table)getValueByName(x, y string) (int, error) {
    var xi, yi int = -1, -1
    for i, v := range t.th {
        if v == x {
            xi = i
        }
        if v == y {
            yi = i
        }
    }
    if xi == -1 || yi == -1 {
        return 0, ErrTableNotFoundName
    }
    return t.content[xi][yi], nil
}

func (t table)printTable() {
    fmt.Print("  ")
    for _, v := range t.th {
        fmt.Print(v, " ")
    }
    fmt.Println()
    for i, v := range t.th {
        fmt.Print(v, " ")
        for _, v2 := range t.content[i] {
            if v2 == LT {
                fmt.Print("<", " ")
            }else if v2 == EQ {
                fmt.Print("=", " ")
            }else if v2 == GT {
                fmt.Print(">", " ")
            }else {
                fmt.Print(0, " ")
            }
        }
        fmt.Println()
    }
}

func buildTable(pack prodPack) table {
    var th []string
    th = append(th, pack.nonTerminals...)
    th = append(th, pack.terminals...)
    th = append(th, "#")
    ps := getPrefixSuffixOfNon(pack.prods)
    t := initTable(th)
    for _, v1 := range pack.prods {
        for _, rightPart := range v1.rightParts {
            if len(rightPart) >= 2 {
                for i:=0; i<len(rightPart)-1; i++{
                    a, b := rightPart[i], rightPart[i+1]
                    //equal
                    t.insert(a.name, b.name, EQ)
                    //less than
                    if b.elemType == NONTERMINATE {
                        index := getIndexOfPresuff(ps, b.name)
                        for _, v := range ps[index].prefix {
                            t.insert(a.name, v, LT)
                        }
                    }
                    //greater than
                    if a.elemType == NONTERMINATE {
                        index := getIndexOfPresuff(ps, a.name)
                        if b.elemType == NONTERMINATE {
                            indexB := getIndexOfPresuff(ps, b.name)
                            tempPrefix := make([]string, 0)
                            tempPrefix = append(tempPrefix, ps[indexB].prefix...)
                            tempPrefix = append(tempPrefix, b.name)
                            for _, v := range ps[index].suffix {
                                for _, v2 := range tempPrefix {
                                    t.insert(v, v2, GT)
                                }
                            }
                        }else {
                            for _, v := range ps[index].suffix {
                                t.insert(v, b.name, GT)
                            }
                        }
                    }
                    //process '#'
                    index := getIndexOfPresuff(ps, pack.startSym)
                    for _, v := range ps[index].prefix {
                        t.insert("#", v, LT)
                    }
                    for _, v := range ps[index].suffix {
                        t.insert(v, "#", GT)
                    }
                }
            }
        }
    }
    return t
}

//nonPreSuffix is used to build the relationship table.
type nonPreSuffix struct {
    nonT string
    prefix []string
    suffix []string
}

func getIndexOfPresuff(ps []nonPreSuffix, non string) int {
    for i, v := range ps {
        if v.nonT == non {
            return i
        }
    }
    return -1
}

//getPrefixSuffixOfNon assume that the prods are valid.
func getPrefixSuffixOfNon(prods []prod) []nonPreSuffix {
    res := make([]nonPreSuffix, len(prods))
    for i, v := range prods {
        res[i].nonT = v.leftPart.name
    }
    //get prefix
    for isChanged, times := true, 1; isChanged; times++ {
       isChanged = false
       for i, v := range prods {
            for _, v2 := range v.rightParts {
                var isInserted bool
                if times == 1 {
                    if temp := v2[0]; temp.elemType == NONTERMINATE {
                        res[i].prefix, _ = myutils.NoRepeatAppend(res[i].prefix, temp.name)
                        index := getIndexOfProd(prods, temp.name)
                        res[i].prefix, isInserted = myutils.NoRepeatAppends(res[i].prefix, res[index].prefix)
                    }else {
                        res[i].prefix, isInserted = myutils.NoRepeatAppend(res[i].prefix, temp.name)
                    }
                }else {
                    if temp := v2[0]; temp.elemType == NONTERMINATE {
                        index := getIndexOfProd(prods, temp.name)
                        res[i].prefix, isInserted = myutils.NoRepeatAppends(res[i].prefix, res[index].prefix)
                    }
                }
                if isInserted {
                    isChanged = true
                }
            }
       }
    }
    //get suffix
    for isChanged, times := true, 1; isChanged; times++ {
       isChanged = false
       for i, v := range prods {
            for _, v2 := range v.rightParts {
                var isInserted bool
                if times == 1 {
                    if temp := v2[len(v2)-1]; temp.elemType == NONTERMINATE {
                        res[i].suffix, _ = myutils.NoRepeatAppend(res[i].suffix, temp.name)
                        index := getIndexOfProd(prods, temp.name)
                        res[i].suffix, isInserted = myutils.NoRepeatAppends(res[i].suffix, res[index].suffix)
                    }else {
                        res[i].suffix, isInserted = myutils.NoRepeatAppend(res[i].suffix, temp.name)
                    }
                }else {
                    if temp := v2[len(v2)-1]; temp.elemType == NONTERMINATE {
                        index := getIndexOfProd(prods, temp.name)
                        res[i].suffix, isInserted = myutils.NoRepeatAppends(res[i].suffix, res[index].suffix)
                    }
                }
                if isInserted {
                    isChanged = true
                }
            }
       }
    }
    return res
}

func getlist(lines [][]byte, index int) []string {
    s := strings.Fields(string(lines[index]))
    return s
}

func getTlist(lines [][]byte) []string {
    return getlist(lines, 1)
}

func getNontlist(lines [][]byte) []string {
    return getlist(lines, 0)
}

func getProduction(lines [][]byte) ([]prod, error) {
    tlist := getTlist(lines)
    nonTlist := getNontlist(lines)
    prods := make([]prod, 0)
    for i := 2; i<len(lines); i++ {
        parts := strings.Fields(string(lines[i]))
        if len(parts) != 2 {
            return nil, errors.New("not two parts")
        }
        left := parts[0]
        if len(left) != 1 {
            return nil, errors.New("left not one character")
        }
        if !myutils.IsExistInStrings(nonTlist, left) {
            return nil, errors.New("left is not a non")
        }
        leftElem := elem{ NONTERMINATE, left }
        rights := strings.Split(parts[1], "|")
        tempA := make([][]elem, 0)
        for _, s := range rights {
            tempB := make([]elem, 0)
            for _, e := range strings.Split(s, "") {
                if myutils.IsExistInStrings(nonTlist, e) {
                    tempB = append(tempB, elem{ NONTERMINATE, e })
                }else if myutils.IsExistInStrings(tlist, e) {
                    tempB = append(tempB, elem{ TERMINATE, e })
                }else {
                    return nil, errors.New("unknown character in right part")
                }
            }
            tempA = append(tempA, tempB)
        }
        prods = append(prods, prod{ leftElem, tempA })
    }
    return prods, nil
}

func printProd(prods []prod) {
    seperator := " -> "
    for _, p := range prods {
        indent := len(p.leftPart.name)
        fmt.Print(p.leftPart.name)
        fmt.Print(seperator)
        for _, k := range p.rightParts[0] {
            fmt.Print(k.name)
        }
        fmt.Println()
        for i:=1; i<len(p.rightParts); i++ {
            for j:=0; j<indent; j++ {
                fmt.Print(" ")
            }
            fmt.Print(seperator)
            for _, k := range p.rightParts[i] {
                fmt.Print(k.name)
            }
            fmt.Println()
        }
    }
}

func parseS(s string) []string {
    res := strings.Split(s, "")
    return res
}

type intStack struct {
    stack [1000]int
    sp int
}

func (st *intStack)push(v int) error {
    if st.sp >= len(st.stack) {
        return ErrStackOverflow
    }
    st.stack[st.sp] = v
    st.sp++
    return nil
}

func (st *intStack)pop() (int, error) {
    if st.sp <= 0 {
        return 0, ErrStackEmpty
    }
    st.sp--
    return st.stack[st.sp], nil
}

//checkRightPartOfProd is used to check the right part of productions to get the left part.
func checkRightPartOfProd(prods []prod, s []string) (string, error) {
    l := len(s)
    for indexOfLeftpart, p := range prods {
        for _, e := range p.rightParts {
            if len(e) != l {
                continue
            }
            var i int
            for i=0; i<l; i++ {
                if s[i] != e[i].name {
                    break
                }
            }
            if i == l {
                return prods[indexOfLeftpart].leftPart.name, nil
            }
        }
    }
    return "", ErrNotFoundRightpartOfProd
}

func checkStrInGrammar(s string, pack prodPack) bool {
    prods := pack.prods
    startSym := pack.startSym
    stack := new(intStack)
    checkS := make([]string, 0)
    checkS = append(checkS, "#")
    checkS = append(checkS, parseS(s)...)
    checkS = append(checkS, "#")
    t := buildTable(pack)
    for i:=0; i<len(checkS); {
        rel, err := t.getValueByName(checkS[i], checkS[i+1])
        if err != nil {
            log.Fatal(err)
        }
        if rel == UNDEFINED {
            log.Fatal("the two combination of ", checkS[i], " ", checkS[i+1], " is invalid")
        }else if rel == LT {
            stack.push(i)//checkS[i] is at the left of LT
            i++
        }else if rel == GT {
            lBound, err := stack.pop()
            if err != nil {
                log.Fatal(err)
            }
            sub := checkS[lBound+1:i+1]
            left, err := checkRightPartOfProd(prods, sub)
            if err != nil {
                log.Fatal(err)
            }
            //change 'checkS'
            temp := checkS[i+1:]
            checkS = checkS[:lBound+1:lBound+1]
            checkS = append(checkS, left)
            checkS = append(checkS, temp...)
            if left == startSym && len(checkS) == 3 {
                return true
            }
            i = lBound
        }else {//situation EQ
            i++
        }
    }
    return false
}

func main() {
    lines := myutils.ReadFileIntoLine("data")
    startSym := "S"
    prods, err := getProduction(lines)
    if err != nil {
        fmt.Println("Good luck!")
    }
    nonTerminals := getNontlist(lines)
    terminals := getTlist(lines)
    p := prodPack{
        startSym,
        nonTerminals,
        terminals,
        prods,
    }
    res := checkStrInGrammar("b((aa)a)b", p)
    fmt.Println(res)
}
