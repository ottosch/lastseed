package table

import "log"

type CellSettings struct {
	Align *TextAlign
	Style *lineStyle
}

func parseSettings(settings int) *CellSettings {
	validateSettings(settings)

	cs := &CellSettings{Align: &TextAlign{alignLeft}, Style: getLineStyle(textMiddle)}

	switch {
	case settings&ALIGN_CENTER != 0:
		cs.Align = &TextAlign{alignCenter}
	case settings&ALIGN_RIGHT != 0:
		cs.Align = &TextAlign{alignRight}
	}

	switch {
	case settings&BORDER_TOP != 0:
		cs.Style = getLineStyle(top)
	case settings&BORDER_BOTTOM != 0:
		cs.Style = getLineStyle(bottom)
	case settings&BORDER_MIDDLE != 0:
		cs.Style = getLineStyle(middle)
	}

	return cs
}

func validateSettings(settings int) {
	if !valid(settings, ALIGN_LEFT|ALIGN_CENTER|ALIGN_RIGHT) {
		log.Fatalf("invalid alignment: %d\n", settings)
	}

	if !valid(settings, BORDER_TOP|BORDER_BOTTOM|BORDER_MIDDLE|TEXT_MIDDLE) {
		log.Fatalf("invalid border: %d\n", settings)
	}
}

func valid(settings int, fullMask int) bool {
	propBits := settings & fullMask // isolates property bits
	somePropSet := propBits != 0    // check if some property bit is on

	if somePropSet {
		flippedBits := propBits - 1                  // flips property bits to get only 1's
		commonBitsFound := propBits&flippedBits != 0 // more than 1 bit is set
		return !commonBitsFound
	}

	return true
}
