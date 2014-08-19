Go09：Multiple results
======================

### 問題 ###

関数が複数の戻り値を返すことを確認する

### ソースコード ###


    //usr/bin/env go run $0 $@ ; exit
    
    package main
    
    import "fmt"
    
    func split(sum int) (x, y int) {
      x = sum * 4 / 9
      y = sum - x
      return 
    }
    
    func main() {
      fmt.Println(split(17))
    }


### 実行結果 ###

    $ ./09/main.go 
    world hello
    $


### 学習したこと ###

* Goの関数は、複数の値を返すことができる

