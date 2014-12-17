package recipes

import (
	"github.com/nfnt/resize"

	"github.com/hawx/img/blend"
	"github.com/hawx/img/blur"
	"github.com/hawx/img/channel"
	"github.com/hawx/img/contrast"
	"github.com/hawx/img/sharpen"
	"github.com/hawx/img/utils"

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
