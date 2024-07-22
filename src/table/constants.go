package table

const (
	topLeft         = '┌'
	topSeparator    = '┬'
	topRight        = '┐'
	bottomLeft      = '└'
	bottomSeparator = '┴'
	bottomRight     = '┘'
	middleLeft      = '├'
	middleSeparator = '┼'
	middleRight     = '┤'
	vertical        = '│'
	horizontal      = '─'

	wordStr    = "Word"
	entropyStr = "Entropy"

	ALIGN_LEFT   = 1 << iota // 1 << 0 == 1
	ALIGN_CENTER             // 1 << 1 == 2
	ALIGN_RIGHT              // 1 << 2 == 4

	BORDER_TOP
	BORDER_BOTTOM
	BORDER_MIDDLE
	TEXT_MIDDLE
)

var (
	top        lineType = 1
	bottom     lineType = 2
	middle     lineType = 3
	textMiddle lineType = 4

	alignLeft   alignType = 1
	alignCenter alignType = 2
	alignRight  alignType = 3
)

type lineType int

type alignType int
