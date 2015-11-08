// Ex05 generates GIF animations of random Lissajous figures in multiple colors.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black}

const paletteSize = 255

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 1; i <= paletteSize; i++ {
		radians := float64(i) * 2 * math.Pi / paletteSize
		r := math.Sin(radians) * math.MaxUint8
		g := math.Sin(radians+2*math.Pi/3) * math.MaxUint8
		b := math.Sin(radians+4*math.Pi/3) * math.MaxUint8
		palette = append(palette, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
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
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(t*math.MaxUint8/(cycles*2*math.Pi)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
