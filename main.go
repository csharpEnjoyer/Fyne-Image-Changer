package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

//Size of window
var (
	height int = 200
	width  int = 100
)

//Update function IMPORTANT!
func updateImage(img *canvas.Image, i uint8) {
	img.Image = generateImage(i)
	time.Sleep(time.Second * 1)
	img.Refresh()
}

func main() {
	a := app.New()
	w := a.NewWindow("Animation")

	img := canvas.NewImageFromImage(generateImage(1))

	updateImage(img, 1)
	w.Resize(fyne.NewSize(float32(height), float32(width)))
	w.SetContent(img)

	go func() {
		var i uint8 = 1
		for i < 10 {
			updateImage(img, i*10)
			fmt.Println(i)
			i++
		}
	}()

	w.ShowAndRun()
}

//image genearation function
func generateImage(seed uint8) image.Image {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			col := color.RGBA{1 * seed, 2 * seed, 3 * seed, 0xff}
			img.Set(x, y, col)
		}
	}
	return img
}
