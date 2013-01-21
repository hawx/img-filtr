package main

import (
	"github.com/hawx/img/utils"
	"fmt"
	"flag"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

const FACTOR = 10

func main() {
	var long  = flag.Bool("long", false, "")
	var short = flag.Bool("short", false, "")
	var usage = flag.Bool("usage", false, "")

	flag.Parse()

	if *long {
		fmt.Println("  Repaints the image with a random dominant colour.")

	} else if *short {
		fmt.Println("the smallest possible transmittable unit")

	} else if *usage {
		fmt.Println("brdl [options]")

	} else {
		img := utils.ReadStdin().(draw.Image)

		table := map[color.Color] int {}

		utils.EachColor(img, func(c color.Color) {
			if v, ok := table[c]; ok {
				table[c] = v + 1
			} else {
				table[c] = 1
			}
		})

		r := rand.Intn(len(table) / FACTOR)
		var c color.Color

		for k,_ := range table {
			r -= 1
			if r <= 0 {
				c = k
			}
		}

		brdl := image.NewUniform(c)
		draw.Draw(img, img.Bounds(), brdl, image.Point{0,0}, draw.Src)

		utils.WriteStdout(img)
	}
}
