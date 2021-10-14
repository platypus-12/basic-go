package main

import (
	"fmt"
)

type Point struct{ X, Y int }
type address struct {
	hostname string
	port     int
}

func main() {
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++

	z := Point{1, 2}
	x := Point{2, 1}
	fmt.Println(z.X == x.X && z.Y == x.Y)
	fmt.Println(z == x)

	p := Point{1, 2}
	k := Scale(p, 2)
	j := Scale1(p, 2)

	fmt.Println(p, k, *j)
	fmt.Printf("%T %T %T\n", p, k, j)

	pp := &Point{1, 2}
	kk := new(Point)
	*kk = Point{1, 2}
	fmt.Println("a", pp, kk)

}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Scale1(p Point, factor int) *Point {
	return &Point{p.X * factor, p.Y * factor}
}
