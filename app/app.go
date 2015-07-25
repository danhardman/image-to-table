package app

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime"
	"mime/multipart"
	"strings"
)

//Table describes a table
type Table struct {
	Width  int
	Height int
	Rows   string
}

//Image is a wrapper for an image and image config
type Image struct {
	Config image.Config
	Image  image.Image
}

func getFileType(fh *multipart.FileHeader) string {
	s := strings.Split(fh.Filename, ".")
	ext := "." + s[len(s)-1]
	mimetype := mime.TypeByExtension(ext)

	if mimetype == fh.Header["Content-Type"][0] {
		return mimetype
	}

	return "Unknown"
}

func getImage(f multipart.File, ft string) (*Image, error) {
	switch ft {
	case "image/jpeg":
		return decodeJPEG(f)
	case "image/gif":
		return decodeGIF(f)
	case "image/png":
		return decodePNG(f)
	}

	return nil, fmt.Errorf("Invalid file type.")
}

func decodeJPEG(f multipart.File) (*Image, error) {
	var image Image

	ic, err := jpeg.DecodeConfig(f)
	_, err = f.Seek(0, 0)
	i, err := jpeg.Decode(f)

	if err != nil {
		return nil, fmt.Errorf("Could not decode JPEG file.")
	}

	image.Config = ic
	image.Image = i

	return &image, nil
}

func decodeGIF(f multipart.File) (*Image, error) {
	var image Image

	ic, err := gif.DecodeConfig(f)
	_, err = f.Seek(0, 0)
	i, err := gif.Decode(f)

	if err != nil {
		return nil, fmt.Errorf("Could not decode GIF file.")
	}

	image.Config = ic
	image.Image = i

	return &image, nil
}

func decodePNG(f multipart.File) (*Image, error) {
	var image Image

	ic, err := png.DecodeConfig(f)
	_, err = f.Seek(0, 0)
	i, err := png.Decode(f)

	if err != nil {
		return nil, fmt.Errorf("Could not decode PNG file.")
	}

	image.Config = ic
	image.Image = i

	return &image, nil
}

//ImageToTable takes an image file and returns a table struct
func ImageToTable(f multipart.File, fh *multipart.FileHeader) (*Table, error) {
	ft := getFileType(fh)
	image, err := getImage(f, ft)
	if err != nil {
		return nil, err
	}

	t := &Table{
		Width:  image.Config.Width,
		Height: image.Config.Height,
	}

	var buffer bytes.Buffer
	for y := 0; y < t.Height; y++ {
		buffer.WriteString("<tr>")
		for x := 0; x < t.Width; x++ {
			color := image.Image.At(x, y)
			hex := ColorToHex(color)

			buffer.WriteString("<td style=\"background-color:" + hex + "\"></td>")
		}
		buffer.WriteString("</tr>")
	}

	t.Rows = buffer.String()

	return t, nil
}

//RGBToHex converts RGB values to a hex string
func rgbToHex(r, g, b uint8) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

//ColorToHex takes a color struct and returns the hexidecimal representation
//of its RGB color value
func ColorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return rgbToHex(uint8(r>>8), uint8(g>>8), uint8(b>>8))
}
