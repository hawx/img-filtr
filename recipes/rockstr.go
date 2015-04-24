package recipes

import (
	"image"

	"hawx.me/code/img/blur"
	"hawx.me/code/img/channel"
	"hawx.me/code/img/contrast"
	"hawx.me/code/img/gamma"
	"hawx.me/code/img/greyscale"
	"hawx.me/code/img/sharpen"
	"hawx.me/code/img/utils"
)

func Rockstr(in image.Image) image.Image {
	in = sharpen.UnsharpMask(in, 2, 1.5, 1.0, 0.05)

	in = utils.MapColor(in, utils.Compose(
		channel.AdjustC(utils.Multiplier(1.75), channel.Brightness),
		channel.AdjustC(utils.Multiplier(1.5), channel.Saturation),
		contrast.AdjustC(3.0),
	))

	in = blur.Gaussian(in, 1, 2, blur.WRAP)

	in = utils.MapColor(in, utils.Compose(
		gamma.AdjustC(1.0/1.8),
		greyscale.GreyscaleC(),
		contrast.AdjustC(1.0),
	))

	in = sharpen.UnsharpMask(in, 5, 25, 1.0, 0.05)

	return in
}
