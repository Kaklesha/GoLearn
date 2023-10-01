package main

import ("fmt"
"math"
)

func sayGreeting( n string )  {
  fmt.Printf("Good morning %v  \n"  ,n)
}

func sayBye( n string )  {
  fmt.Printf("Good bye %v  \n"  ,n)
}

func cycleNames( n [] string , f func(string))  {
  for _,value := range n {
    f(value)
  }
}

func circleArea(r float64 ) float64 {
  return math.Pi * r * r
}

func main() {

 

 cycleNames([]string{"mario","luigi", "yoshi", "peache"}, sayGreeting)

 cycleNames([]string{"mario","luigi", "yoshi", "peache"}, sayBye)
//names := [] string {"mario","luigi", "yoshi", "peache"}

a1:=circleArea(4.2)
a2:=circleArea(4)

fmt.Println( a1,a2 )
fmt.Printf( "circle 1 is %0.3f and circel 2 is %0.3f",a1,a2 )

}