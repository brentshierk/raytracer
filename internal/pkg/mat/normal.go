package mat

func NormalAt(s Shape, worldPoint Set, intersection *Intersection) Set {

	// transform point from world to object space, including recursively traversing any parent object
	// transforms.
	localPoint := WorldToObject(s, worldPoint)

	// normal in local space given the shape's implementation
	objectNormal := s.NormalAtLocal(localPoint, intersection)

	// convert normal from object space back into world space, again recursively applying any
	// parent transforms.
	return NormalToWorld(s, objectNormal)
}

//
//func NormalAtPtr(s Shape, worldPoint Tuple4, intersection *Intersection, localPoint *Tuple4) Tuple4 {
//
//	// transform point from world to object space, including recursively traversing any parent object
//	// transforms.
//	WorldToObjectPtr(s, worldPoint, localPoint)
//
//	// normal in local space given the shape's implementation
//	objectNormal := s.NormalAtLocal(*localPoint, intersection)
//
//	// convert normal from object space back into world space, again recursively applying any
//	// parent transforms.
//	NormalToWorldPtr(s, &objectNormal)
//	return objectNormal
//}

// in - normal * 2 * dot(in, normal)
func Reflect(vec Set, normal Set) Set {
	dotScalar := Dot(vec, normal)
	norm := Scalar(Scalar(normal, 2.0), dotScalar)
	return Subtract(vec, norm)
}

// in - normal * 2 * dot(in, normal)
func ReflectPtr(vec Set, normal Set, out *Set) {
	dotScalar := Dot(vec, normal)
	norm := Scalar(Scalar(normal, 2.0), dotScalar)
	SubPtr(vec, norm, out)
}