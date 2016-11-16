// Ex05 generates GIF animations of random Lissajous curves in multiple colors.
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
	"os"
	"time"
)

var palette = []color.Color{color.Black}

const paletteSize = 256

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	addToPalette := func(r, g, b int) {
		palette = append(palette, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
	}

	// 255/3 is exactly 85, so each loop adds 85 entries
	for i := 1; i < paletteSize; i += 3 {
		addToPalette(paletteSize-i, i, 0) // Red turning to Green
	}
	for i := 1; i < paletteSize; i += 3 {
		addToPalette(0, paletteSize-i, i) // Green turning to Blue
	}
	for i := 1; i < paletteSize; i += 3 {
		addToPalette(i, 0, paletteSize-i) // Blue turning to Red
	}

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	for i := 0; i < nframes; i++ {
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// Range from 1 to paletteSize-1 (use full palette exactly once)
			paletteIndex := uint8(t*(paletteSize-1)/(cycles*2*math.Pi)) + 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				paletteIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
