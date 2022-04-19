package model

// Winner returns X, O or EMPTY in case no winner
func (b *Board) Winner() FieldContents {
	wins := []struct {
		p1, p2, p3 int
	}{
		{1, 2, 3}, {1, 4, 7}, {1, 5, 9},
		{2, 5, 8}, {3, 6, 9}, {3, 5, 7},
		{4, 5, 6},
		{7, 8, 9},
	}

	for _, w := range wins {
		// only check for X, O
		if b.Field[w.p1] == X || b.Field[w.p1] == O {
			if b.Field[w.p1] == b.Field[w.p2] && b.Field[w.p2] == b.Field[w.p3] {
				return b.Field[w.p1]
			}
		}
	}

	return EMPTY
}
