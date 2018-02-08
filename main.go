package main

func main() {
	var image *Image
	image = MakeImage(250, 250)
	image.DrawLineOctantI(Color{r: 0, g: 255, b: 0}, 0, 0, 150, 100)
	image.DrawLineOctantI(Color{r: 255, g: 0, b: 0}, 0, 0, 150, 150)
	image.SavePPM("test.ppm")
}
