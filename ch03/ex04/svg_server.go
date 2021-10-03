package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = 30 * math.Pi / 180
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var width int
var height int
var xyscale int
var zscale float64

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	height_r := r.URL.Query().Get("height")
	width_r := r.URL.Query().Get("width")
	color := r.URL.Query().Get("color")
	fmt.Println(height_r, width_r, color)

	if height_r == "" || width_r == "" || color == "" {
		fmt.Fprintf(w, "plz attach query string whose field is cycles: e.g.http://localhost:8000/?height=200&width=300&color=ff0000")
		return
	}
	var err1 error
	var err2 error
	height, err1 = strconv.Atoi(height_r)
	width, err2 = strconv.Atoi(width_r)
	if err1 != nil || err2 != nil || len(color) != 6 {
		fmt.Fprintf(w, "plz attach query string whose field is cycles: e.g.http://localhost:8000/?height=200&width=300&color=ff0000")
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, color)
}
func surface(out io.Writer, color string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	xyscale = width / 2 / xyrange
	zscale = float64(height) * 0.4
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:#%s;'/>", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) || math.IsInf(z, -1) {
		return 0, 0, fmt.Errorf("err")
	}
	sx := float64(width/2) + (x-y)*cos30*float64(xyscale)
	sy := float64(height/2) + (x+y)*sin30*float64(xyscale) - z*float64(zscale)
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
