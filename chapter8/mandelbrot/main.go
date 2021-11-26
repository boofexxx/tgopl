package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	var wg sync.WaitGroup

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		wg.Add(1)
		go func(py int) {
			defer wg.Done()
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z
				img.Set(px, py, mandelbrot(z))
			}
		}(py)
	}
	wg.Wait()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	fmt.Fprintln(os.Stderr, time.Now().Sub(start))
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, 0xFF, 0xFF, 0xFF}
		}
	}
	return color.RGBA{0x00, 0x00, 0x00, 0xFF}
}
