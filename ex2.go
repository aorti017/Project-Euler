package main

import "fmt"

func main() {
    fib(2, 1, 0)
}

func fib(curr int, prev int, sum int){
    if curr % 2 == 0{
        sum += curr
    }
    if curr + prev > 4000000{
        fmt.Println(sum)
        return
    }
    temp := prev
    prev = curr
    curr += temp
    fib(curr, prev, sum)
}

