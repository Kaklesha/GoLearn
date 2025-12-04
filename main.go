package main

import "fmt"

// func convert(foos []Foo) []Bar{
//   n:= len(foos)
//   bars:= make([]Bar,0,n)//создали срез с заданой емкостью
//   for _, foo:= range foos {
//     bars = append(bars, fooToBar(foo)) //// foToBar - lorem for example
//   }
//   return bars
// }

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
	//Array
	var a [2]string
	a[0] = "hello"
	a[1] = "World!"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

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
	var ages = [3]int{20, 25, 30}

	names := [4]string{"fff", "gggg", "vvvvv", "bbbbb"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// //slice (use arrays under the hood)
	var scores = []int{100, 50, 60}

	fmt.Println(scores, len(scores))

	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	////slice ranges
	fmt.Println("//======slice=======")
	fmt.Println(names, len(names))

	rangeOne := names[1:3]
	//slice default
	rangeTwo := names[2:]
	//slice default
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	scoresh := append(rangeOne, "koopa")

	fmt.Println(scoresh, len(scores))

	primes = [6]int{2, 3, 5, 7, 11, 13}

	var sss []int = primes[1:4]
	vvv := &primes
	var ttt []int = (*vvv)[0:4]
	fmt.Println(sss)
	fmt.Println(ttt)

	namees := [4]string{
		"john",
		"Paul",
		"george",
		"rino",
	}

	fmt.Println(namees)

	aa := namees[0:2]
	bb := namees[1:3]
	fmt.Println(aa, bb)

	bb[0] = "XXXX"
	fmt.Println(aa, bb)
	fmt.Println(names)
	//use make for create map\slice\channel, for slice init len and capacity
	s1 := make([]int, 3, 6) // [0 0 0]  // 3 - len, 6 - power
	s2 := s1[1:3]
	s1[1] = 1
	fmt.Println(s1, s2)
	//Slice literals
	qq := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(qq)

	rrr := []bool{true, false, true}
	fmt.Println(rrr)

	ssss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(ssss)
	//NIl slices
	var sssss []int
	fmt.Println(sssss, len(sssss), cap(sssss))
	if sssss == nil {
		fmt.Println("nil!")
	}
	//appending to a slice
	sssss = append(sssss, 3, 5, 6, 7)
	fmt.Println(sssss, len(sssss), cap(sssss))
}
