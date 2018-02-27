package main

import (
	"github.com/7dev7/mandelbrot-set/mandelbrot"
	"os"
	"log"
	"image/png"
	"flag"
)

var (
	width     int
	height    int
	autoScale bool
	fileName  string
)

func init() {
	flag.IntVar(&width, "width", 1920, "Width of generated image")
	flag.IntVar(&height, "height", 1080, "Height of generated image")
	flag.BoolVar(&autoScale, "autoscale", true, "Use auto scale for image")
	flag.StringVar(&fileName, "file", "example.png", "Output filename of generated image")
	flag.Parse()
}

func main() {
	log.Printf("started generation Mandelbrot set with params: width=%d, height=%d, autoscale=%t, fileName=%s\n", width, height, autoScale, fileName)
	img := mandelbrot.Create(width, height, autoScale)
	if file, e := os.Create(fileName); e != nil {
		log.Fatalln(e)
	} else {
		png.Encode(file, img)
		log.Println("done")
	}
}
