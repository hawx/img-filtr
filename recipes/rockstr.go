package recipes

import (
	"github.com/hawx/img/blur"
	"github.com/hawx/img/channel"
	"github.com/hawx/img/contrast"
	"github.com/hawx/img/gamma"
	"github.com/hawx/img/greyscale"
	"github.com/hawx/img/utils"
	"github.com/hawx/img/sharpen"
	"image"
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
		gamma.AdjustC(1.0 / 1.8),
		greyscale.DefaultC(),
		contrast.AdjustC(1.0),
	))

	in = sharpen.UnsharpMask(in, 5, 25, 1.0, 0.05)

	return in
}
