package simulation

import (
	"math/rand/v2"
)

const (
	TurnNorth int = iota
	TurnEast
	TurnSouth
	TurnWest
)

type Xel struct {
	Energy int64
}

func (xel *Xel) Step(ownPos Vector2, grid *XelGrid) {
	// get other xel
	var otherXelPos Vector2
	switch rand.IntN(4) {
	case 0:
		otherXelPos = Vector2{X: ownPos.X + 1, Y: ownPos.Y}
	case 1:
		otherXelPos = Vector2{X: ownPos.X - 1, Y: ownPos.Y}
	case 2:
		otherXelPos = Vector2{X: ownPos.X, Y: ownPos.Y + 1}
	case 3:
		otherXelPos = Vector2{X: ownPos.X, Y: ownPos.Y - 1}
	}
	otherXel := grid.GetXel(otherXelPos)
	if otherXel == nil { return }
	// interact
	if (xel.Energy > 0 && otherXel.Energy > 0) || (xel.Energy < 0 && otherXel.Energy < 0) {
		if rand.IntN(4) == 0 {
			// create a higher energy xel and an inverted xel
			xel.Energy += (2 * otherXel.Energy)
			otherXel.Energy *= -1
		} else {
			xel.Energy += otherXel.Energy
			otherXel.Energy = 0
		}
		return
	}
	if (xel.Energy < 0 && otherXel.Energy > 0) || (xel.Energy > 0 && otherXel.Energy < 0) {
		// cancel out
		xel.Energy += otherXel.Energy
		otherXel.Energy = 0
		return
	}
	if (xel.Energy == 0 && otherXel.Energy != 0) || (xel.Energy != 0 && otherXel.Energy == 0) {
		if xel.Energy > 1 || xel.Energy < -1 || otherXel.Energy > 1 || otherXel.Energy < -1 { // prevent integer 1 / 0
			// split
			splitE := (xel.Energy + otherXel.Energy) / 2
			xel.Energy = splitE + ((xel.Energy + otherXel.Energy) % 2)
			otherXel.Energy = splitE
		} else {
			// move
			if xel.Energy != 0 {
				otherXel.Energy = xel.Energy
				xel.Energy = 0
			} else {
				xel.Energy = otherXel.Energy
				otherXel.Energy = 0
			}
		}
		return
	}
}
