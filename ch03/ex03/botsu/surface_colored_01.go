package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = 30 * math.Pi / 180
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0.0; i < cells; i++ {
		for j := 0.0; j < cells; j++ {
			ax, ay, az, err := corner(float64(i+1), float64(j))
			if err != nil {
				continue
			}
			bx, by, bz, err := corner(float64(i), float64(j))
			if err != nil {
				continue
			}
			cx, cy, cz, err := corner(float64(i), float64(j+1))
			if err != nil {
				continue
			}
			dx, dy, dz, err := corner(float64(i+1), float64(j+1))
			if err != nil {
				continue
			}

			//	bは、(-1,0) (0,-1)と比較、dは、(0,1)(1,0)と比較、aは(1,0)(0,-1), cは(-1,0) (0,1)

		  _, _, a1, _ := corner(float64(i+1+0.1), float64(j))
			_, _, a2, _ := corner( float64(i+1), float64(j-0.1))
			_, _, b1, _ := corner(float64(i-0.1), float64(j))
			_, _, b2, _ := corner(float64(i), float64(j-0.1))
			_, _, c1, _ := corner(float64(i-0.1), float64(j+0.1))
			_, _, c2, _ := corner(float64(i), float64(j+1+0.1))
			_, _, d1, _ := corner(float64(i+1+0.1), float64(j+1))
			_, _, d2, _ := corner(float64(i+1), float64(j+1+0.1))

			color := "#000000"
			if az >= math.Max(a1, a2) ||  bz >= math.Max(b1, b2) || cz == math.Max(c1, c2) || dz == math.Max(d1, d2) {
				color = "#ff0000"
			}else if az <= math.Min(a1, a2) ||  bz <= math.Min(b1, b2) || cz <= math.Min(c1, c2) || dz <= math.Min(d1, d2) {
				color = "#0000ff"
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s;'/>", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

//sxとsyは、x, y, zを2次元に変換したもの
func corner(i, j float64) (float64, float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) || math.IsInf(z, -1) {
		return 0, 0, 0, fmt.Errorf("err")
	}
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
