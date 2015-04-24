package recipes

import (
	"image"
	"image/color"
	"image/draw"

	"hawx.me/code/img/crop"
	"hawx.me/code/img/utils"
)

const (
	TOP_BORDER    = 0.1
	SIDE_BORDER   = 0.1
	BOTTOM_BORDER = 0.3
)

func Edwn(in image.Image) image.Image {
	// Crop:
	in = crop.Square(in, -1, utils.Right)
	rect := in.Bounds()

	// Calculate border widths:
	topBorder := int(float64(rect.Dy()) * TOP_BORDER)
	sideBorder := int(float64(rect.Dx()) * SIDE_BORDER)
	bottomBorder := int(float64(rect.Dy()) * BOTTOM_BORDER)

	fullWidth := rect.Dx() + 2*sideBorder + 2
	fullHeight := rect.Dy() + topBorder + bottomBorder + 2

	// Bounds now start off (0, 0) since the image has been cropped, so need to
	// correct for this when calculating the image size and remove the margin
	imageSize := image.Rect(sideBorder+1-rect.Min.X, topBorder+1-rect.Min.Y, sideBorder+1+rect.Dx(), topBorder+1+rect.Dy())

	// Put white border round image
	whiteSize := image.Rect(1, 1, fullWidth-1, fullHeight-1)

	// Put 1px black border round
	blackSize := image.Rect(0, 0, fullWidth, fullHeight)

	white := image.NewUniform(color.White)
	black := image.NewUniform(color.Black)

	final := image.NewRGBA(blackSize)
	draw.Draw(final, blackSize, black, image.Pt(0, 0), draw.Src)
	draw.Draw(final, whiteSize, white, image.Pt(0, 0), draw.Src)
	draw.Draw(final, imageSize, in, image.Pt(0, 0), draw.Src)

	return final
}
