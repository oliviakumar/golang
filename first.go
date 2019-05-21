package main

import (
  "fmt"
  "time"
  "math/rand"
  // "math"
)

// func type() string {
//     v := 42
//   	// return
//     v := 42
//     fmt.Printf("v is of type %T\n", v)
// }

func swap(x, y string) (string, string) {
  return y, x
}

func add(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println("hello")
  fmt.Println("the time is", time.Now())
  fmt.Println("here's a number", rand.Intn(10))

  z := add(5, 2)
  fmt.Println(z)
  // fmt.Println(add(5, 2))

  // fmt.Println(swap("hello", "world"))
  a, b := swap("hello", "world")
  fmt.Println(a, b)

  /*https://tour.golang.org/basics/7 -- named/naked return vals */

  // type()
  v := 42
  fmt.Printf("v is of type %T\n", v)
  p := "hi"
  fmt.Printf("p is of type %T\n", p)
}
