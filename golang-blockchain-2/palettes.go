package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const paletteFile = "palette_%s.dat"

// Palettes stores a collection of palettes
type Palettes struct {
	Palettes map[string]*Palette
}

// NewPalettes creates Palettes and fills it from a file if it exists
func NewPalettes(nodeID string) (*Palettes, error) {
	palettes := Palettes{}
	palettes.Palettes = make(map[string]*Palette)

	err := palettes.LoadFromFile(nodeID)

	return &palettes, err
}

// CreatePalette adds a Palette to Palettes
func (ps *Palettes) CreatePalette() string {
	palette := NewPalette()
	ps.Palettes[palette.Id] = palette

	return palette.Id
}

// GetPalette returns a Palette by id
func (ps Palettes) GetPalette(id string) Palette {
	return *ps.Palettes[id]
}

func (ps *Palettes) CreatePaint(paletteId string) string {

	palette := ps.Palettes[paletteId]

	paint := NewPaint()

	palette.Paints[paint.Id] = paint

	return paint.Id
}

// LoadFromFile loads palettes from the file
func (ps *Palettes) LoadFromFile(nodeID string) error {
	paletteFile := fmt.Sprintf(paletteFile, nodeID)
	if _, err := os.Stat(paletteFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(paletteFile)
	if err != nil {
		log.Panic(err)
	}

	var palettes Palettes
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&palettes)
	if err != nil {
		log.Panic(err)
	}

	ps.Palettes = palettes.Palettes

	return nil
}

// SaveToFile saves palettes to a file
func (ps Palettes) SaveToFile(nodeID string) {
	var content bytes.Buffer
	paletteFile := fmt.Sprintf(paletteFile, nodeID)

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ps)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(paletteFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
