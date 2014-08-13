//usr/bin/env go run $0 $@ ; exit                                                                                                                                                     

package main

import "fmt"
import "reflect"

// 関数と型を宣言する
func add(x int, y int) int {
  return x + y
}

// 42+13=55
func main() {

  // 型を表示する
  fmt.Println(reflect.TypeOf(add))

  // 関数を実行し、標準出力に表示する
  fmt.Println(add(42, 13))

}

