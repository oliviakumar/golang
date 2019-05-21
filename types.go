package main

import (
  "fmt"
)


// todo
/* https://tour.golang.org/moretypes/5 */

type Vertex struct {
  X int
  Y int
}

func slicerRet() var []int {
  primes := [6]int{2, 3, 5, 7, 11, 13}
  var s []int = primes[1:4]
  return s//fmt.Println(s)
}

// func slice(array []primes) {
func slicer() {
  primes := [6]int{2, 3, 5, 7, 11, 13}
  var s []int = primes[1:4]
  fmt.Println(s)
}

func arrays(x string) {//var int {// var {
  // if x == "a" {
  var a [2]string
  a[0] = "hello"
  a[1] = "world"
  fmt.Println(a)
    // return a
  // }
  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)
  // return primes
}

func types(v Vertex) Vertex {
  v.X = 4
  return v
}

func main() {
  v := Vertex{1, 2}
  fmt.Println(v.X)
  v = types(v) // for fun: try a global static-y Vertex v
  fmt.Println(v.X)

  // fmt.Println(arrays("a"))
  // fmt.Println(arrays("b"))
  arrays("b")


  // fmt.Println(slicer())
  slicer()
  // fmt.Println(slicerRet())
}
