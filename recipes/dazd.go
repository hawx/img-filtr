package recipes

import (
	"image"

	"hawx.me/code/img/blur"
	"hawx.me/code/img/channel"
	"hawx.me/code/img/contrast"
	"hawx.me/code/img/sharpen"
	"hawx.me/code/img/utils"
)

func Dazd(img image.Image) image.Image {
	// Unsharp with radius and sigma 1.5, others are default IM values.
	// Radius is rounded up, as it is by IM internally.
	img = sharpen.UnsharpMask(img, 2, 1.5, 1.0, 0.05)

	img = utils.MapColor(img, utils.Compose(
		channel.AdjustC(utils.Multiplier(1.75), channel.Lightness),
		channel.AdjustC(utils.Multiplier(1.5), channel.Saturation),
		contrast.AdjustC(3.0),
	))

	// gaussian with radius=1, sigma=2
	img = blur.Gaussian(img, 1, 2, blur.CLAMP)

	return img
}
