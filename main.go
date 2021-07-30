package main

import (
	"fmt"
	"github.com/brentshierk/raytracer/mat"

)


func main(){

	x := mat.NewVector(1,2,3)
	y := mat.NewVector(1,2,3)
	n := mat.Dot(x,y)
	fmt.Println(n)


}