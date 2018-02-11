package main

func main() {
	var image *Image
	image = MakeImage(499, 499)
	// Octant I (Or slope 1 which may not be I)
	image.DrawLine(Color{r: 0, g: 255, b: 0}, 250, 250, 400, 350)
	image.DrawLine(Color{r: 255, g: 0, b: 0}, 250, 250, 400, 400)
	image.DrawLine(Color{r: 127, g: 0, b: 127}, 250, 250, 450, 300)

	// // Octant II
	image.DrawLine(Color{r: 0, g: 0, b: 255}, 250, 250, 300, 400)

	// Octant III
	image.DrawLine(Color{r: 127, g: 127, b: 127}, 250, 250, 100, 450)

	// Octant IV
	image.DrawLine(Color{r: 127, g: 127, b: 127}, 250, 250, 100, 300)

	// Octant V
	image.DrawLine(Color{r: 127, g: 127, b: 127}, 250, 250, 100, 200)

	// Octant VI
	image.DrawLine(Color{r: 127, g: 127, b: 127}, 250, 250, 100, 0)

	// Octant VII
	image.DrawLine(Color{r: 127, g: 127, b: 0}, 250, 250, 300, 150)

	// Octant VIII
	image.DrawLine(Color{r: 0, g: 127, b: 127}, 250, 250, 450, 150)

	image.DrawLine(Color{r: 0, g: 0, b: 0}, 0, 250, 499, 250)
	image.DrawLine(Color{r: 0, g: 0, b: 0}, 250, 0, 250, 499)

	image.DrawLine(Color{r: 127, g: 127, b: 127}, 100, 0, 250, 250)
	image.SavePPM("lines.ppm")
}
