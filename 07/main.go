//usr/bin/env go run $0 $@ ; exit                                                                                                                                                     

package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

// 42+13=55
func main() {
  fmt.Println(add(42, 13))

}

