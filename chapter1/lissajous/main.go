package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.RGBA{0x00, 0x80, 0x00, 0xFF}, color.Black, color.White}

const (
	whiteIndex = 0
	blackIndex = 1
	thirdIndex = 2
)

func Lissajous(out io.Writer, res float64, cycles, size, nframes, delay int) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.00
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if i%2 == 0 {
				img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), thirdIndex)
			} else {
				img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
