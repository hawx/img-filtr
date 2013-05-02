package recipes

import (
	"github.com/hawx/img/utils"
	"github.com/hawx/img/greyscale"
	"image"
	"image/color"
	"image/draw"
)

func Dthr(in image.Image) image.Image {
	img := greyscale.Greyscale(in).(draw.Image)
	b := img.Bounds()

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			pixel := img.At(x, y)
			oldGrey, _, _, a := utils.NormalisedRGBA(pixel)
			var newGrey uint32 = 0

			if oldGrey > 127 {
				newGrey = 255
			}

			err := (int(oldGrey) - int(newGrey)) / 8

			img.Set(x, y, color.NRGBA{uint8(newGrey), uint8(newGrey), uint8(newGrey), uint8(a)})

			points := []image.Point{
				image.Point{x+1, y  },
				image.Point{x+2, y  },
				image.Point{x-1, y+1},
				image.Point{x,   y+1},
				image.Point{x+1, y+1},
				image.Point{x,   y+2},
			}

			for _, point := range points {
				if point.X >= b.Min.X && point.X < b.Max.X && point.Y >= b.Min.Y && point.Y < b.Max.Y {
					g,_,_,a := utils.NormalisedRGBA(img.At(point.X, point.Y))
					ng := uint8(utils.Truncate(uint32(int(g) + err)))
					newColor := color.NRGBA{ng, ng, ng, uint8(a)}
					img.Set(point.X, point.Y, newColor)
				}
			}
		}
	}

	return img
}
