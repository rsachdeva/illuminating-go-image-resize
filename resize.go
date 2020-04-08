package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"log"
)

func main() {
	fmt.Println("Hello World.")
	// Open a test image.
	img, err := imaging.Open("assets/ScreenShot-RustGoP.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	// Resize preserving the aspect ratio.
	fmt.Println("Before resize dimensions are", img.Bounds());
	// The color method returns the image's `ColorType`.
	fmt.Printf("RGBA color code %v\n", img.At(img.Bounds().Min.X, img.Bounds().Min.Y));

	scaled := imaging.Resize(img, 1400, 2250, imaging.Lanczos)
	// The dimensions method returns the images width and height.
	fmt.Println("After resize dimensions are", scaled.Bounds());

	// Save the resulting image as JPEG.
	err = imaging.Save(scaled, "assets/ScreenShot-RustGoJ.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}