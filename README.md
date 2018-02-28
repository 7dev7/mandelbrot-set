Mandelbrot set Visualisation
============================
[Mandelbrot set](https://en.wikipedia.org/wiki/Mandelbrot_set) is a famous example of a fractal in mathematics.
Rendering the image is performed using goroutine per column.

### Run
```
$ go run main.go
```
<img src="https://raw.githubusercontent.com/7dev7/mandelbrot-set/master/images/1.png"/>

If you want to know all parameters:
```
$ go run main.go --help
```
Example:
```
$ go run main.go -re -1.96680095 -im 0.00000478 -r 0.00000014 -file "pic.png"
```
<img src="https://raw.githubusercontent.com/7dev7/mandelbrot-set/master/images/3.png"/>

```
$ go run main.go -re -1.7433419053321 -im 0.0000907687489 -r 0.00000000374 -file "pic.png" -w 4096 -h 2048
```
<img src="https://raw.githubusercontent.com/7dev7/mandelbrot-set/master/images/2.png"/>
