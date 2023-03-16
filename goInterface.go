package main

import (
    "fmt"
    "math"
)
//Interface
type geometry interface {
    area() float64
    perim() float64
}

//rectangle
type rect struct {
    width, height float64
}

//circle
type circle struct {
    radius float64
}

//area of rectangle
func (r rect) area() float64 {
    return r.width * r.height
}

//area of circle
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

//perimeter of rectangle
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

//perimeter of circle
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}


func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 5, height: 4}
    c := circle{radius: 7}

    measure(r)
    measure(c)
}