package main

import (
    "fmt"
    "math/rand"
    "time"
)

func testFor1() {
    var i int
    for i = 1; i < 10; i++ {
        fmt.Printf("i=%d\n", i)
    }
    fmt.Printf("final:i=%d\n", i)
}

func testFor2() {
    var i int
    for i = 1; i < 10; i++ {
        fmt.Printf("i=%d\n", i)
        if i > 5 {
            break
        }
    }
    fmt.Println(i)
}

//打印奇数
func testFor3() {
    var i int
    for i = 1; i < 1000; i++ {
        //正整数，就调出本次循环，所以不打印
        if i%2 == 0 {
            continue
        }
        fmt.Printf("i=%d\n", i)
    }

}
func testFor4() {
    i := 1
    for i <= 10 {
        fmt.Printf("i=%d\n", i)
        i = i + 2
    }
}

func testFor5() {
    i := 1
    for i <= 10 {
        fmt.Printf("i=%d\n", i)
        i = i + 2
    }
}

func testMultiSign() {
    a, b, c := 10, "hello", 100
    fmt.Printf("a=%d b=%s c=%d\n", a, b, c)
}

func testFor6() {
    for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
        fmt.Printf("%d*%d=%d\n", no, i, no*i)
    }
}

func testFor7() {
    for {
        fmt.Printf("hello\n")
        time.Sleep(time.Second)
    }
}

//峰云大神-http://xiaorui.cc/2016/03/23/golang%E9%9A%8F%E6%9C%BAtime-sleep%E7%9A%84duration%E9%97%AE%E9%A2%98/
func fengyun() {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 10; i++ {
        x := rand.Intn(10)
        fmt.Println(x)
        time.Sleep(time.Duration(x) * time.Second)
    }
}

//入口执行函数
func main() {
    //testFor1()
    //testFor2()
    //testFor3()
    //testFor4()
    //testFor5()
    //testFor6()
    testFor7()
    //fengyun()
}
