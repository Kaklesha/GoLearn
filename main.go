package main

import ("fmt"
"strings"
"sort"
)

func main() {

  // var ages [3] int = [3]int{20,25, 30} 
var greeting  = " hello there fret "
//substring including
// fmt.Println(strings.Contains(greeting,"hello") )

// fmt.Println(strings.ReplaceAll(greeting,"hello","hi") )

//fmt.Println(strings.ToUpper(greeting) )

//fmt.Println(strings.Index(greeting, "hello") )

fmt.Println(strings.Split(greeting, " ") )


fmt.Println("Original str:", greeting )

ages := []int {45,20,35,30,75,60,50,25}
sort.Ints(ages)

fmt.Println(ages)

index:= sort.SearchInts(ages, 30)
fmt.Println(index)

names := []string{"yoshi","mario","peach","bowser","luigi"}
fmt.Println(names)

sort.Strings(names)
fmt.Println(names)
fmt.Println(sort.SearchStrings(names,"bowser"))


}