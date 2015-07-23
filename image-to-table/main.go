package main

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"strconv"
)

// func init() {
// 	// damn important or else At(), Bounds() functions will
// 	// caused memory pointer error!!
// 	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
// }

func main() {
	abs, err := filepath.Abs("/Users/danhardman/Documents/images/logo.jpg")
	if err != nil {
		panic(err)
	}

	f, err := os.Open(abs)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	ic, err := jpeg.DecodeConfig(f)
	if err != nil {
		panic(err)
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	image, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	abs, err = filepath.Abs("/Users/danhardman/Documents/output/test.html")
	if err != nil {
		panic(err)
	}

	output, err := os.Create(abs)
	if err != nil {
		panic(err)
	}

	widthString := strconv.Itoa(ic.Width)

	heightString := strconv.Itoa(ic.Height)

	defer output.Close()
	hs := []byte("<html><body><table width=\"" + widthString + "\" height=\"" + heightString + "\" style=\"border-spacing:0;border-collapse:collapse;\">")
	_, err = output.Write(hs)
	if err != nil {
		panic(err)
	}

	for y := 0; y < ic.Height; y++ {
		rs := []byte("<tr style=\"height:1px;padding:0;\">")
		_, err = output.Write(rs)
		if err != nil {
			panic(err)
		}

		for x := 0; x < ic.Width; x++ {
			color := image.At(x, y)
			hex := ColorToHex(color)

			cell := []byte("<td style=\"width:1px;height:1px;padding:0;background-color:" + hex + "\"></td>")
			_, err = output.Write(cell)
			if err != nil {
				panic(err)
			}
		}

		re := []byte("</tr>")
		_, err = output.Write(re)
		if err != nil {
			panic(err)
		}
	}

	he := []byte("</table></body></html>")
	_, err = output.Write(he)
	if err != nil {
		panic(err)
	}
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
