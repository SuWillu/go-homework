package main

import "math"

func utopiantree(cycles int) int{
    var height int = 1
    var growth int = 0

    for growth < cycles { 
        if math.Mod(float64(growth), 2) == 0{
             height = height * 2 
        } else {
             height = height + 1 
        }
        growth++
    }

    return height;
}

