package audiofingerprint

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

// Save the spectrogram image
func SaveSpectrogramImage(magSpec [][]float64, filename string) error {
	width := len(magSpec)
	height := len(magSpec[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Normalize range
	min, max := magSpec[0][0], magSpec[0][0]
	for _, row := range magSpec {
		for _, v := range row {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			val := (magSpec[x][y] - min) / (max - min)
			c := JetColorMap(val)
			img.Set(x, height-y-1, c)
		}
	}

	// Save to file
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()
	return png.Encode(outFile, img)
}

// JetColorMap maps a normalized value [0,1] to an RGB color using the Jet colormap
func JetColorMap(v float64) color.RGBA {
	v = math.Max(0, math.Min(1, v)) // clamp to [0,1]

	var r, g, b float64
	if v < 0.25 {
		r = 0
		g = 4 * v
		b = 1
	} else if v < 0.5 {
		r = 0
		g = 1
		b = 1 + 4*(0.25-v)
	} else if v < 0.75 {
		r = 4 * (v - 0.5)
		g = 1
		b = 0
	} else {
		r = 1
		g = 1 + 4*(0.75-v)
		b = 0
	}

	return color.RGBA{
		R: uint8(math.Max(0, math.Min(255, r*255))),
		G: uint8(math.Max(0, math.Min(255, g*255))),
		B: uint8(math.Max(0, math.Min(255, b*255))),
		A: 255,
	}
}
