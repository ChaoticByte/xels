package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ChaoticByte/xels/simulation"
)


type Application struct {
	simulation simulation.Simulation
}

func (a *Application) Update() error {
	a.simulation.Update()
	return nil
}

func (a *Application) Draw(screen *ebiten.Image) {
	// create image from xel grid
	pixels := make([]byte, 4 * screen.Bounds().Dx() * screen.Bounds().Dy())
	for i := range len(a.simulation.Grid.Xels) {
		xel := a.simulation.Grid.Xels[i]
		var v byte = 0
		if xel.Energy > 0 {
			v = byte(max(0, min(255, xel.Energy * 2 + 38)))
		}
		// r
		pixels[i * 4]     = v
		// g
		pixels[(i*4) + 1] = v
		// b
		pixels[(i*4) + 2] = v
		// a
		pixels[(i*4) + 3] = 255
	}
	screen.WritePixels(pixels)
}

func (a *Application) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return CanvasWidth, CanvasHeight
}


func main() {
	ebiten.SetWindowSize(CanvasWidth, CanvasHeight)
	ebiten.SetWindowTitle("Pixels - Main Window")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()
	ebiten.SetTPS(MaxTps)
	app := &Application{
		simulation: *simulation.NewSimulation(CanvasWidth, CanvasHeight, SimStepsPerUpdate),
	}
	InitGrid(app.simulation.Grid)
	err := ebiten.RunGame(app)
	if err != nil { panic(err) }
}
