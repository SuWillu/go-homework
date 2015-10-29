package main

import "fmt"
func main(){
	PrintNumber(1)
	PrintNumber(2)
	PrintNumber(3)
}

func PrintNumber(number int){
        fmt.Println(number)
	time.Sleep(time.Second)
}
