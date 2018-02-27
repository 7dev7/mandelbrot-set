package mandelbrot

import (
	"math/cmplx"
	"image"
	"image/draw"
	"image/color"

	"github.com/nfnt/resize"
	"sync"
	"fmt"
)

type Config struct {
	Width     int
	Height    int
	RePos     float64
	ImPos     float64
	Radius    float64
	AutoScale bool
	FileName  string
}

func (cfg Config) String() string {
	return fmt.Sprintf("{Width: %d; Height: %d; RePos: %e; ImPos: %e; Radius: %e; Autoscale: %t; Filename: %s}",
		cfg.Width, cfg.Height, cfg.RePos, cfg.ImPos, cfg.Radius, cfg.AutoScale, cfg.FileName)
}

var cfg *Config

const (
	MaxIterations = 500
	Threshold     = 2
	Red           = 248
	Green         = 212
	Blue          = 160
)

func Create(config *Config) image.Image {
	cfg = config
	reMin, reMax := cfg.RePos-float64(cfg.Radius)/2., cfg.RePos+float64(cfg.Radius)/2.
	imMin, imMax := cfg.ImPos-float64(cfg.Radius)/2., cfg.ImPos+float64(cfg.Radius)/2.

	scale := float64(float64(cfg.Width) / (reMax - reMin))
	autoHeight := int(scale * (imMax - imMin))
	b := generateImage(cfg.Width, autoHeight, scale, reMin, imMin)
	if cfg.AutoScale {
		return b
	}
	return resize.Resize(uint(cfg.Width), uint(cfg.Height), b, resize.Lanczos3)
}

func generateImage(width, height int, scale, reMin, imMin float64) image.Image {
	bounds := image.Rect(0, 0, width, height)
	b := image.NewRGBA(bounds)
	draw.Draw(b, bounds, image.NewUniform(color.Black), image.ZP, draw.Src)
	var wg sync.WaitGroup
	wg.Add(width)
	for x := 0; x < width; x++ {
		go func(x int) {
			defer wg.Done()
			for y := 0; y < height; y++ {
				approximation := calcApproximation(complex(float64(x)/scale+reMin, float64(y)/scale+imMin))
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
