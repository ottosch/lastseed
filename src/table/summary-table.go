package table

import (
	"fmt"
	"math"
	"strings"

	"github.com/ottosch/lastseed/src/seed"
)

// DrawSummary prints a summary table
func DrawSummary(s *seed.Seed) {
	colSizes := getSummaryColSizes(s.GetChecksumSize())

	table := make([]*TableRow, 0)
	table = append(table, GridRow(colSizes, BORDER_TOP))

	table = append(table,
		TextRow(
			[]string{
				fmt.Sprintf("%d words input", len(s.GetWords())),
				fmt.Sprintf("%d words total", s.GetWordCount()),
			},
			colSizes, ALIGN_CENTER,
		),
	)

	table = append(table,
		TextRow(
			[]string{
				fmt.Sprintf("%d bits input", len(s.GetWords())*11),
				fmt.Sprintf("%d bits total", s.GetWordCount()*11),
			}, colSizes, ALIGN_CENTER,
		),
	)

	table = append(table,
		TextRow(
			[]string{
				fmt.Sprintf("%d bits entropy", uint(s.GetWordCount()*11)-s.GetChecksumSize()),
				fmt.Sprintf("%d bits checksum", s.GetChecksumSize()),
			}, colSizes, ALIGN_CENTER,
		),
	)

	totalWords := int(math.Pow(2, float64(11-s.GetChecksumSize())))
	table = append(table,
		TextRow(
			[]string{
				fmt.Sprintf("%d missing bits", 11-s.GetChecksumSize()),
				fmt.Sprintf("%d possibilities", totalWords),
			}, colSizes, ALIGN_CENTER,
		),
	)

	pad := getPaddingString(s, colSizes)
	table = append(table, GridRow(colSizes, BORDER_BOTTOM))

	var sb strings.Builder
	for _, tt := range table {
		sb.WriteString(pad)
		sb.WriteString(tt.String())
	}

	fmt.Println(sb.String())
}

func getSummaryColSizes(checksumSize uint) []int {
	const maxWordLength = 8
	const wordCol = maxWordLength + 2     // leading and traili ang space
	entropyCol := int(8*checksumSize) + 2 // use entropy to make this table the same width as one of words table
	lineLength := wordCol + entropyCol

	colSizes := make([]int, 2)
	colSizes[0] = (lineLength) / 2
	colSizes[1] = lineLength - colSizes[0]

	return colSizes
}

// getPaddingString returns a string to centralise the summary table
func getPaddingString(s *seed.Seed, colSizes []int) string {
	if s.GetWordCount() >= 21 {
		return ""
	}

	// get result table size
	const maxWordLength = 8
	const wordCol = maxWordLength + 2
	entropyCol := int(8*s.GetChecksumSize()) + 2

	var tableCount int
	if s.GetWordCount() == 12 {
		tableCount = 4
	} else {
		tableCount = 2
	}

	gridBorders := 3
	tableMargins := tableCount - 1

	size := (wordCol+entropyCol+gridBorders)*tableCount + tableMargins

	// summary table size
	totalColSize := colSizes[0] + colSizes[1] + gridBorders

	centerPosition := (size - totalColSize) / 2
	return strings.Repeat(" ", centerPosition)

}
