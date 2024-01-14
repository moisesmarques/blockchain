package main

import (
	"fmt"
)

func (cli *CLI) createPalette(nodeID string) {

	palettes, _ := NewPalettes(nodeID)

	paletteId := palettes.CreatePalette()

	palettes.SaveToFile(nodeID)

	fmt.Printf("Palette generated '%s'", paletteId)
}
