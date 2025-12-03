package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// exercise: Loops and funcs + by using Set
func Sqrt(x float64) float64 {
	z := float64(1)
	//init Set empty
	m := make(map[float64]bool)
	var count int
	for count = 1; count < 15; count++ {
		z -= (z*z - x) / (2 * z)
		//iterate Set for check
		_, found := m[z]
		if found {
			fmt.Printf("z is %v\n", z)
			return z
		}
		fmt.Println(z)
		//add value to Set
		m[z] = true
	}
	return z
}

var someNamw = "heel"

// if and else with a short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g>=%g\n", v, lim)
	}
	//can't use v here , though
	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func stackingDefersCallView() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("StackingDone")
}

func main() {
	//Same below but example with func
	// stackingDefersCallView()
	//Stacking defers calls
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("StackingDone")
	//defer
	defer fmt.Println("FILO Iam last but i was first ")

	//Switch with no condition instead of if-then-else chaining to lightly and simply
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good mooning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")

	}

	//Switch evaluation order
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In to days.")
	default:
		fmt.Println("Too far away.")

	}

	//Switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		//windows, freebsd ..etc
		fmt.Println("s%.\n", os)
	}

	fmt.Printf("Answer is %v\n", Sqrt(2))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(sqrt(2), sqrt(-4))

	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// loop forever

	// for {
	// 	}

	//while in Go
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// for sum:=1 ;sum < 1000; {
	// 	sum += sum
	// }
	// fmt.Println(sum)
	age := 35
	name := "shss"

	//Print
	fmt.Print("hello, ")
	fmt.Print("world!")
	fmt.Print("hello, \n ")

	//Println
	fmt.Println("me name", age, "ddddd", name)

	//Printf(formatted) %_ = format specifier
	fmt.Printf("me age is %v and my name is %v \n", age, name)

	fmt.Printf("me age is %q and my name is %q \n", age, name)

	fmt.Printf("me age is %T and my name is %T \n", age, name)

	fmt.Printf("you scored  %f points \n", 255.33)

	fmt.Printf("you scored  %0.1f points \n", 255.33)

	//Sprintf (save formatted string)

	var srt = fmt.Sprintf(" my age is %v and my name is %v \n", age, name)

	fmt.Println("the saved stirfg is ", srt)

}
