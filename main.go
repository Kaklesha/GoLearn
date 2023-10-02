package main

import (
  "fmt"

)

//var score = 99.5

func updateName (x *string){
  *x="wedge"
}

func main() {
name:="tifa"
 
m:=&name

 fmt.Println("memory address oa mane is: ",&name)
fmt.Println(name)

fmt.Println("memory address oa mane is: ",*m)
*m="dddd"
fmt.Println("memory address oa msse is: ",name)


updateName(m)
fmt.Println("memory address oa mane is: ",*m)
}