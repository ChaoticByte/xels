package simulation

type XelGrid struct {
	Width int  // do not overwrite this value after init!
	Height int // do not overwrite this value after init!
	Xels []*Xel
}

func (grid *XelGrid) GetCenterPosition() Vector2 { // just a helper
	return Vector2{
		X: (grid.Width / 2) + (grid.Width % 2),
		Y: (grid.Height / 2) + (grid.Height % 2),
	}
}

func (grid *XelGrid) GetXel(pos Vector2) *Xel {
	if pos.X < 0 || pos.X >= grid.Width || pos.Y < 0 || pos.Y >= grid.Height {
		return nil // out of bounds
	}
	i := (pos.Y * grid.Width) + pos.X
	return grid.Xels[i]
}

func NewXelGrid(width int, height int) *XelGrid {
	grid := &XelGrid{
		Width: width,
		Height: height,
		Xels: make([]*Xel, width*height),
	}
	// init
	for i := range len(grid.Xels) {
		grid.Xels[i] = &Xel{}
	}
	//
	return grid
}
