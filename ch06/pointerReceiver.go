package main

import "fmt"

type Point struct{ X, Y float64 }

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p)

	q := Point{1, 2}
	(&q).ScaleBy(2)
	fmt.Println(q)

	d := Point{1, 2}
	d.ScaleBy(2)
	fmt.Println(d)
}
