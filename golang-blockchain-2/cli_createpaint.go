package main

import (
	"fmt"
)

func (cli *CLI) createPaint(paletteId, nodeID string) {

	palettes, _ := NewPalettes(nodeID)

	paintId := palettes.CreatePaint(paletteId)

	palettes.SaveToFile(nodeID)

	fmt.Printf("Paint generated '%s'", paintId)
}
