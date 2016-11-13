package main

import (
  "fmt"
)

var a [3]int = [3]int{1, 2, 3}  // array of 3 integers

func main(){
fmt.Println(a[0])

for i, v := range a {
  fmt.Printf("%d %d\n", i, v)
}

for _, v := range a {
  fmt.Printf("%d\n", v)
}
}
