Go04：Package
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

### 学習したこと ###

* 処理はfunc main()から始まる。
* すべてのGoプログラムは、パッケージで構成されている。
* 規約でパッケージ名は、インポートパスの最後の要素と同一となる。
  * go.net/websocketだった場合は、websocketとなる。
  * 大事なことなのでもう一度、インポートパスとパッケージ名は別物

