package recipes

import (
	"github.com/nfnt/resize"

	"hawx.me/code/img/blend"
	"hawx.me/code/img/blur"
	"hawx.me/code/img/channel"
	"hawx.me/code/img/contrast"
	"hawx.me/code/img/sharpen"
	"hawx.me/code/img/utils"

	"image"
)

func Postcrd(in image.Image) image.Image {
	b := in.Bounds()

	thumbWidth := int(b.Dx() / 10)
	thumbHeight := int(b.Dy() / 10)

	// Create mask
	mask := resize.Resize(uint(thumbWidth), uint(thumbHeight), in, resize.Bilinear)
	mask = contrast.Adjust(mask, 1.0)
	mask = channel.Adjust(mask, utils.Multiplier(1.5), channel.Saturation)
	mask = blur.Gaussian(mask, 1, 2, blur.CLAMP)

	mask = resize.Resize(uint(b.Dx()), uint(b.Dy()), mask, resize.Bilinear)
	mask = blur.Gaussian(mask, 0, 5, blur.CLAMP)
	mask = channel.Adjust(mask, utils.Multiplier(1.8), channel.Lightness)
	mask = channel.Adjust(mask, utils.Multiplier(1.5), channel.Saturation)

	// Create lomo
	lomo := sharpen.UnsharpMask(in, 2, 1.5, 1.0, 0.05)
	lomo = channel.Adjust(in, utils.Multiplier(1.75), channel.Lightness)
	lomo = contrast.Adjust(lomo, 3.0)
	lomo = blur.Gaussian(lomo, 1, 2, blur.CLAMP)

	// Compose
	final := blend.Multiply(mask, lomo)

	return final
}
