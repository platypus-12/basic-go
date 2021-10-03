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

	super_sampling_delta_x := 0.5 * (xmax - xmin) / height
	super_sampling_delta_y := 0.5 * (ymax - ymin) / width
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width*(xmax); px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			z1 := complex(x+super_sampling_delta_x, y)
			z2 := complex(x, y+super_sampling_delta_y)
			z3 := complex(x+super_sampling_delta_x, y+super_sampling_delta_y)

			samplings := []color.Color{
				mandelbrot(z),
				mandelbrot(z1),
				mandelbrot(z2),
				mandelbrot(z3),
			}
			var sum_r, sum_g, sum_b, sum_a uint32
			for _, sampling := range samplings {
				r, g, b, a := sampling.RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8
				sum_r += r
				sum_g += g
				sum_b += b
				sum_a += a
			}
			img.Set(px, py, color.RGBA{uint8(sum_r / 4), uint8(sum_g / 4), uint8(sum_b / 4), uint8(sum_a / 4)})
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
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
