package main

import (
	"fmt"
	"github.com/brentshierk/raytracer/mat"

)


func main(){

	x := mat.NewVector(3.3,4.2,1.7)
	y := mat.NewPoint(2.2,5.1,-0.7)
	n := mat.Add(x,y)
	b := mat.NegateSet(y)
	fmt.Println(b)
	fmt.Println(n)
}