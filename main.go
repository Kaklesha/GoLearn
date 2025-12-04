package main

import "fmt"

func pointerViewer() {
	var pp *int
	fmt.Println(pp)
	i, j := 42, 2701
	p := &i //point to i
	fmt.Println(p)
	fmt.Println("----")
	fmt.Println(*p) //read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  //see the new value of i
	//dereferencing OR indirecting !!!!
	p = &j       //point to j
	*p = *p / 37 //devide j through the pointer
	fmt.Println(j)

}

type Vertex struct {
	X int
	Y int
}

func main() {
	//Struct Literals
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1} //Y:0 is implicit
		v3 = Vertex{}     // X:0 and Y:0
		//BELOW has type *Vertex // The special prefix "&"
		//  returns a pointer to the struct value
		pp = &Vertex{1, 2}
	)
	//Struct Literals demo
	fmt.Println(v1, pp, v2, v3)
	//Pointers to structs
	vv := Vertex{2, 3}
	p := &vv
	p.X = 1e9 // OR (*p).X = 1e9 However is cumbersome
	fmt.Println(vv)
	//Struct fields
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	//Structs
	fmt.Println(Vertex{1, 2})

	//Pointers AND 	//dereferencing OR indirecting !!!!
	pointerViewer()
	fmt.Println("====================================")
	// var ages [3] int = [3]int{20,25, 30}
	//	var ages = [3]int{20, 25, 30}

	//	names := [4]string{"fff", "gggg", "vvvvv", "bbbbb"}

	//fmt.Println(ages, len(ages))
	//fmt.Println(names, len(names))

	// //slice (use arrays under the hood)
	//var scores = []int{100, 50, 60}

	//	fmt.Println(scores, len(scores))

	//scores[2] = 25
	//scores = append(scores, 85)

	//fmt.Println(scores, len(scores))

	////slice ranges

	//rangeOne := names[1:3]

	//rangeTwo := names[2:]

	//rangeThree := names[:3]

	//fmt.Println(rangeOne, rangeTwo, rangeThree)

	//scoresh := append(rangeOne, "koopa")

	//fmt.Println(scoresh, len(scores))

}
