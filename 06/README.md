Go06：Exported names
======================

### 問題 ###

math.piと記述されたコードをmath.Piと書き換えてコードを実行する

### ソースコード ###

    //usr/bin/env go run $0 $@ ; exit                                                                                                                                                     
    
    package main
    
    import (
      "fmt"
      "math"
    )
    
    func main() {
      fmt.Println(math.Pi)
    }

### 実行結果 ###

    $ ./06/main.go 
    3.141592653589793
    $ 

### 学習したこと ###

* パッケージをインポートすると外部に公開（エクスポート）している関数を参照することができる。
* Goでは、最初の文字が大文字で始まる関数の場合は、エクスポートされている。


