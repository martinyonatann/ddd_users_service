package serr

import (
	"fmt"
	"reflect"
)

// color type
type color int

// colorType type
type colorType int

const (
	// colorTypeForeground constant for foreground color type
	colorTypeForeground colorType = 1 + iota

	// colorTypeBackground constant for background color type
	colorTypeBackground
)

const (
	// colorDeafult constant for default color
	colorDeafult color = 1 + iota

	// colorBlack constant for black color
	colorBlack

	// colorRed constant for red color
	colorRed

	// colorGreen constant for green color
	colorGreen

	// colorYellow constant for yellow color
	colorYellow

	// colorBlue constant for blue color
	colorBlue

	// colorMagenta constant for magenta color
	colorMagenta

	// colorCyan constant for cyan color
	colorCyan

	// colorLightGray constant for light gray color
	colorLightGray

	// colorDarkGray constant for dark gray color
	colorDarkGray

	// colorLightRed constant for light red color
	colorLightRed

	// colorLightGreen constant for light green color
	colorLightGreen

	// colorLightYellow constant for light yellow color
	colorLightYellow

	// colorLightBlue constant for light blue color
	colorLightBlue

	// colorLightMagenta constant for light magenta color
	colorLightMagenta

	// colorLightCyan constant for light cyan color
	colorLightCyan

	// colorWhite constant for white color
	colorWhite
)

const (
	// escChar constant
	escChar = "\x1B"

	// resetChar constant
	resetChar = escChar + "[0m"
)

// getColorCode function
func getColorCode(c color, t colorType) (string, bool) {
	fcs := map[color][]string{
		colorDeafult:      []string{"39", "49"},
		colorBlack:        []string{"30", "40"},
		colorRed:          []string{"31", "41"},
		colorGreen:        []string{"32", "42"},
		colorYellow:       []string{"33", "43"},
		colorBlue:         []string{"34", "44"},
		colorMagenta:      []string{"35", "45"},
		colorCyan:         []string{"36", "46"},
		colorLightGray:    []string{"37", "47"},
		colorDarkGray:     []string{"90", "100"},
		colorLightRed:     []string{"91", "101"},
		colorLightGreen:   []string{"92", "102"},
		colorLightYellow:  []string{"93", "103"},
		colorLightBlue:    []string{"94", "104"},
		colorLightMagenta: []string{"95", "105"},
		colorLightCyan:    []string{"96", "106"},
		colorWhite:        []string{"97", "107"},
	}

	if colorCode, ok := fcs[c]; ok {
		switch t {
		case colorTypeForeground:
			return colorCode[0], true

		case colorTypeBackground:
			return colorCode[1], true
		}
	}
	return "", false
}

// applyForeColor function
func applyForeColor(msg string, color color) string {
	if color, ok := getColorCode(color, colorTypeForeground); ok {
		return fmt.Sprintf("%s[%sm%s%s", escChar, color, msg, resetChar)
	}
	return msg
}

// isExists function to check are is exists in array?
func isExists(value interface{}, array interface{}) (exist bool) {
	exist = false
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) == true {
				exist = true
				return exist
			}
		}
	}

	return exist
}
