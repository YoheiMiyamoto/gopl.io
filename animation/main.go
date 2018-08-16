// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math/rand"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	handler := func(w http.ResponseWriter, r *http.Request) {
		generate(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func generate(out io.Writer) {
	const nframes = 64
	anim := gif.GIF{LoopCount: nframes}
	anim = setNewFrames(anim, nframes)
	star(anim.Image, 20, 20, 10)
	star(anim.Image, 10, 10, 20)
	gif.EncodeAll(out, &anim)
}

func setNewFrames(anim gif.GIF, num int) gif.GIF {
	const (
		delay = 8
		size  = 100
	)
	for i := 0; i < num; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	return anim
}

func star(imgs []*image.Paletted, x, y, interval int) {
	bool := true
	count := 0
	for _, img := range imgs {
		if bool {
			img.SetColorIndex(x, y, blackIndex)
		}
		count++
		if count == interval {
			count = 0
			bool = !bool
		}
	}
}

func dot(img *image.Paletted, x, y, size int) {
	img.SetColorIndex(x, y, blackIndex)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	// freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	// phase := 0.0 // phase difference
	// for i := 0; i < nframes; i++ {
	// 	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	// 	img := image.NewPaletted(rect, palette)
	// 	for t := 0.0; t < cycles*2*math.Pi; t += res {
	// 		x := math.Sin(t)
	// 		y := math.Sin(t*freq + phase)
	// 		img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
	// 			blackIndex)
	// 	}
	// 	phase += 0.1
	// 	anim.Delay = append(anim.Delay, delay)
	// 	anim.Image = append(anim.Image, img)
	// }

	// for i := 0; i < 10; i++ {
	// 	for i2 := 0; i2 < 100; i2++ {
	// 		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	// 		img := image.NewPaletted(rect, palette)
	// 		img.SetColorIndex(i2, i2, blackIndex)
	// 		img.SetColorIndex(i2+1, i2+1, blackIndex)
	// 		img.SetColorIndex(i2, i2+1, blackIndex)
	// 		img.SetColorIndex(i2+1, i2, blackIndex)
	// 		anim.Delay = append(anim.Delay, delay)
	// 		anim.Image = append(anim.Image, img)
	// 	}
	// }

	anim = move(100, 8, anim)

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func move(rectSize, delay int, anim gif.GIF) gif.GIF {
	for i := 0; i < 100; i++ {
		rect := image.Rect(0, 0, 2*rectSize+1, 2*rectSize+1)
		img := image.NewPaletted(rect, palette)

		img.SetColorIndex(i, i, blackIndex)
		img.SetColorIndex(i+1, i+1, blackIndex)
		img.SetColorIndex(i, i+1, blackIndex)
		img.SetColorIndex(i+1, i, blackIndex)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	return anim
}

// func dot(img *image.Paletted, x, y int) *image.Paletted {
// 	img.SetColorIndex(x, y, blackIndex)
// 	img.SetColorIndex(x+1, y+1, blackIndex)
// 	img.SetColorIndex(x, y+1, blackIndex)
// 	img.SetColorIndex(x+y, i, blackIndex)
// }

//!-main
