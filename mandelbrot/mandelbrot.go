package mandelbrot

import (
	"math/cmplx"
	"image"
	"image/draw"
	"image/color"

	"github.com/nfnt/resize"
	"sync"
)

const (
	ReMin         = -2.5
	ReMax         = 1.5
	ImMin         = -1.
	ImMax         = 1.
	MaxIterations = 300
	Threshold     = 7
	Red           = 248
	Green         = 212
	Blue          = 160
)

func Create(width, height int, autoScale bool) image.Image {
	scale := float64(width / (ReMax - ReMin))
	autoHeight := int(scale * (ImMax - ImMin))
	b := generateImage(width, autoHeight, scale)
	if autoScale {
		return b
	}
	return resize.Resize(uint(width), uint(height), b, resize.Lanczos3)
}

func generateImage(width, height int, scale float64) image.Image {
	bounds := image.Rect(0, 0, width, height)
	b := image.NewRGBA(bounds)
	draw.Draw(b, bounds, image.NewUniform(color.Black), image.ZP, draw.Src)
	var wg sync.WaitGroup
	wg.Add(width)
	for x := 0; x < width; x++ {
		go func(x int) {
			defer wg.Done()
			for y := 0; y < height; y++ {
				approximation := calcApproximation(complex(float64(x)/scale+ReMin, float64(y)/scale+ImMin))
				b.Set(x, y, color.NRGBA{R: uint8(Red * approximation), G: uint8(Green * approximation), B: uint8(Blue * approximation), A: 255})
			}
		}(x)
	}
	wg.Wait()
	return b
}

func calcApproximation(a complex128) float64 {
	i := 0
	for z := a; cmplx.Abs(z) < Threshold && i < MaxIterations; i++ {
		z = z*z + a
	}
	return float64(MaxIterations-i) / MaxIterations
}
