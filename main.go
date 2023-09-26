package main

import "fmt"

var someNamw="heel"

func main() {

   age:= 35
   name:="shss"

   //Print
   fmt.Print("hello, ")
   fmt.Print("world!")
   fmt.Print("hello, \n ")

   //Println
   fmt.Println("me name",age, "ddddd", name)

   //Printf(formatted) %_ = format specifier
   fmt.Printf("me age is %v and my name is %v \n",age,name)

   fmt.Printf("me age is %q and my name is %q \n",age,name)
 
   fmt.Printf("me age is %T and my name is %T \n",age,name)

   fmt.Printf("you scored  %f points \n",255.33)

   fmt.Printf("you scored  %0.1f points \n",255.33)

   //Sprintf (save formatted string)

  var srt =  fmt.Sprintf(" my age is %v and my name is %v \n",age , name)

  fmt.Println("the saved stirfg is ",srt)

}