all:
	go run image.go main.go

run:
	display lines.ppm

clean:
	rm *.ppm
