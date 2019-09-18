package images

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"testing"
)

func Test_Image(t *testing.T) {
	// just print the image in base64 encoding
	// the base64 can then be converted in an image using https://base64.guru/converter/decode/file
	m := Image{}
	var buf bytes.Buffer
	png.Encode(&buf, m)
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	t.Log("PNG image (base64): " + enc)
}
