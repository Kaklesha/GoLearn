package main

import (
  "fmt"

)

var score = 99.5

func main() {

 menu:= map[string] float64{
   "soup": 4.99,
   "pie": 7.99,
   "salad": 6.99,
   "coffee pudding": 3.55,
 }

 fmt.Println(menu) 
 fmt.Println("\n",menu["pie"])

  for k, v := range menu {
    fmt.Println(k," - ",v)
  }

  // ints as key type
  phonebook := map[int]string{
    23223323: "mario", 
    897887897: "luigig", 
    34342342: "peache",
  }

  fmt.Println(phonebook) 
  fmt.Println("\n",phonebook[23223323])
 
  phonebook[897887897]= "bpwser"

  fmt.Println(phonebook) 
}