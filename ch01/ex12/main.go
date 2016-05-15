// Ex12 serves GIF animations of random Lissajous curves in multiple colors,
// accepting query parameters for frequency, phase step and cycles, number of
// frames, delay between frames, image size and angular resolution.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black}

const paletteSize = 255

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fillPalette()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func fillPalette() {
	addToPalette := func(r, g, b int) {
		palette = append(palette, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
	}

	for i := 1; i < paletteSize; i += 3 {
		addToPalette(paletteSize-i, i, 0) // Red turning to Green
	}
	for i := 1; i < paletteSize; i += 3 {
		addToPalette(0, paletteSize-i, i) // Green turning to Blue
	}
	for i := 1; i < paletteSize; i += 3 {
		addToPalette(i, 0, paletteSize-i) // Blue turning to Red
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		return
	}
	freq, err := strconv.ParseFloat(r.Form.Get("freq"), 64)
	if err != nil {
		freq = rand.Float64() * 3.0
	}
	phaseStep, err := strconv.ParseFloat(r.Form.Get("phaseStep"), 64)
	if err != nil {
		phaseStep = 0.01
	}
	cycles, err := strconv.ParseFloat(r.Form.Get("cycles"), 64)
	if err != nil {
		cycles = 5
	}
	nframes, err := strconv.Atoi(r.Form.Get("nframes"))
	if err != nil {
		nframes = 64
	}
	delay, err := strconv.Atoi(r.Form.Get("delay"))
	if err != nil {
		delay = 8
	}
	size, err := strconv.Atoi(r.Form.Get("size"))
	if err != nil {
		size = 100
	}
	res, err := strconv.ParseFloat(r.Form.Get("res"), 64)
	if err != nil {
		res = 0.001
	}
	lissajous(w, freq, phaseStep, cycles, nframes, delay, size, res)
}

func lissajous(out io.Writer, freq float64, phaseStep float64, cycles float64,
	nframes int, delay int, size int, res float64) {
	fmt.Printf("freq=%f, phaseStep=%f, cycles=%f, nframes=%d, delay=%d, "+
		"size=%d, res=%f\n", freq, phaseStep, cycles, nframes, delay, size, res)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	for i := 0; i < nframes; i++ {
		img := image.NewPaletted(rect, palette)
		upperLimit := cycles * 2 * math.Pi
		for t := 0.0; t < upperLimit; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// Range from 1 to MaxUint8
			img.SetColorIndex(size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5),
				uint8(t*(math.MaxUint8-1)/upperLimit)+1)
		}
		phase += phaseStep
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
