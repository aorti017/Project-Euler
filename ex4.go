package main

import (
    "fmt"
    "strconv"
)

func main(){
    great := 0
    for i := 0; i < 1000; i++{
        for j := 0; j < 1000; j++{
            var num = i * j
            var num_str = strconv.Itoa(num)
            if palindrome(num_str){
                if num > great{
                    great = num
                }
            }
        }
    }
    fmt.Println(great)
}

func palindrome(num string) bool {
    var str = ""
    for i := len(num) - 1; i >= 0; i--{
        str = str + string(num[i])
    }
    if str == num{
        return true
    }
    return false
}
