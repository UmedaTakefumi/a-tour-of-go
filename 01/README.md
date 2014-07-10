Go Hello World
======================

ここのツアーで学習できること

* 基本コンセプト
* メソッドとインターフェース
* 並行性

Go01：Hello, 世界
------
### 問題 ###
コンソールにHello, 世界と表示する

### ソースコード ###

    //usr/bin/env go run $0 $@ ; exit                                                                                                                                                     
    
    package main
    
    import "fmt"
    
    func main() {
      fmt.Printf("Hello, 世界\n")
    }

### 実行結果 ###

    $ ./hello.go 
    Hello, 世界
    $ 



