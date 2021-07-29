package mat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestSet_IsPoint(t *testing.T) {
	p := NewPoint(2,3,4)
	assert.True(t, p.IsPoint())
	assert.False(t, p.IsVector())

	assert.Equal(t, 2.0, p.Get(0))
	assert.Equal(t, 3.0, p.Get(1))
	assert.Equal(t, 4.0, p.Get(2))
}

func TestSet_IsVector(t *testing.T) {
	v := NewVector(4.3, -4.2, 3.1)
	assert.True(t, v.IsVector())
	assert.False(t, v.IsPoint())

	assert.Equal(t, 4.3, v.Get(0))
	assert.Equal(t, -4.2, v.Get(1))
	assert.Equal(t, 3.1, v.Get(2))
}

func TestSet_Adding(t *testing.T) {
	p := NewPoint(3,5,7)
	v := NewVector(2,1,0)
	t1 := Add(p,v)

	assert.Equal(t,5.0,t1.Get(0))
	assert.Equal(t,6.0,t1.Get(1))
	assert.Equal(t, 7.0,t1.Get(2))
}

func TestSet_Subtract(t *testing.T) {
	p := NewPoint(1,2,3)
	v := NewVector(3,2,1)
	t1 := Subtract(p,v)
	assert.Equal(t, -2.0,t1.Get(0))
	assert.Equal(t, 0.0,t1.Get(1))
	assert.Equal(t, 2.0,t1.Get(2))

}
//func TestSet_SubVectorFromPoint(t *testing.T) {
//	v := NewVector(2,3,7)
//	p := NewPoint(3,6,5)
//
//	t1 := Subtract(v,p)
//	assert.Equal(t, -1.0,t1.Get(0))
//	assert.Equal(t, -3.0,t1.Get(1))
//	assert.Equal(t, 2.0,t1.Get(2))
//	assert.Equal(t, 1.0,t1.Get(3))
//}
func TestSet_SubVectorFromPoint(t *testing.T) {
	v := NewVector(2,3,7)
	p := NewPoint(3,6,5)

	t1 := Subtract(p,v)
	assert.Equal(t, 1.0,t1.Get(0))
	assert.Equal(t, 3.0,t1.Get(1))
	assert.Equal(t, -2.0,t1.Get(2))
	assert.Equal(t, 1.0,t1.Get(3))
}
