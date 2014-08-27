Go10：Multiple results
======================

### 問題 ###

戻り値パラメータに名前をつけてreturnステートメントに戻り値変数名を記載せずに関数を実行する

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

    $ ./10/main.go 
    7 10
    $ 

### 学習したこと ###

* Goの関数は、戻りを複数返すことができる
* 複数の戻り値は、名前をつけて変数のように扱うことができる
* 戻り値パラメータに名前が付けられている場合は、return文に名前を書かなくてもよい



