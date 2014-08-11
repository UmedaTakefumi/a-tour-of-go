//usr/bin/env go run $0 $@ ; exit                                                                                                                                                     

package main

import (
  "fmt"
  "math"  
)

func main() {
  fmt.Printf("Now you have %g problems.\n",
    math.Nextafter(2, 3))
}


