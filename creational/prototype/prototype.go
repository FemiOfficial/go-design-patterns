package prototype

import (
	"fmt"
	"math"
)

type ShapeType string;

type Shape interface {
	ShapeType() ShapeType
	Clone() Shape
	Display()
}

const (
	CircleShape = "circle"
	RectangleShape = "rectangle"
)

type Circle struct {
	Radius float32
}

type Rectangle struct {
	Length float32
	Width float32
}

func (c Circle) ShapeType() ShapeType {
	return CircleShape
}

func GetCirclePrototype(r float32) Circle {
	return Circle{Radius: r}
}

func (c Circle) Display() {
	area := math.Pi * c.Radius * c.Radius
	fmt.Printf("Circle with radius %f and area %f %s %v\n", c.Radius, area, c.ShapeType(), c)
}

func (c Circle) Clone() Shape {
	return GetCirclePrototype(c.Radius)
}

func GetRectanglePrototype(l float32, w float32) Rectangle {
	return Rectangle{Length: l, Width: w}
}

func (r Rectangle) ShapeType() ShapeType {
	return RectangleShape
}

func (r Rectangle) Display() {
	area := r.Length * r.Width
	fmt.Printf("Rectangle with length %f and width %f and area %f %s %v\n", r.Length, r.Width, area, r.ShapeType(), r)
}

func (r Rectangle) Clone() Shape {
	return GetRectanglePrototype(r.Length, r.Width)
}


var cloneRegistry = make(map[ShapeType]Shape)

func init() {
	cloneRegistry[CircleShape] = GetCirclePrototype(0)
	cloneRegistry[RectangleShape] = GetRectanglePrototype(0, 0)
}
