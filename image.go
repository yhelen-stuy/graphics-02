package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

type Color struct {
	r int
	g int
	b int
}

type Image struct {
	img    [][]Color
	height int
	width  int
}

func MakeImage(height, width int) *Image {
	img := make([][]Color, height)
	for i := range img {
		img[i] = make([]Color, width)
	}
	image := &Image{
		img:    img,
		height: height,
		width:  width,
	}
	image.Clear()
	return image
}

func (image Image) plot(c Color, x, y int) error {
	if x < 0 || x > image.height || y < 0 || y > image.width {
		return errors.New("Error: Coordinate invalid")
	}
	image.img[x][y] = c
	return nil
}

func (image Image) fill(c Color) {
	for x := 0; x < image.height; x++ {
		for y := 0; y < image.width; y++ {
			image.plot(c, x, y)
		}
	}
}

func (image Image) Clear() {
	image.fill(Color{r: 255, g: 255, b: 255})
}

func (image Image) SavePPM(filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	// TODO: Take variant max color
	buffer.WriteString(fmt.Sprintf("P3 %d %d 255\n", image.height, image.width))

	for x := 0; x < image.height; x++ {
		for y := 0; y < image.width; y++ {
			buffer.WriteString(fmt.Sprintf("%d %d %d\n", image.img[x][y].r, image.img[x][y].g, image.img[x][y].b))
		}
	}

	f.WriteString(buffer.String())
	f.Close()
	return nil
}

func (image Image) DrawLineOctantI(c Color, x0, y0, x1, y1 int) error {
	y := y0
	lA := y1 - y0
	lB := x0 - x1
	lD := 2*lA + lB
	for x := x0; x < x1; x++ {
		err := image.plot(c, x, y)
		if err != nil {
			return err
		}
		if lD > 0 {
			y++
			lD += 2 * lB
		}
		x++
		lD += 2 * lA
	}
	return nil
}
