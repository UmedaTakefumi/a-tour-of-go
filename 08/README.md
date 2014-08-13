Go08：Functions continued
======================

### 問題 ###

関数宣言時に設定する引数に同一の型の記述を省略し、正しく関数が実行されることを確認する

### ソースコード ###

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

### 実行結果 ###

    $ ./08/main.go 
    55
    $ 

### 学習したこと ###

* 関数の引数が2つ以上の同じ型であれば、最後の型を残して省略することができる


    x int, y int
    
    　↓ 次のように省略
    
    x, y int


