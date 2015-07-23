package app

import (
	"bytes"
	"fmt"
	"image/color"
	"image/jpeg"
	"mime/multipart"

	"github.com/danhardman/image-to-table/utils"
)

//Table describes a table
type Table struct {
	Width  int
	Height int
	Rows   string
}

//ImageToTable takes an image file and returns a table struct
func ImageToTable(f multipart.File) *Table {
	ic, err := jpeg.DecodeConfig(f)
	utils.PanicOnError(err)

	t := &Table{
		Width:  ic.Width,
		Height: ic.Width,
	}

	_, err = f.Seek(0, 0)
	utils.PanicOnError(err)

	image, err := jpeg.Decode(f)
	utils.PanicOnError(err)

	var buffer bytes.Buffer
	for y := 0; y < ic.Height; y++ {
		buffer.WriteString("<tr>")
		for x := 0; x < ic.Width; x++ {
			color := image.At(x, y)
			hex := ColorToHex(color)

			buffer.WriteString("<td style=\"background-color:" + hex + "\"></td>")
		}
		buffer.WriteString("</tr>")
	}

	t.Rows = buffer.String()

	return t
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
