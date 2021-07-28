package mat

import "fmt"


//type Set [4]float64

type Set [4]float64



func (s Set) Vector()  Set{
	s[0] = 0.0
	s[1] = 0.0
	s[2] = 0.0
	s[3] = 1.0
	return s
}
func (s Set) Point()  Set{
	s[0] = 0.0
	s[1] = 0.0
	s[2] = 0.0
	s[3] = 1.0
	return s

}

func NewVector(x,y,z float64) Set{
	return Set{x,y,z,1.0}
}

func NewPoint(x,y,z float64) Set{
	return Set{x,y,z,0.0}
}

func Add(a,b Set) Set{
	newSet := Set{}
	for i :=0;i<4;i++{
		newSet[i] = a[i] +b[i]
	}
	return newSet
}

func tuple(){
	x := NewVector(3.3,4.2,1.7)
	y := NewPoint(2.2,5.1,-0.7)
	n := Add(x,y)
	fmt.Println(n)
}