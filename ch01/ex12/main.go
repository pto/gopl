// Ex12 serves GIF animations of random Lissajous figures, accepting query
// parameters for cycles, image size, number of frames, and delay between
// frames.
package main

import (
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

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}
	cycles, err := strconv.Atoi(r.Form.Get("cycles"))
	if err != nil {
		cycles = 5
	}
	size, err := strconv.Atoi(r.Form.Get("size"))
	if err != nil {
		size = 100
	}
	nframes, err := strconv.Atoi(r.Form.Get("nframes"))
	if err != nil {
		nframes = 64
	}
	delay, err := strconv.Atoi(r.Form.Get("delay"))
	if err != nil {
		delay = 8
	}
	lissajous(w, cycles, size, nframes, delay)
}

func lissajous(out io.Writer, cycles int, size int, nframes int, delay int) {
	const res = 0.001            // angular resolution
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
