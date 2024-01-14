package main

func (cli *CLI) getPalette(paletteId, nodeID string) {

	palettes, _ := NewPalettes(nodeID)

	palette := palettes.Palettes[paletteId]

	PrintJson(palette)
}
