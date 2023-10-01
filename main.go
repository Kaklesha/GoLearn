package main

import ("fmt")

func main() {

  
names := []string{"yoshi","mario","peach","luigi"}


  
  // for i := 0; i < len(names); i++ {
  //   fmt.Println("valur of x is: ", names[i])
  // }

// for index, value := range names {
//   fmt.Printf("valur of x is index: %v  is %v \n", index, value)
// }

for _, value := range names {
  fmt.Printf("valur of x is index:  is %v \n", value)
}



}