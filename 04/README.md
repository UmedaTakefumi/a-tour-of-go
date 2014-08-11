Go01：Hello, 世界
======================

### 問題 ###

math/rand Packageのrand.Intn関数を利用する


### ソースコード ###

    //usr/bin/env go run $0 $@ ; exit                                                                                                                                                     
    
    package main
    
    import (
      "fmt"
      "math/rand"
    )
    
    func main() {
      fmt.Println("My favortite number is", rand.Intn(10))
    }

### 実行結果 ###

    $ ./main.go 
    My favortite number is 1
    $ 

