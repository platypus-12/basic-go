package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/cmplx"
	"os"
	"testing"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

type Coordinate struct {
	x int
	y int
	c color.Color
}

func main() {
	result := testing.Benchmark(func(b *testing.B) {
		b.ResetTimer()
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		coordinates := make(chan Coordinate)
		for py := 0; py < height; py++ {
			go concurrent(py, coordinates, img)
		}

		// 質問rangeだとうまくいかない理由がわからない
		// for x := range coordinates {
		// 	img.Set(x.x, x.y, x.c)
		// }

		for i := 0; i < height*width*(xmax); i++ {
			c := <-coordinates
			img.Set(c.x, c.y, c.c)
		}

		png.Encode(os.Stdout, img)
	})
	fmt.Printf("%s\n", result)
}

func concurrent(py int, coordinates chan Coordinate, img draw.Image) {
	y := float64(py)/height*(ymax-ymin) + ymin
	for px := 0; px < width*(xmax); px++ {
		x := float64(px)/width*(xmax-xmin) + xmin
		z := complex(x, y)
		coordinates <- Coordinate{px, py, mandelbrot(z)}
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
