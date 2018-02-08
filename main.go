package main

func main() {
	var image *Image
	image = MakeImage(500, 500)
	image.SavePPM("test.ppm")
}
