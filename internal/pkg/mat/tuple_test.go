package mat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestSet_IsPoint(t *testing.T) {
	p := Point(2,3,4)
	if p.w != 1.0{
		t.Errorf("p.w = %b; is a NOT a point",p.w)
	}
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
	p := Point(3,5,7)
	v := Vector(2,1,0)
	n := Set{}
	n.x = p.x + v.x
	n.y = p.y + v.y
	n.z = p.z + v.z
	return
}
