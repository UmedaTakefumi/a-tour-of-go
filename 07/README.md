Go07：Functions
======================

### 問題 ###

型と関数を作成し

### ソースコード ###

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

### 実行結果 ###

    $ ./07/main.go 
    55
    $ 

### 学習したこと ###

* 関数は0個以上の引数を取ることができる
* 型は、変数名のあとにくる
* なぜ型の記述が、変数名のあとに記述する理由については下記のURLを参照 
* [Go's Declaration Syntax - The Go Blog](http://blog.golang.org/gos-declaration-syntax)

