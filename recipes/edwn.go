package recipes

import (
	"image"
	"image/color"
	"image/draw"
)

func crop(img image.Image, rect image.Rectangle) image.Image {
	o := image.NewRGBA(rect)

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			o.Set(x, y, img.At(x, y))
		}
	}

	return o
}

const (
	TOP_BORDER = 0.1
	SIDE_BORDER = 0.1
	BOTTOM_BORDER = 0.3
)

func Edwn(in image.Image) image.Image {
	b := in.Bounds()

	// Crop:

	rect := image.Rect(b.Min.X, b.Min.X, b.Max.X, b.Max.X)
	if b.Dx() > b.Dy() {
		rect = image.Rect(b.Min.Y, b.Min.Y, b.Max.Y, b.Max.Y)
	}

	in = crop(in, rect)

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
