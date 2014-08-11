Go05：imports
======================

### 問題 ###

複数のパッケージをimportで指定する

### ソースコード ###


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

### 実行結果 ###

    $ 05/main.go 
    Now you have 2.0000000000000004 problems.
    $ 

### 学習したこと ###

* import文に複数のパッケージを記述することができる

