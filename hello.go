package main

import "fmt"

func main() {
    fmt.Println("Hello World")
    
    for i := 1 ; i < 10 ; i++ {
        fmt.Println(i)
    } 
    
    for i := 1 ; i < 9 ; i++ {
        for j := 1 ; j < 9 ; j++ {
            print(i*j)
        }
    }
}