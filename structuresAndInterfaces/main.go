package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
	newSize() (float64, float64)
	getName() string
}

type Rect struct {
	width, height float64
	name          string
}
type Circle struct {
	radius, diameter float64
	name             string
}

func (r *Rect) area() float64 {
	return r.width * r.height
}
func (r *Rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func (c *Circle) newSize() (float64, float64) {
	c.radius = 25
	c.diameter = c.radius * 2
	return c.radius, c.diameter
}
func (r *Rect) newSize() (float64, float64) {
	r.width = 25
	r.height = 40
	return r.width, r.height
}

func (r *Rect) getName() string {
	return r.name
}
func (c *Circle) getName() string {
	return c.name
}

func calculation(shapes ...Shape) {
	var tA float64
	for _, s := range shapes {
		tA += s.area()
		printInformation(s)
	}
	fmt.Println("total area of all shapes = ", tA)
}

func printInformation(s Shape) {
	fmt.Printf("%#v\n", s)
	fmt.Println(s.getName(), " area = ", s.area())
	fmt.Println(s.perimeter())
	fmt.Println(s.newSize())
	fmt.Printf("\"new Size \"%#v\n", s)
	fmt.Println("----------------------------------")
}

func main() {
	r := Rect{width: 3, height: 4, name: "rect"}
	c := Circle{radius: 5, diameter: 5 * 2, name: "circle"}
	calculation(&r, &c)
} 
/       git tested
