package main

import (
	"fmt"
	"github.com/brentshierk/raytracer/internal/pkg/mat/"

)


func main(){

	var x = NewVector(3.3,4.2,1.7)
	y := NewPoint(2.2,5.1,-0.7)
	n := Add(x,y)
	fmt.Println(n)
}