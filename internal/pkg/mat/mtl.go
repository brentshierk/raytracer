package mat

type Mtl struct {
	Ambient         Set
	Diffuse         Set
	Specular        Set
	Shininess       float64
	Reflectivity    float64
	Transparency    float64
	RefractiveIndex float64
	Name            string
}