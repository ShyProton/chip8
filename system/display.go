package system

import (
	"fmt"

	"github.com/ShyProton/chip8/system/ops"
	"github.com/ShyProton/chip8/system/utils"
	"github.com/msoap/tcg"
)

const (
	DisplayWidth  = 64
	DisplayHeight = 32

	TitleOffset = 3
)

type IO struct {
	graphics         *tcg.Tcg
	originX, originY int
}

func (io *IO) Init(romPath string) error {
	mode := tcg.Mode1x2

	graphics, err := tcg.New(mode)
	if err != nil {
		return fmt.Errorf("could not initialize display graphics: %w", err)
	}

	io.graphics = graphics

	io.initPattern()
	io.setOriginChars()

	err = io.graphics.SetClipCenter(DisplayWidth, DisplayHeight)
	if err != nil {
		io.graphics.Finish()

		return fmt.Errorf("window is too small to initialize graphics: %w", err)
	}

	io.graphics.Buf.Rect(0, 0, io.graphics.Width, io.graphics.Height, tcg.White)

	io.initTitles("CHIP-8", romPath)

	return nil
}

func (io *IO) initPattern() {
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
}

func (io *IO) initTitles(title, subtitle string) {
	centeredTitle := centerStringToDisplay(title)
	centeredSubtitle := centerStringToDisplay(subtitle)

	io.graphics.PrintStr(io.originX, io.originY-TitleOffset, centeredTitle)
	io.graphics.PrintStr(io.originX, io.originY-TitleOffset+1, centeredSubtitle)

	io.graphics.Show()
}

func centerStringToDisplay(str string) string {
	padding := (DisplayWidth - len(str)) / 2
	oddOffset := len(str) % 2

	return fmt.Sprintf("%*s%s%*s", padding, "", str, padding+oddOffset, "")
}

func (io *IO) setOriginChars() {
	screenWidthChars, screenHeightChars := io.graphics.ScreenSize()

	originCharsX := (screenWidthChars - DisplayWidth) / 2
	originCharsY := (screenHeightChars - DisplayHeight) / 2

	io.originX, io.originY = originCharsX, originCharsY
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

	for i := range ops.BitsPerByte {
		drawX, drawY := getDrawXY(x+i, y)

		newPixel := utils.GetBinaryDigit(row, ops.BitsPerByte-1-i)
		oldPixel := io.graphics.Buf.At(drawX, drawY)

		newPixel ^= oldPixel

		if newPixel == 0 && oldPixel == 1 {
			erasure = true
		}

		io.graphics.Buf.Set(drawX, drawY, newPixel)
	}

	return erasure
}
