package recipes

import (
	"image"

	"hawx.me/code/img/blend"
)

func Postr(in image.Image) image.Image {
	return blend.Multiply(in, Postcrd(in))
}
