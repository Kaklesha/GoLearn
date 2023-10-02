package main

import ("fmt"
"bufio"
"strings"
"os")

func getInput( prompt string, r *bufio.Reader) (string,error) {
  fmt.Print(prompt)
  input, err := r.ReadString('\n')

  return strings.TrimSpace(input),err
}

func createBill() bill {
  reader :=bufio.NewReader(os.Stdin)


  name, _ :=getInput("Create a new bill name:", reader)

  b:= newBill(name)
  fmt.Println("Created the bill - ", b.name)
return b
}

func promptOptions(b bill) {
  reader:= bufio.NewReader(os.Stdin)
  opt, _ :=getInput("Choose option (a - add item , s - save bill , t - add tip):", reader)

  fmt.Println(opt)


}

func main() {

mybill:=createBill()
promptOptions(mybill)
fmt.Println("Created the  ",mybill)
//   mybill:= newBill("mario's bill")

// //mybill.format()

//   mybill.addItem("onion soup", 4.50)

//   mybill.addItem("veg pie", 8.95)
//   mybill.addItem("toffe pudding", 4.95)
//   mybill.addItem("coffee", 3.25)
//   mybill.updateTip(10)

//   fmt.Println(mybill.format())

}