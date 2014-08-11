//usr/bin/env go run $0 $@ ; exit                                                                                                                                                     

package main

import (
  "fmt"
  "math/rand"
)

func main() {
  fmt.Println("My favortite number is", rand.Intn(10))
}

