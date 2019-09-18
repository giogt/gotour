package slices

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"testing"
)

func Test_pic(t *testing.T) {
	// just print the image in base64 encoding
	// the base64 can then be converted in an image using https://base64.guru/converter/decode/file
	show(t, Pic)
}

func show(t *testing.T, f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{dx, dy},
	})
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	showImage(t, m)
}

func showImage(t *testing.T, m image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, m)
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	t.Log("PNG image (base64): " + enc)
}
