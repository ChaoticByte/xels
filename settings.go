package main

import "github.com/ChaoticByte/xels/simulation"

const CanvasWidth = 200
const CanvasHeight = 200
const MaxTps = 1000
const SimStepsPerUpdate = 10

func InitGrid(grid *simulation.XelGrid) {
	grid.GetXel(grid.GetCenterPosition()).Energy = 50000
}
