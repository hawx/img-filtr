package recipes

import (
	"github.com/hawx/img/blend"
	"image"
)

func Postr(in image.Image) image.Image {
	return blend.Multiply(in, Postcrd(in))
}
