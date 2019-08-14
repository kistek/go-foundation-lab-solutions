package main

import "fmt"

type iGeoshape interface {
	draw()
	print()
}

type Rectangle struct{}

func (r Rectangle) draw() {
	fmt.Println("drawing the rectangle")
}

func (r Rectangle) print() {
	fmt.Println("printing the rectangle. This may look symmetrical on rectangular paper.")
}

type Circle struct{}

func (r Circle) draw() {
	fmt.Println("drawing the circle")
}

func (r Circle) print() {
	fmt.Println("printing the circle. This may look sophisticated on rectangular paper.")
}

func render(shapes ...iGeoshape) {
	for _, shape := range shapes {
		shape.draw()
	}
}

func main() {
	var r1, r2, r3 Rectangle
	var c1, c2 Circle

	render(r1, r2, r3, c1, c2)
}
