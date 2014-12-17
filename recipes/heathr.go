package recipes

import (
	"github.com/hawx/img/crop"
	"github.com/hawx/img/utils"
	"github.com/nfnt/resize"

	"image"
	"image/color"
	"image/draw"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Heathr(left, right image.Image) image.Image {
	// assume left and right have been filtrd before being passed here
	// this just needs to compose them onto a white background

	left = crop.Square(left, -1, utils.Centre)
	right = crop.Square(right, -1, utils.Centre)

	leftSize := left.Bounds().Dx()
	rightSize := right.Bounds().Dx()
	central := int(float64(leftSize) * 0.01)

	if leftSize > rightSize {
		left = resize.Resize(uint(rightSize), uint(rightSize), left, resize.Bilinear)

	} else if leftSize < rightSize {
		right = resize.Resize(uint(leftSize), uint(leftSize), right, resize.Bilinear)
	}

	left = Edwn(left)
	right = Edwn(right)

	finalBounds := image.Rect(0, 0, left.Bounds().Dx()*2+central, left.Bounds().Dy())
	leftBounds := image.Rect(0, 0, left.Bounds().Dx(), left.Bounds().Dy())
	rightBounds := image.Rect(left.Bounds().Dx()+central, 0, left.Bounds().Dx()*2+central, left.Bounds().Dy())

	final := image.NewRGBA(finalBounds)
	white := image.NewUniform(color.White)

	draw.Draw(final, finalBounds, white, image.Pt(0, 0), draw.Src)
	draw.Draw(final, leftBounds, left, image.Pt(0, 0), draw.Src)
	draw.Draw(final, rightBounds, right, image.Pt(0, 0), draw.Src)

	return final
}
