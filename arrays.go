package main

import "fmt"

func main() {
  arr2()
}

func arr1 () {
  names := [...]string{
    "John Doe",
    "Jane Doe",
    "don tito",
    "Harry Potter",
    "Iron man"}

  for index, name := range names {
    fmt.Printf("Index no of %s is %d\n", name, index)
  }
}

func arr2 () {
   s := make([]int, 10)
   s[0] = 1
   s[1] = 3
   s[2] = 5
   s[3] = 7
   n := append(s, 9)
   m := append(n, 10)
   fmt.Println(s)
   fmt.Println(n)
   fmt.Println(m) //[1 3 5 7 0 0 0 0 0 0 9]
}
