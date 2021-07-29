package mat

import "fmt"


//type Set [4]float64

type Set [4]float64



func (s Set) Vector()  Set{
	s[0] = 0.0
	s[1] = 0.0
	s[2] = 0.0
	s[3] = 0.0
	return s
}
func (s Set) Point()  Set{
	s[0] = 0.0
	s[1] = 0.0
	s[2] = 0.0
	s[3] = 1.0
	return s

}


// NewVector creates a new vector
func NewVector(x,y,z float64) Set{
	return Set{x,y,z,0.0}
}
// NewPoint creates new point
func NewPoint(x,y,z float64) Set{
	return Set{x,y,z,1.0}
}
func (s Set) IsVector()  bool{
	return s[3] == 0.0
}
func (s Set) IsPoint()  bool{
	return s[3] == 1.0
}
func (s Set) Get(i int)  float64{
	return s[i]
}

func Add(a,b Set) Set{
	newSet := Set{}
	for i :=0;i<4;i++{
		newSet[i] = a[i] +b[i]
	}
	return newSet
}

func Subtract(s1, s2 Set) Set {
	s3 := Set{}
	for i := 0; i < 4; i++ {
		s3[i] = s1[i] - s2[i]
	}
	return s3
}
// im subtracting each index of p,p1 p[3] has not changed "still 1.0"
func SubPointFromPoint(p,p1 Set)  Set{
	newSet := Subtract(p,p1)
	newSet[3] = 0.0
	//now a vector
	return newSet
}
func SubVectorFromPoint(p,v Set)  Set{
	 newSet := Subtract(p,v)
	newSet[3] = 1.0
	//now a point
	return newSet
}
func SubVectorFromVector(v,v1 Set)  Set{
	newSet := Subtract(v,v1)
	newSet[3] = 0.0
	return newSet
}
func NegateVector(v Set)  Set{
	zeroVector := NewVector(0,0,0)
	newSet := Subtract(zeroVector,v)
	return newSet
}
func NegateSet(a Set ) Set{
	zeroSet := Set{}
	negativeSet := Subtract(zeroSet,a)
	return negativeSet
}





func tuple(){
	x := NewVector(3.3,4.2,1.7)
	y := NewPoint(2.2,5.1,-0.7)
	n := Add(x,y)
	fmt.Println(n)
}