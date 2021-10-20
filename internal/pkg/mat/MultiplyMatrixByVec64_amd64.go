//+build !noasm
//+build !appengine

package mat

import "unsafe"

//go:noescape
func __MultiplyMatrixByVec64(m, vec4, result unsafe.Pointer)

func MultiplyByTuplePtr(m *Mat4x4, f2 *Set, _f4 *Set) {
	__MultiplyMatrixByVec64(unsafe.Pointer(m), unsafe.Pointer(f2), unsafe.Pointer(_f4))
}