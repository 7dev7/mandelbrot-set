package main

import (
	"github.com/7dev7/mandelbrot-set/mandelbrot"
	"os"
	"log"
	"image/png"
	"flag"
)

var (
	cfg *mandelbrot.Config
)

func init() {
	cfg = &mandelbrot.Config{}
	flag.IntVar(&cfg.Width, "w", 1920, "width of generated image")
	flag.IntVar(&cfg.Height, "h", 1080, "height of generated image")
	flag.Float64Var(&cfg.RePos, "re", -1., "real value")
	flag.Float64Var(&cfg.ImPos, "im", 0., "imaginary value")
	flag.Float64Var(&cfg.Radius, "r", 2., "radius")
	flag.BoolVar(&cfg.AutoScale, "autoscale", true, "Use auto scale for image")
	flag.StringVar(&cfg.FileName, "file", "example.png", "Output filename of generated image")
	flag.Parse()
}

func main() {
	log.Printf("started generation Mandelbrot set with cfg: %v\n", cfg)
	img := mandelbrot.Create(cfg)
	if file, e := os.Create(cfg.FileName); e != nil {
		log.Fatalln(e)
	} else {
		png.Encode(file, img)
		log.Println("done")
	}
}
