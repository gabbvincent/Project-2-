package main

import "fmt"

func main() {

 var first int = 10
 var second int = 10
 var input int


  for x :=1; first >= 1 && second >= 1; x++ {

if x%2 == 0 { 
fmt.Println("Pile one =", first)
fmt.Println("Pile two =", second)
fmt.Println()
fmt.Println("Enter '1' or '2' to choose which pile.")
fmt.Scanln(&input)
if input == 1 {
first -= input 
} else if input == 2 {
second -= input

}
}
}
}