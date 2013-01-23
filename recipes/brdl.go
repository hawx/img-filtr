package recipes

import (
	"github.com/hawx/img/utils"
	"image"
	"image/color"
	"image/draw"
	"math/rand"
)

const FACTOR = 10

func Brdl(in image.Image) image.Image {
	table := map[color.Color] int {}
	img   := in.(draw.Image)

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

	return img
}