package main

import (
	"fmt"

	"github.com/msoap/tcg"
)

const (
	DisplayWidth  = 64
	DisplayHeight = 32

	TitleOffset = 3
)

type IO struct {
	graphics *tcg.Tcg
}

func (io *IO) Init(romName string) error {
	mode := tcg.Mode1x2

	graphics, err := tcg.New(mode)
	if err != nil {
		return fmt.Errorf("could not initialize display graphics: %w", err)
	}

	io.graphics = graphics

	pattern := tcg.MustNewBufferFromStrings([]string{
		"**  **",
		" **** ",
		"*    *",
		" **** ",
		"**  **",
		"  **  ",
	})

	io.graphics.Buf.Fill(0, 0, tcg.WithPattern(pattern))
	io.graphics.Show()

	originX, originY := io.getOriginChars()

	err = io.graphics.SetClipCenter(DisplayWidth, DisplayHeight)
	if err != nil {
		io.graphics.Finish()

		return fmt.Errorf("window is too small to initialize graphics: %w", err)
	}

	io.graphics.Buf.Rect(0, 0, io.graphics.Width, io.graphics.Height, tcg.White)

	title := centerStringToDisplay("CHIP-8")
	subTitle := centerStringToDisplay(romName)

	io.graphics.PrintStr(originX, originY-TitleOffset, title)
	io.graphics.PrintStr(originX, originY-TitleOffset+1, subTitle)

	io.graphics.Show()

	return nil
}

func centerStringToDisplay(str string) string {
	padding := (DisplayWidth - len(str)) / 2

	return fmt.Sprintf("%*s%s%*s", padding, "", str, padding, "")
}

func (io *IO) getOriginChars() (int, int) {
	screenWidthChars := io.graphics.Width / tcg.Mode1x2.Width()
	screenHeightChars := io.graphics.Height / tcg.Mode1x2.Height()

	originCharsX := (screenWidthChars - DisplayWidth) / 2
	originCharsY := (screenHeightChars - DisplayHeight) / 2

	return originCharsX, originCharsY
}

func getDrawXY(x, y int) (int, int) {
	drawX, drawY := x, y

	if drawX >= DisplayWidth {
		drawX -= DisplayWidth
	}

	if drawY >= DisplayHeight {
		drawY -= DisplayHeight
	}

	return drawX, drawY
}

func (io *IO) DrawRow(x int, y int, row byte) bool {
	var erasure bool

	for i := range BitsPerByte {
		drawX, drawY := getDrawXY(x+i, y)

		newPixel := GetBinaryDigit(row, BitsPerByte-i-1)
		oldPixel := io.graphics.Buf.At(drawX, drawY)

		newPixel ^= oldPixel

		if newPixel == 0 && oldPixel == 1 {
			erasure = true
		}

		io.graphics.Buf.Set(drawX, drawY, newPixel)
	}

	return erasure
}
