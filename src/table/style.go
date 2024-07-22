package table

import (
	"strings"
)

// TextAlign contains a text align of alignType
type TextAlign struct {
	align alignType
}

// AlignText aligns the text, padding it with size and according to alignType
func (ta *TextAlign) AlignText(text string, size int) string {
	padding := size - len(text)
	switch ta.align {
	case alignLeft:
		return " " + text + strings.Repeat(" ", padding-1)
	case alignCenter:
		spacesLeft := strings.Repeat(" ", padding/2)
		spacesRight := strings.Repeat(" ", padding-len(spacesLeft))
		return spacesLeft + text + spacesRight
	default:
		return strings.Repeat(" ", padding-1) + text + " "
	}
}

// lineStyle style of table row, with borders and between-cell separator
type lineStyle struct {
	leftBorder  rune
	separator   rune
	rightBorder rune
}

// getLineStyle returns the suitable grid runes for the lineType
func getLineStyle(t lineType) *lineStyle {
	switch t {
	case top:
		return &lineStyle{leftBorder: topLeft, separator: topSeparator, rightBorder: topRight}
	case middle:
		return &lineStyle{leftBorder: middleLeft, separator: middleSeparator, rightBorder: middleRight}
	case textMiddle:
		return &lineStyle{leftBorder: vertical, separator: vertical, rightBorder: vertical}
	default:
		return &lineStyle{leftBorder: bottomLeft, separator: bottomSeparator, rightBorder: bottomRight}
	}
}
