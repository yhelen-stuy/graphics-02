package main

func main() {
	var image *Image
	image = MakeImage(500, 500)
	// image.DrawLine(Color{r: 0, g: 255, b: 0}, 250, 250, 400, 350)
	// image.DrawLine(Color{r: 255, g: 0, b: 0}, 250, 250, 400, 400)
	// image.DrawLine(Color{r: 127, g: 0, b: 127}, 250, 250, 450, 300)

	// image.DrawLine(Color{r: 0, g: 0, b: 255}, 250, 250, 300, 400)

	// image.DrawLine(Color{r: 122, g: 127, b: 0}, 250, 250, 400, 150)

	// image.DrawLine(Color{r: 0, g: 127, b: 127}, 250, 250, 300, 100)

	image.DrawLine(Color{r: 0, g: 0, b: 0}, 0, 250, 500, 250)
	image.DrawLine(Color{r: 0, g: 0, b: 0}, 250, 0, 250, 500)
	image.SavePPM("lines.ppm")
}
