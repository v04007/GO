package main

// import fyne
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
)

func main() {
	// New app
	a := app.New()
	// new windown and title
	w := a.NewWindow("Scroll")
	// resize
	w.Resize(fyne.NewSize(400, 400))
	// red rectangle
	red := canvas.NewRectangle(
		color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	red.SetMinSize(fyne.NewSize(400, 400))
	// blue rectangle
	blue := canvas.NewRectangle(
		color.NRGBA{R: 0, G: 0, B: 255, A: 255})
	blue.SetMinSize(fyne.NewSize(400, 400))
	// New Scroll & Vbox
	c := container.NewVScroll(
		container.NewHBox( // New Horizontal Box
			red,
			blue,
		),
	)
	// Change scroll direction
	c.Direction = container.ScrollHorizontalOnly
	// setup content
	w.SetContent(c)
	// show and run
	w.ShowAndRun()
}
