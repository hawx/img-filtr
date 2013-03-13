package recipes

import (
	"github.com/hawx/img/blur"
	"github.com/hawx/img/channel"
	"github.com/hawx/img/contrast"
	"github.com/hawx/img/greyscale"
	"github.com/hawx/img/utils"
	"github.com/hawx/img/sharpen"
	"image"
)

func Rockstr(in image.Image) image.Image {

	in = sharpen.UnsharpMask(in, 2, 1.5, 1.0, 0.05)
	// in = channel.Brightness(in, utils.Multiplier(1.75))
	// in = channel.Saturation(in, utils.Multiplier(1.5))
	// in = contrast.Adjust(in, 1.3) // guess

	in = utils.MapColor(in, utils.Compose(
		channel.BrightnessC(utils.Multiplier(1.75)),
		channel.SaturationC(utils.Multiplier(1.5)),
		contrast.AdjustC(3.0), // guess
	))

	in = blur.Gaussian(in, 1, 2, blur.WRAP)

	// in = greyscale.Photoshop(in)
	// in = contrast.Adjust(in, 1.1) // guess

	in = utils.MapColor(in, utils.Compose(
		greyscale.PhotoshopC(),
		contrast.AdjustC(1.0), // guess
	))

	in = sharpen.Sharpen(in, 5, 5)

	return in
}
