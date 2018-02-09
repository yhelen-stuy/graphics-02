package main

func main() {
	var image *Image
	image = MakeImage(500, 500)
	image.DrawLineOctantI(Color{r: 0, g: 255, b: 0}, 250, 250, 400, 350)
	image.DrawLineOctantI(Color{r: 255, g: 0, b: 0}, 250, 250, 400, 400)
	image.DrawLineOctantI(Color{r: 0, g: 0, b: 0}, 250, 250, 450, 300)

	image.DrawLineOctantII(Color{r: 0, g: 0, b: 255}, 250, 250, 300, 400)

	image.DrawLineOctantVIII(Color{r: 122, g: 122, b: 0}, 250, 250, 400, 150)

	image.DrawLineOctantVII(Color{r: 0, g: 122, b: 122}, 250, 250, 300, 100)
	image.SavePPM("lines.ppm")
}
