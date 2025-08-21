package simulation

import (
	"math/rand/v2"
)

type Simulation struct {
	Grid *XelGrid
}

func (sim *Simulation) Update() {
	for range 10 {
		// get all available xels with energy != 0
		available_positions := []Vector2{}
		for i, xel := range sim.Grid.Xels {
			if xel.Energy != 0 {
				available_positions = append(available_positions, Vector2{X: i % sim.Grid.Width, Y: i/sim.Grid.Height})
			}
		}
		// choose random pos
		pos := available_positions[rand.IntN(len(available_positions))]
		xel := sim.Grid.GetXel(pos)
		if xel.Energy == 0 { panic("xel energy null") }
		xel.Step(pos, sim.Grid)
	}
}

func NewSimulation(width int, height int) *Simulation {
	return &Simulation{
		Grid: NewXelGrid(width, height),
	}
}
