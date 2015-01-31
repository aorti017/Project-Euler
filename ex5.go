package main

import "fmt"

func main(){
    low := true
    for i := 2; i <= i; i++{
        for j:= 2; j <= 20; j++{
            if i % j != 0{
                low = false
                break
            }
        }
        if low{
            fmt.Println(i)
            return
        }
        low = true
    }
}
