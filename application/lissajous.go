package application

import (
	"image/color"
	"math/rand"
	"time"
	"image/gif"
	"image"
	"math"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajou() {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		max := 2*size + 1
		rect := image.Rect(0, 0, max, max)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := size + int(math.Sin(t)*size+0.5)
			y := size + int(math.Sin(t*freq+phase)*size+0.5)
			img.SetColorIndex(x, y, blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(os.Stdout, &anim)
}
