package utils

import (
	"strconv"
)

var (
	Black, White *Color
)

func init() {
	initColorCube()
	Black, _ = NewColor("#000000")
	White, _ = NewColor("#ffffff")
}

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func NewColor(hex string) (*Color, error) {
	red, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return nil, err
	}
	green, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return nil, err
	}
	blue, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return nil, err
	}

	return &Color{
		Red:   uint8(red),
		Green: uint8(green),
		Blue:  uint8(blue),
	}, nil
}

var x6colorIndexes = [6]uint8{0, 95, 135, 175, 215, 255}
var x6colorCube [216]Color

func initColorCube() {
	i := 0
	for iR := 0; iR < 6; iR++ {
		for iG := 0; iG < 6; iG++ {
			for iB := 0; iB < 6; iB++ {
				x6colorCube[i] = Color{
					x6colorIndexes[iR],
					x6colorIndexes[iG],
					x6colorIndexes[iB],
				}
				i++
			}
		}
	}
}
