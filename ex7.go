package main

import "fmt"

func main(){
    i := 0
    prime := true
    curr_p := 2
    for j := 2; i < 10001; j++{
        for k := 2; k <= j -1; k++{
            if j % k == 0{
                prime = false
                break
            }
        }
        if prime{
            curr_p = j
            i += 1
        }
        prime = true
        fmt.Println(i)
    }
    fmt.Println(curr_p)
}
