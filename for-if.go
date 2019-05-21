package main

import (
  "fmt"
)



// TODO: exercise
/* https://tour.golang.org/flowcontrol/8 */

// func pow()
/* https://tour.golang.org/flowcontrol/6 */

func deferOrder() {
  defer fmt.Println("world")

  fmt.Println("hello")

  defer fmt.Println("to")

  defer fmt.Println("my")
}

func ifFuncStr(num int) string {
  // num := 1
  limit := 5
  if num < limit {
    return fmt.Sprint(num)
  }
  return fmt.Sprint(limit)
  // return limit
}

func ifFunc(num int) int {
  // num := 1
  limit := 5
  if num < limit {
    return num
  }
  // return fmt.Printf(limit)
  return limit
}

func while() int {
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  return sum
}

func loopAlt() int {
  sum := 1
  for ; sum < 1000; {
    sum += sum
  }
  return sum
}

func loop() int {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  return sum
}

func main() {
  // fmt.Println(loop())
  // fmt.Println(loopAlt())
  // fmt.Println(while())
  // fmt.Println(ifFunc(2))
  // fmt.Println(ifFuncStr(2))

  deferOrder()
}
