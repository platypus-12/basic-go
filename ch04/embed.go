package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	v := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Println(w, v)
	fmt.Printf("%#v\n", w)
	fmt.Printf("%v\n", w)
	
	w.X = 43
	fmt.Printf("%#v\n", w)
}
