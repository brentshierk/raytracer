package main

import (
	"fmt"
	"github.com/brentshierk/raytracer/mat"
	"io/ioutil"
	"os"
)


func main(){

projectileDemo()


}
func projectileDemo() {
	VelocityVector := mat.NewVector(1,1.8,0)
	NormVelocityVector := mat.NormalizeMagnitude(VelocityVector)
	v:= mat.Scaler(5.2,NormVelocityVector)
	p := mat.NewPoint(0,1,0)

	prj := NewProjectile(p,v)
	env := NewEnvironment(mat.NewVector(0, -0.1, 0), mat.NewVector(-0.01, 0, 0))
	c := mat.NewCanvas(900, 550)
	red := mat.NewColor(1, 1, 1)
	for prj.pos.Get(1) > 0.0 {
		tick(prj, env)
		//time.Sleep(time.Millisecond * 100)
		fmt.Printf("Projectile pos %v at height %v with velocity %v\n", mat.VectorMagnitude(prj.pos), prj.pos.Get(1), prj.velocity)
		fmt.Printf("Drawing at: %d %d\n", int(prj.pos.Get(0)), c.H-int(prj.pos.Get(1)))
		c.WritePixel(int(prj.pos.Get(0)), c.H-int(prj.pos.Get(1)), red)
	}
	fmt.Printf("Projectile flew %v\n", mat.VectorMagnitude(prj.pos))
	data := c.ToPPM()
	err := ioutil.WriteFile("pic.ppm", []byte(data), os.FileMode(0755))
	if err != nil {
		fmt.Println(err.Error())
	}
}
func tick(prj *Projectile, env *Environment) {
	prj.pos = mat.Add(prj.pos, prj.velocity)
	prj.velocity = mat.Add(prj.velocity, env.gravity)
	prj.velocity = mat.Add(prj.velocity, env.wind)
}

type Environment struct {
	gravity mat.Set
	wind    mat.Set
}

func NewEnvironment(gravity mat.Set, wind mat.Set) *Environment {
	return &Environment{gravity: gravity, wind: wind}
}

type Projectile struct {
	pos      mat.Set
	velocity mat.Set
}

func NewProjectile(pos mat.Set, velocity mat.Set) *Projectile {
	return &Projectile{pos: pos, velocity: velocity}
}