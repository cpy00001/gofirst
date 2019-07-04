package main

import (
    "fmt"
    "os"
)

func test()  {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)


    var a, c int
    var d string
    a, c, d = 5, 7, "abc"
    fmt.Println(a,c,d)
}


//一个可以返回多个值的函数
func numbers()(int,int,string){
    a , b , c := 1 , 2 , "str"
    return a,b,c
}

func g(){
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        sep = " "
        s += sep + os.Args[i]
        fmt.Println(s)
    }
    fmt.Println(s)
}

func main() {
    //test()
    //fmt.Println(numbers())
    g()
}