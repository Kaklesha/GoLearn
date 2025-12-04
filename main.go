package main

import (
	"fmt"
	"math"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

// MUTATE ABOUT: add the pointer and value receivers methons
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}
func (v *Vertex) Scale(f float64) {
	v.Long = v.Long * f
	v.Lat = v.Lat * f
}

var m map[string]Vertex

// Methods for example need it below
type Vertexe struct {
	X, Y float64
}

// //Methods
// rem a method is just a func with receiver argument and unstandart syntax
func (v Vertexe) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods continued for non-struct types too
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

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

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// create func closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println("//=======pointer and value receivers methons==========")
	n := Vertex{3, 4}
	fmt.Println(n)
	fmt.Println(n.Abs())
	fmt.Println(n)
	n.Scale(10)
	fmt.Println(n)

	fmt.Println(n.Abs())

	fmt.Println("//=======method continued for non-struct==========")
	g := MyFloat(-math.Sqrt2)
	fmt.Println(g.Abs())
	fmt.Println("//=======Methods using=============")
	//Methods using
	f := Vertexe{3, 4}
	fmt.Println(f.Abs())
	fmt.Println("//====================")
	//Release func closure
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	//Function values aka as signatire into another func
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))

	fmt.Println(compute(math.Pow))
	//====================
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
