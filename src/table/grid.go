package table

import (
	"log"
	"strings"
)

// TableCell a single table cell and its text
type TableCell struct {
	text string
}

// NewCell creates a table cell, content aligned
func NewCell(text string, length int, align *TextAlign) *TableCell {
	cellText := align.AlignText(text, length)
	return &TableCell{text: cellText}
}

// TableRow a table row, composed of various, styled cells
type TableRow struct {
	cells []*TableCell
	style *lineStyle
}

// TextRow creates a row of text cells
func TextRow(texts []string, sizes []int, settings int) *TableRow {
	if len(texts) != len(sizes) || len(texts)%2 != 0 {
		log.Fatalf("invalid number of columns: %d, %d\n", len(texts), len(sizes))
	}

	cellSettings := parseSettings(settings)
	cells := make([]*TableCell, len(texts))

	for i, text := range texts {
		cells[i] = NewCell(text, sizes[i], cellSettings.Align)
	}

	return &TableRow{cells: cells, style: cellSettings.Style}
}

// GridRow creates a row of grid (border) content only
func GridRow(sizes []int, settings int) *TableRow {
	if len(sizes)%2 != 0 {
		log.Fatalf("invalid number of grid cells: %v\n", sizes)
	}

	cellSettings := parseSettings(settings)
	cells := make([]*TableCell, len(sizes))

	for i, size := range sizes {
		cells[i] = &TableCell{text: strings.Repeat(string(horizontal), size)}
	}

	return &TableRow{cells: cells, style: cellSettings.Style}
}

func (g *TableRow) String() string {
	var result strings.Builder

	strList := make([]string, 0, len(g.cells)/2)
	for i := 0; i < len(g.cells); i += 2 {
		var sb strings.Builder
		sb.WriteRune(g.style.leftBorder)
		sb.WriteString(g.cells[i].text)
		sb.WriteRune(g.style.separator)
		sb.WriteString(g.cells[i+1].text)
		sb.WriteRune(g.style.rightBorder)
		strList = append(strList, sb.String())
	}

	result.WriteString(strings.Join(strList, " "))
	result.WriteRune('\n')
	return result.String()
}
