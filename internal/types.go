package internal

import (
	"errors"
	"strconv"
	"strings"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

func NewColorFromHex(hex string) (Color, error) {
	var color Color

	hex = strings.ToLower(hex)
	if hex[0] == '#' {
		hex = hex[1:]
	}

	red_value, err_red := strconv.ParseUint(hex[0:2], 16, 8)
	green_value, err_green := strconv.ParseUint(hex[2:4], 16, 8)
	blue_value, err_blue := strconv.ParseUint(hex[4:6], 16, 8)

	if err_red != nil || err_green != nil || err_blue != nil {
		return color, errors.New("Malformed hex value.")
	}

	color = Color{R: uint8(red_value), G: uint8(green_value), B: uint8(blue_value)}

	return color, nil
}

func NewColorFromRGB(r_value uint8, g_value uint8, b_value uint8) Color {
	return Color{R: r_value, G: g_value, B: b_value}
}

/* Default colors */

type StandardColor string

const (
	Reset StandardColor = "reset"

	// Foreground
	FgBlack   = "black"
	FgRed     = "red"
	FgGreen   = "green"
	FgYellow  = "yellow"
	FgBlue    = "blue"
	FgMagenta = "magenta"
	FgCyan    = "cyan"
	FgWhite   = "white"
	FgDefault = "default"

	// Background
	BgBlack   = "b;black"
	BgRed     = "b;red"
	BgGreen   = "b;green"
	BgYellow  = "b;yellow"
	BgBlue    = "b;blue"
	BgMagenta = "b;magenta"
	BgCyan    = "b;cyan"
	BgWhite   = "b;white"
	BgDefault = "b;default"

	// Styles
	StyleBold          = "bold"
	StyleDim           = "dim"
	StyleItalic        = "italic"
	StyleUnderline     = "underline"
	StyleBlink         = "blink"
	StyleInverse       = "inverse"
	StyleInvisible     = "invisible"
	StyleStrikethrough = "strikethrough"
)

var StandardColorList = map[string]uint8{
	"reset": 0,

	"black":   30,
	"red":     31,
	"green":   32,
	"yellow":  33,
	"blue":    34,
	"magenta": 35,
	"cyan":    36,
	"white":   37,
	"default": 38,

	"b;black":   40,
	"b;red":     41,
	"b;green":   42,
	"b;yellow":  43,
	"b;blue":    44,
	"b;magenta": 45,
	"b;cyan":    46,
	"b;white":   47,
	"b;default": 49,

	"bold":          1,
	"dim":           2,
	"italic":        3,
	"underline":     4,
	"blink":         5,
	"inverse":       7,
	"invisible":     8,
	"strikethrough": 9,
}

func (standard_color StandardColor) StandardColorToId() uint8 {
	return StandardColorList[string(standard_color)]
}

/* Parser */
type ColorTagType uint8

const (
	ColorTagHex ColorTagType = iota
	ColorTagStandard
	ColorTagStyle
	ColorTagCombined
)
