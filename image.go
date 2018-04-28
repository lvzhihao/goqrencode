package goqrencode

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

const (
	IMAGE_OUTPUT_JPEG string = "jpeg"
	IMAGE_OUTPUT_JPG  string = "jpg"
	IMAGE_OUTPUT_PNG  string = "png"
)

var (
	ImageDefaultSize            int           = 200
	ImageDefaultBackgroundColor color.Color   = color.White
	ImageDefaultForegroundColor color.Color   = color.Black
	DefaultJPEGOptions          *jpeg.Options = &jpeg.Options{
		//Quality: 92, // gives a very high-quality image while gaining a significant reduction on the original 100% file size.
		Quality: 85, // gives a greater file size reduction with almost no loss in quality.
		//Quality: 75, // and lower begins to show obvious differences in the image, which can reduce your website user experience.
	}
)

type Image struct {
	size            int         // pixel for image
	backgroundColor color.Color // background color
	foregroundColor color.Color // foreground color
	image           image.Image // golang image.Image
}

func ImageNew() (ret *Image) {
	ret = new(Image)
	ret.New()
	return
}

func (c *Image) New() {
	c.size = ImageDefaultSize
	c.backgroundColor = ImageDefaultBackgroundColor
	c.foregroundColor = ImageDefaultBackgroundColor
	c.image = nil
}

func (c *Image) SetSize(size int) *Image {
	c.size = size
	return c
}

func (c *Image) GetSize() int {
	return c.size
}

func (c *Image) SetBackgroundColor(back color.Color) *Image {
	c.backgroundColor = back
	return c
}

func (c *Image) GetBackgroundColor() color.Color {
	return c.backgroundColor
}

func (c *Image) SetForegroundColor(fore color.Color) *Image {
	c.foregroundColor = fore
	return c
}

func (c *Image) GetForegroundColor() color.Color {
	return c.foregroundColor
}

func (c *Image) DrawQRcode(qrcode *QRcode) *Image {
	c.Draw(qrcode.Bitmap())
	return c
}

func (c *Image) Draw(bitmap [][]bool) *Image {
	width := len(bitmap)
	// adjust size
	if c.size < width {
		c.size = width
	}
	// per bit pixels
	pixels := c.size / width
	// pedding offset, used real size
	offset := (c.size - width*pixels) / 2
	// draw size
	rectangle := image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{c.size, c.size},
	}
	// create color palette
	palette := color.Palette([]color.Color{
		c.backgroundColor,
		c.foregroundColor,
	})
	// create palette image
	paletted := image.NewPaletted(rectangle, palette)
	// fix background color
	for i := 0; i < c.size; i++ {
		for j := 0; j < c.size; j++ {
			paletted.Set(i, j, c.backgroundColor)
		}
	}
	// saves qr bitmaps
	for y, row := range bitmap {
		for x, v := range row {
			if !v {
				continue
			}
			// per bit pixels loc
			sx := x*pixels + offset
			sy := y*pixels + offset
			// fix foreground color
			for i := sx; i < sx+pixels; i++ {
				for j := sy; j < sy+pixels; j++ {
					paletted.Set(i, j, c.foregroundColor)
				}
			}
		}
	}
	// image.Image interface
	c.image = paletted.SubImage(rectangle)
	return c
}

func (c *Image) Encode(outputType string, w io.Writer) error {
	switch outputType {
	case IMAGE_OUTPUT_JPEG:
		return c.JPEG(w, nil)
	case IMAGE_OUTPUT_JPG:
		return c.JPEG(w, nil)
	default:
		return c.PNG(w)
	}
}

func (c *Image) JPEG(w io.Writer, options *jpeg.Options) error {
	return jpeg.Encode(w, c.image, options)
}

func (c *Image) PNG(w io.Writer) error {
	return png.Encode(w, c.image)
}

func (c *Image) WriteFile(filename string) error {
	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fp.Close()
	if strings.HasSuffix(strings.ToLower(filename), "."+IMAGE_OUTPUT_JPG) {
		return c.JPEG(fp, DefaultJPEGOptions)
	} else {
		return c.PNG(fp)
	}
}
