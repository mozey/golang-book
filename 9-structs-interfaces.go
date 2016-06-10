package main

import (
	"fmt"; "math"
)

// Structs .....................................................................
// The type keyword is used to create a new type,
// Circle is the name of this struct type.
type Circle struct {
	// Each field of the struct has a name and a type.
	x float64
	y float64
	r float64
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a * a + b * b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func setCircleRadius(c *Circle, area float64) {
	c.r = math.Sqrt(area / math.Pi)
}

// Methods .....................................................................
// Are a special type of function,
// they are associated with a type.
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}

func (r *Rectangle) perimeter() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return 2 * (l + w)
}

// Embedded Types ..............................................................
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

//type Android struct {
//	Person Person
//	Model  string
//}

type Android struct {
	Person
	Model string
}

// Interfaces ..................................................................

type Shape interface {
  area() float64
}

func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}

// Interfaces can also be used as fields!
type MultiShape struct {
  shapes []Shape
}

func (m *MultiShape) area() float64 {
  var area float64
  for _, s := range m.shapes {
    area += s.area()
  }
  return area
}


func main() { // ...............................................................
	// This creates a local Circle variable with all fields set to zero:
	// 		float		=> 0.0
	// 		int			=> 0
	// 		string		=> ""
	// 		pointer 	=> nil
	//var c Circle

	// The new function allocates memory for all the fields,
	// sets them to zero and returns a pointer *Circle
	//c := new(Circle)

	// We can also create a struct and assign values at the same time
	//c := Circle{x: 0, y: 0, r: 5}
	// Or omit the field names if the values are passed in order
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}

	// Using a struct, as below, is superior to passing co-ordinates
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	fmt.Println("rectangleArea", rectangleArea(rx1, ry1, rx2, ry2))

	// Remember that arguments are always copied in go
	fmt.Println("circleArea", circleArea(c))

	// We have to pass a pointer to assign values
	// to a field of Circle inside a function.
	setCircleRadius(&c, 80)
	fmt.Println("circleArea", circleArea(c))

	// Type methods allow us to use the . operator to call them,
	// we also don't need the & operator, since go knows to pass a pointer.
	fmt.Println("circleArea", c.area())

	// Go supports relationships like "android is a person"
	// by using an embedded type, a.k.a anonymous fields
	a := new(Android)
	a.Name = "Slim Shady"
	a.Person.Talk()
	a.Person.Name = "The real Slim Shady"
	// People can talk, an android is a person,
	// therefore an android can talk.
	a.Talk()

	fmt.Println("totalArea", totalArea(&c, &r))

	fmt.Println("circle perimeter", c.perimeter())
	fmt.Println("rect perimeter", r.perimeter())
}