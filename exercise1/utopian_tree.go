package main

import "fmt"
import "math"

func main(){
    var height int = 1

    fmt.Println("Cycles -> Result")

    for i := 0; i < 5; i++ { 
        if i==0 || i == 1 || i == 4 {
             fmt.Println(i, " -> ", height)
        }    
        if math.Mod(float64(i), 2) == 0{
             height = height * 2 
        } else {
             height = height + 1 
         }
    }
}
