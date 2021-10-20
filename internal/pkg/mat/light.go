package mat

import (
	"math"
	"math/rand"
)

type LightSource interface {
	Pos() Set
	Intens() Set
}

type Light struct {
	Position  Set
	Intensity Set
}

func NewLight(position Set, intensity Set) Light {

	return Light{
		Position:  position,
		Intensity: intensity,
	}
}

func (l Light) Pos() Set {
	return l.Position
}

func (l Light) Intens() Set{
	return l.Intensity
}

type AreaLight struct {
	Corner    Set
	UVec      Set
	USteps    int
	VVec      Set
	VSteps    int
	Intensity Set
	Samples   float64
	Position  Set
}

func (al AreaLight) Pos() Set {
	return al.Position
}

func (al AreaLight) Intens() Set{
	return al.Intensity
}

func OrientAreaLight(light *AreaLight, source Set, target Set) {
	n := Normalize(Subtract(target, source)) // Desired direction of the area light normal

	// Compute tangent and bitangent vectors
	a := NewVector(0, 1, 0)
	t := Normalize(Cross(a, n))
	b := Cross(t, n)

	// Replace the uvec and vvec vectors, but preserve their length
	light.UVec = Scalar(t, Magnitude(light.UVec))
	light.VVec = Scalar(b, Magnitude(light.VVec))

	// Set the corner so that the source position falls in the center
	light.Corner = Subtract(source, Subtract(Scalar(light.UVec, 0.5), Scalar(light.VVec, 0.5)))
}

func NewAreaLight(corner Set, uVec Set, usteps int, vVec Set, vsteps int, intensity Set) AreaLight {
	return AreaLight{
		Corner:    corner,
		UVec:      DivideByScalar(uVec, float64(usteps)),
		USteps:    usteps,
		VVec:      DivideByScalar(vVec, float64(vsteps)),
		VSteps:    vsteps,
		Intensity: intensity,
		Samples:   float64(usteps * vsteps),
		Position: Add(corner, NewPoint(
			(uVec[0]+vVec[0])/2,
			(uVec[1]+vVec[1])/2,
			(uVec[2]+vVec[2])/2)),
	}
}

func PointOnLight(light AreaLight, u, v float64) Set {
	return Add(light.Corner,
		Add(
			Scalar(light.UVec, u+rand.Float64()*0.5),
			Scalar(light.VVec, v+rand.Float64()*0.5)))
}
func PointOnLightNoRandom(light AreaLight, u, v float64) Set {
	return Add(light.Corner,
		Add(
			Scalar(light.UVec, u+0.5),
			Scalar(light.VVec, v+0.5)))
}

func Lighting(material Material, object Shape, light LightSource, position, eyeVec, normalVec Set, intensity float64, lightData LightData) Set {
	var color Set

	if material.HasPattern() {
		color = PatternAtShape(material.Pattern, object, position)
	} else {
		color = material.Color
	}
	if intensity == 0.0 {
		MultiplyByScalarPtr(color, material.Ambient, &lightData.EffectiveColor)
		return lightData.EffectiveColor
	}

	HadamardPtr(&color, light.Intens(), &lightData.EffectiveColor)

	// sample each point on area light
	l := light.(AreaLight)

	sum := NewColor(0, 0, 0)
	for u := 0; u < l.USteps; u++ {
		for v := 0; v < l.VSteps; v++ {

			// get vector from point on sphere to light source by subtracting, normalized into unit space.
			//p := light.(AreaLight).Corner // renders OK?
			p := PointOnLight(l, float64(u), float64(v)) // ???
			//p := PointOnLightNoRandom(l, float64(u), float64(v)) // Works with unit test

			SubPtr(p, position, &lightData.LightVec)
			NormalizePtr(&lightData.LightVec, &lightData.LightVec)

			// Add the ambient portion
			MultiplyByScalarPtr(lightData.EffectiveColor, material.Ambient, &lightData.Ambient)

			lightDotNormal := Dot(lightData.LightVec, normalVec)

			// get dot product (angle) between the light and normal  vectors. If negative, it means the light source is
			// on the backside
			if lightDotNormal < 0 {
				lightData.Diffuse = black
				lightData.Specular = black
			} else {
				// Diffuse contribution first here??

				MultiplyByScalarPtr(lightData.EffectiveColor, material.Diffuse*lightDotNormal, &lightData.Diffuse)

				// reflect_dot_eye represents the cosine of the angle between the
				// reflection vector and the eye vector. A negative number means the
				// light reflects away from the eye.
				// Note that we negate the light vector since we want to angle of the bounce
				// of the light rather than the incoming light vector.

				ReflectPtr(Negate(lightData.LightVec), normalVec, &lightData.ReflectVec)
				reflectDotEye := Dot(lightData.ReflectVec, eyeVec)

				if reflectDotEye <= 0.0 {
					lightData.Specular = black
				} else {
					// compute the specular contribution
					factor := math.Pow(reflectDotEye, material.Shininess)

					// again, check precedense here
					MultiplyByScalarPtr(light.Intens(), material.Specular*factor, &lightData.Specular)
				}
			}
			sum = Add(sum, Add(lightData.Diffuse, lightData.Specular))
		}
	}

	// Add the three contributions together to get the final shading
	// Uses standard Tuple addition
	// for soft shadows, multiply by intensity
	return lightData.Ambient.Add(Scalar(DivideByScalar(sum, l.Samples), intensity))
}

// LightingPointLight computes the color of a given pixel given phong shading a point light
func LightingPointLight(material Material, object Shape, light Light, position, eyeVec, normalVec Set, inShadow bool, lightData LightData) Set {
	var color Set
	if material.HasPattern() {
		color = PatternAtShape(material.Pattern, object, position)
	} else {
		color = material.Color
	}
	if inShadow {
		MultiplyByScalarPtr(color, material.Ambient, &lightData.EffectiveColor)
		return lightData.EffectiveColor
	}

	HadamardPtr(&color, light.Intensity, &lightData.EffectiveColor)

	// get vector from point on sphere to light source by subtracting, normalized into unit space.
	SubPtr(light.Position, position, &lightData.LightVec)
	NormalizePtr(&lightData.LightVec, &lightData.LightVec)

	// Add the ambient portion
	MultiplyByScalarPtr(lightData.EffectiveColor, material.Ambient, &lightData.Ambient)

	// get dot product (angle) between the light and normal  vectors. If negative, it means the light source is
	// on the backside
	lightDotNormal := Dot(lightData.LightVec, normalVec)
	specular := lightData.Specular
	diffuse := lightData.Diffuse

	if lightDotNormal < 0 {
		diffuse = black
		specular = black
	} else {
		// Diffuse contribution Precedense here??

		MultiplyByScalarPtr(lightData.EffectiveColor, material.Diffuse*lightDotNormal, &diffuse)

		// reflect_dot_eye represents the cosine of the angle between the
		// reflection vector and the eye vector. A negative number means the
		// light reflects away from the eye.
		// Note that we negate the light vector since we want to angle of the bounce
		// of the light rather than the incoming light vector.

		ReflectPtr(Negate(lightData.LightVec), normalVec, &lightData.ReflectVec)
		reflectDotEye := Dot(lightData.ReflectVec, eyeVec)

		if reflectDotEye <= 0.0 {
			specular = black
		} else {
			// compute the specular contribution
			factor := math.Pow(reflectDotEye, material.Shininess)

			// again, check precedense here
			MultiplyByScalarPtr(light.Intensity, material.Specular*factor, &specular)
		}
	}
	// Add the three contributions together to get the final shading
	// Uses standard Tuple addition
	return lightData.Ambient.Add(diffuse.Add(specular))
}