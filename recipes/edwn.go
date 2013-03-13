package recipes

import (
	"github.com/hawx/img/crop"
	"github.com/hawx/img/utils"
	"image"
	"image/color"
	"image/draw"
)

const (
	TOP_BORDER = 0.1
	SIDE_BORDER = 0.1
	BOTTOM_BORDER = 0.3
)

func Edwn(in image.Image) image.Image {
	// Crop:

	in = crop.Square(in, -1, utils.Centre)
	rect := in.Bounds()

	// Calculate border widths:

	topBorder    := int(float64(rect.Dy()) * TOP_BORDER)
	sideBorder   := int(float64(rect.Dx()) * SIDE_BORDER)
	bottomBorder := int(float64(rect.Dy()) * BOTTOM_BORDER)

	fullWidth  := rect.Dx() + 2 * sideBorder + 2
	fullHeight := rect.Dy() + topBorder + bottomBorder + 2

	// Put white border round image
	// Put 1px black border round

	imageSize := image.Rect(sideBorder+1, topBorder+1, sideBorder+1 + rect.Dx(), topBorder+1 + rect.Dy())
	whiteSize := image.Rect(1, 1, fullWidth - 1, fullHeight - 1)
	blackSize := image.Rect(0, 0, fullWidth, fullHeight)

	white     := image.NewUniform(color.White)
	black     := image.NewUniform(color.Black)

	final := image.NewRGBA(blackSize)

	draw.Draw(final, blackSize, black, image.Pt(0, 0), draw.Src)
	draw.Draw(final, whiteSize, white, image.Pt(0, 0), draw.Src)
  draw.Draw(final, imageSize, in, image.Pt(0, 0), draw.Src)

	return final
}
