package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

 var first int = 10
 var second int = 10
 var inputA int
 var inputB int

fmt.Println()
fmt.Println("----------------------------------------------------")

 //rules
 fmt.Println("1.) In this game there are two piles of ten pebbles.")
 fmt.Println()
 fmt.Println("2.) You can take up to 3 pebbles but only from one pile each turn.")
 fmt.Println()
 fmt.Println("3.) To win, be the person to take the last pebble.")
 
fmt.Println("----------------------------------------------------")
 
 //game begins
  for x :=0; first >=0 && second >= 0; x++ {

if x%2 == 0 { 
fmt.Println("Pile one =", first)
fmt.Println("Pile two =", second)
fmt.Println()
fmt.Println("Enter '1' or '2' to choose which pile.")
fmt.Scanln(&inputA)
if inputA == 1 {
  fmt.Println()
  fmt.Println("Now choose between 1 and 3 pebbles to remove from Pile 1.")
  fmt.Scanln(&inputB)
first -= inputB 
} else if inputA == 2 {
  fmt.Println()
  fmt.Println("Now choose between 1 and 3 pebbles to remove from Pile 2.")
  fmt.Scanln(&inputB)
second -= inputB
}
} else if x%2 !=0 {
  if rand.Intn(2) == 1 {
  
  first -= r1.Intn(3)
  
  
  } else if r1.Intn(2) == 2 {
    second -= r1.Intn(3)
  

  } else if r1.Intn(2) == 0 {
    
    first -= r1.Intn(3)
  }
}
}
   fmt.Println("Game Over")
}
