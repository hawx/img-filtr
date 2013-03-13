package recipes

import (
	"github.com/nfnt/resize"

	"github.com/hawx/img/contrast"
	"github.com/hawx/img/channel"
	"github.com/hawx/img/utils"
	"github.com/hawx/img/blend"
	"github.com/hawx/img/blur"
	"github.com/hawx/img/sharpen"

	"image"
)

func Postcrd(in image.Image) image.Image {
	b := in.Bounds()

	thumbWidth  := int(b.Dx() / 10)
	thumbHeight := int(b.Dy() / 10)

	// Create mask
	mask := resize.Resize(uint(thumbWidth), uint(thumbHeight), in, resize.Bilinear)
	mask  = contrast.Adjust(mask, 1.0) // guess amount
	mask  = channel.Saturation(mask, utils.Multiplier(1.5))
	mask  = blur.Gaussian(mask, 1, 2, blur.CLAMP)

	mask  = resize.Resize(uint(b.Dx()), uint(b.Dy()), mask, resize.Bilinear)
	mask  = blur.Gaussian(mask, 0, 5, blur.CLAMP)
	mask  = channel.Brightness(mask, utils.Multiplier(1.8))
	mask  = channel.Saturation(mask, utils.Multiplier(1.5))

	// Create lomo
	lomo := sharpen.UnsharpMask(in, 2, 1.5, 1.0, 0.05)
	lomo  = channel.Brightness(in, utils.Multiplier(1.75))
	lomo  = contrast.Adjust(lomo, 3.0) // again, guessing
	lomo  = blur.Gaussian(lomo, 1, 2, blur.CLAMP)

	// Compose
	final := blend.Multiply(mask, lomo)

	return final
}
