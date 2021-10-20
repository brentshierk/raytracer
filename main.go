package main

import (
	"fmt"
	"github.com/brentshierk/raytracer/internal/pkg/mat"
	"io/ioutil"
	"math"
	"os"

	//"io/ioutil"
	//"os"
)








func circleDemo() {
	c := mat.NewCanvas(100, 100)

	rayOrigin := mat.NewPoint(0, 0, -15.0)
	wallZ := 20.0
	wallSize := 7.0
	pixelSize := wallSize / float64(c.W)
	half := wallSize / 2
	color := mat.NewColor(1, 0, 0)
	sphere := mat.NewSphere()

	//mat.SetTransform(sphere, mat.Scale(1, 0.5, 1))
	//mat.SetTransform(sphere, mat.Multiply(mat.RotateZ(math.Pi/4), mat.Scale(0.5, 1, 1)))

	for row := 0; row < c.W; row++ {
		worldY := half - pixelSize*float64(row)

		for col := 0; col < c.H; col++ {
			worldX := -half + pixelSize*float64(col)
			posOnWall := mat.NewPoint(worldX, worldY, wallZ)

			rayFromOriginToPosOnWall := mat.NewRay(rayOrigin, mat.Normalize(mat.Sub(posOnWall, rayOrigin)))

			// check if our ray intersects the sphere
			intersections := mat.IntersectRayWithShape(sphere, rayFromOriginToPosOnWall)
			_, found := mat.Hit(intersections)
			if found {
				c.WritePixel(col, c.H-row, color)
			}
		}
	}
	// write
	data := c.ToPPM()
	err := ioutil.WriteFile("circle.ppm", []byte(data), os.FileMode(0755))
	if err != nil {
		fmt.Println(err.Error())
	}
}
