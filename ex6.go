package main

import "fmt"

func main(){
    pre_square := 0
    post_square :=0

    for i := 1; i <= 100; i++{
        pre_square += i*i
        post_square += i
    }

    post_square *= post_square

    fmt.Println(post_square-pre_square)
}
