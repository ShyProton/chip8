package main

import (
	"fmt"

	"github.com/msoap/tcg"
)

type Display tcg.Tcg

func (display *Display) Init() error {
	tcg, err := tcg.New(tcg.Mode1x1)
	if err != nil {
		return fmt.Errorf("could not initialize display: %w", err)
	}

	*display = Display(*tcg)

	return nil
}

func (display *Display) Finish() {
	(*tcg.Tcg)(display).Finish()
}
