//usr/bin/env go run $0 $@ ; exit                                                                                                                                                     

package main

import "fmt"

func swap(x, y string) (string, string) {
  return y, x
}

func main() {
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}


