package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//var someNamw = "heel"

func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxINt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// NUmeric constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// Constants
	const World = "Piace"
	fmt.Println("hello", World)
	//Type inference
	var iii int = 5
	jj := iii //jj is an int
	jj = 55
	fmt.Printf("%T type %v\n", jj, jj)
	//----
	var (
		//type conversions
		x, y int     = 3, 4
		fff  float64 = math.Sqrt(float64(x*x + y*y))
		zz   uint    = uint(fff)
	)
	fmt.Println(x, y, zz)

	//zero values
	var (
		ii int
		ff float64
		bb bool
		ss string
	)
	//factored declaratiot variables
	fmt.Printf("%v %v %v %q\n", ii, ff, bb, ss)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxINt, MaxINt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//short declarate
	i, j := 1, 3
	//initialazed value fully without typed exclide
	var c, python, java = true, false, "no!"

	fmt.Println(i, j, c, python, java)
	//return naked
	fmt.Println(split(17))
	// fmt.Println(222)

	//strings
	var nameOne string = "ff"

	// var nameTwo = "luigu"

	//  var nameThree string

	// fmt.Println(nameOne,nameTwo,nameThree)

	nameOne = "fffas"

	nameFour := "peach"

	fmt.Println(nameOne, nameFour)

	//ints

	var ageOne int = 10
	var ageTwo int = 20
	var ageThree int = 40
	fmt.Println(ageOne, ageTwo, ageThree)

	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256

	fmt.Println(`num `, numOne, numTwo, numThree)
	fmt.Println(ageOne, ageTwo, ageThree)

	var scoreOne float32 = 25.98
	var scoreTwe float64 = 888999898989889899889.7

	scoreThree := 1.5
	fmt.Println(scoreOne, scoreTwe, scoreThree)

	a, b := swap("'hello'", "'world'")
	fmt.Println(a, b)

}
