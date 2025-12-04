package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

// Exercise: Maps
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	sub := strings.Fields(s)
	//for _, i := range strings.Fields(s) {  // we could did it as here but
	// It is complicated solition
	for _, i := range sub {
		m[i] += 1
	}
	return m
}
func main() {

	names := []string{"yoshi", "mario", "peach", "luigi"}
	fmt.Printf("%T", names)
	// for i := 0; i < len(names); i++ {
	//   fmt.Println("valur of x is: ", names[i])
	// }

	// for index, value := range names {
	//   fmt.Printf("valur of x is index: %v  is %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("valur of x is index:  is %v \n", value)
	}

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68, -74.39,
	}
	fmt.Println(m["Bell Labs"])

	var mm = map[string]Vertex{
		//MAp literals
		"BEll": Vertex{40.324, 455.434},
		//MAp literals continued //если типо топ лвл:А
		// - это просто имя типа его можно опустить
		"Google": {344.55, 432.66},
	}
	fmt.Println(mm["BEll"])

	//Mutating maps
	mmm := make(map[string]int)

	mmm["answer"] = 42
	fmt.Println(mmm["answer"])
	mmm["answer"] = 48
	fmt.Println(mmm["answer"])
	delete(mmm, "answer")
	fmt.Println(mmm["answer"])

	v, ok := mmm["answer"]
	fmt.Println("value is ", v, "Present? - ", ok)
	//Exercise: Maps
	fmt.Println(WordCount("I am learning go ! I wanna go to home"))
}
