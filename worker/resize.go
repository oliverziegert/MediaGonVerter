package worker

import (
	"golang.org/x/image/draw"
	"image"
	"math"
)

type ResizeConfig struct {
	Width           int
	Height          int
	Crop            bool
	BackgroundColor *image.Uniform
}

func Resize(imgSrc image.Image, rc ResizeConfig) *image.RGBA {
	var size *image.Rectangle
	var rectDst *image.Rectangle
	var leftTop *image.Point

	if rc.Crop {
		size = &image.Rectangle{Min: image.Point{}, Max: image.Point{X: rc.Width, Y: rc.Height}}
		imgSize := rc.computeSize(imgSrc.Bounds())
		leftTop = rc.computedCropLeftTop(imgSize)
		rectDst = &image.Rectangle{Min: *leftTop, Max: leftTop.Add(imgSize.Max)}
	} else {
		size = rc.computeSize(imgSrc.Bounds())
		leftTop = &image.Point{}
		rectDst = size
	}

	imgDst := image.NewRGBA(*size)

	draw.Draw(imgDst, imgDst.Bounds(), rc.BackgroundColor, image.Point{}, draw.Src)

	draw.NearestNeighbor.Scale(imgDst, *rectDst, imgSrc, imgSrc.Bounds(), draw.Over, nil)
	//draw.NearestNeighbor
	//draw.ApproxBiLinear
	//draw.BiLinear
	//draw.CatmullRom

	return imgDst
}

func (rc *ResizeConfig) computeSize(bounds image.Rectangle) *image.Rectangle {
	ratioW := float64(rc.Width) / float64(bounds.Max.X)
	ratioH := float64(rc.Height) / float64(bounds.Max.Y)

	ratio := math.Min(ratioW, ratioH)

	width := float64(bounds.Max.X) * ratio
	height := float64(bounds.Max.Y) * ratio

	rect := image.Rect(0, 0, int(width), int(height))

	return &rect
}

func (rc *ResizeConfig) computedCropLeftTop(bounds *image.Rectangle) *image.Point {
	if rc.Height == bounds.Max.Y && rc.Width == bounds.Max.X {
		return &image.Point{}
	}

	halfImgDstWidth := float64(rc.Width) / 2
	halfImgDstHeight := float64(rc.Height) / 2

	halfImgSrcWidth := float64(bounds.Max.X) / 2
	halfImgSrcHeight := float64(bounds.Max.Y) / 2

	return &image.Point{X: int(halfImgDstWidth - halfImgSrcWidth), Y: int(halfImgDstHeight - halfImgSrcHeight)}
}
