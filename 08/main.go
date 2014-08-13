//usr/bin/env go run $0 $@ ; exit                                                                                                                                                     

package main

import "fmt"

func add(x, y int) int {
  return x + y
}

func main() {

  // 42+13=55
  fmt.Println(add(42, 13))

}

