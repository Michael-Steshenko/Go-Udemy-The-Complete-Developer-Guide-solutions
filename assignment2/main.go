package main

import "fmt"

type shape interface {
	getArea() float64
}

type triagle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (myTriangle triagle) getArea() float64 {
	return 0.5 * myTriangle.base * myTriangle.height
}

func (mySquare square) getArea() float64 {
	return mySquare.sideLength * mySquare.sideLength
}

func printArea(myShape shape) {
	fmt.Println(myShape.getArea())
}

// main is added to test the program more easily
func main() {
	tri := triagle{5, 3}
	sq := square{7}
	printArea(tri)
	printArea(sq)
}
