package table

import (
	"fmt"
	"strings"

	"github.com/ottosch/lastseed/src/seed"
)

// DrawSummary prints the results table
func DrawResults(s *seed.Seed) {
	tableCount, linesPerTable := getTableCountAndHeight(s)
	data := make([][]*seed.Result, 0, tableCount)

	results := s.GetResults()
	var i int
	for i < len(results) {
		data = append(data, results[i:i+linesPerTable])
		i += linesPerTable
	}

	colSizes := getResultsColSizes(s.GetChecksumSize(), tableCount)

	table := make([]*TableRow, 0)
	table = append(table, GridRow(colSizes, BORDER_TOP))

	headers := make([]string, len(data)*2)
	for i := 0; i < len(headers); i += 2 {
		headers[i] = wordStr
		headers[i+1] = entropyStr
	}

	table = append(table, TextRow(headers, colSizes, ALIGN_CENTER|TEXT_MIDDLE))
	table = append(table, GridRow(colSizes, BORDER_MIDDLE))

	limit := len(data[0])
	for i := 0; i < limit; i++ {
		bodyCells := make([]string, 0, len(data)*2)
		for j := 0; j < len(data); j++ {
			bodyCells = append(bodyCells, data[j][i].LastWord())
			bodyCells = append(bodyCells, data[j][i].Entropy(s.GetChecksumSize()))
		}

		table = append(table, TextRow(bodyCells, colSizes, ALIGN_LEFT|TEXT_MIDDLE))
	}

	table = append(table, GridRow(colSizes, BORDER_BOTTOM))

	var sb strings.Builder
	for _, tt := range table {
		sb.WriteString(tt.String())
	}

	fmt.Println(sb.String())
}

func getTableCountAndHeight(s *seed.Seed) (tables int, linesPerTable int) {
	switch s.GetWordCount() {
	case 12:
		tables = 4
	case 15, 18:
		tables = 2
	default:
		tables = 1
	}

	linesPerTable = len(s.GetResults()) / tables
	return
}

func getResultsColSizes(checksumSize uint, tableCount int) []int {
	const maxWordLength = 8
	const colPadding = 2
	wordColLength := maxWordLength + colPadding
	entropyColLength := 8*int(checksumSize) + colPadding

	colSizes := make([]int, tableCount*2)
	for i := 0; i < len(colSizes); i += 2 {
		colSizes[i] = wordColLength
		colSizes[i+1] = entropyColLength
	}

	return colSizes
}
