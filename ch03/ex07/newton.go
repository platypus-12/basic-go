package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	fillRect(img, color.RGBA{255, 255, 255, 255})

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width*(xmax); px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

/*
https://qiita.com/PlanetMeron/items/09d7eb204868e1a49f49
f'(z1) = Δy/Δx
       = f(z1) / z1 -z2
z1 - z2= f(z1) / f'(z1)
   z2  = z1 - f(z1) / f'(z1)
*/

func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 2
	for n := uint8(0); n < iterations; n++ {
		//z = z - (z*z*z*z - 1) / (4 * z * z*z)
		z = z*3/4 + 1/(4*z*z*z)
		if cmplx.Abs(z*z*z*z-1) < 0.01 {
			return color.RGBA{255 - contrast*n, 255 - contrast*2*n, 255 - contrast*3*n, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func fillRect(img *image.RGBA, col color.Color) {
	rect := img.Rect
	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			img.Set(v, h, col)
		}
	}
}
