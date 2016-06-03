package ewa

import "sort"

type byCorrection func(m1, m2 Correction) bool

type correctionSorter struct {
	waves  Corrections
	sortFn byCorrection
}

func (ws *correctionSorter) Len() int {
	return len(ws.waves)
}

func (ws *correctionSorter) Swap(i, j int) {
	ws.waves[i], ws.waves[j] = ws.waves[j], ws.waves[i]
}

func (ws *correctionSorter) Less(i, j int) bool {
	return ws.sortFn(ws.waves[i], ws.waves[j])
}

func newCorrectionSorter(selector SelectorCorrections, fn byCorrection) *correctionSorter {
	return &correctionSorter{
		waves:  selector.Corr,
		sortFn: fn,
	}
}

//ByDegree sorting of impulses
func (in SelectorCorrections) ByDegree(asc bool) SelectorCorrections {
	sorter := newCorrectionSorter(in, func(i, j Correction) bool {
		if asc {
			return i.Degree() < j.Degree()
		}
		return i.Degree() > j.Degree()
	})

	sort.Sort(sorter)
	return in
}

//ByDuration sorting of impulses
func (in SelectorCorrections) ByDuration(asc bool) SelectorCorrections {
	sorter := newCorrectionSorter(in, func(i, j Correction) bool {
		if asc {
			return i.Duration() < j.Duration()
		}

		return i.Duration() > j.Duration()
	})

	sort.Sort(sorter)
	return in
}

//ByLen sorting of impulses
func (in SelectorCorrections) ByLen(asc bool) SelectorCorrections {
	sorter := newCorrectionSorter(in, func(i, j Correction) bool {
		if asc {
			return i.Len() < j.Len()
		}

		return i.Len() > j.Len()
	})

	sort.Sort(sorter)
	return in
}

//ByRetrace sorting of impulses
func (in SelectorCorrections) ByRetrace(val float64, asc bool) SelectorCorrections {
	sorter := newCorrectionSorter(in, func(i, j Correction) bool {
		if asc {
			if i.Up() {
				return i.Retrace(val) < j.Retrace(val)
			}
			return i.Retrace(val) > j.Retrace(val)
		}

		if i.Up() {
			return i.Retrace(val) > j.Retrace(val)
		}

		return i.Retrace(val) < j.Retrace(val)
	})

	sort.Sort(sorter)
	return in
}

//ByBegins sorting of impulses
func (in SelectorCorrections) ByBegins(asc bool) SelectorCorrections {
	sorter := newCorrectionSorter(in, func(i, j Correction) bool {
		if asc {
			return j.Begins().After(i.Begins())
		}

		return i.Begins().After(j.Begins())
	})

	sort.Sort(sorter)
	return in
}

//ByEnds sorting of impulses
func (in SelectorCorrections) ByEnds(asc bool) SelectorCorrections {
	sorter := newCorrectionSorter(in, func(i, j Correction) bool {
		if asc {
			return j.Ends().After(i.Ends())
		}

		return i.Ends().After(j.Ends())
	})

	sort.Sort(sorter)
	return in
}

//ByStarts sorting of impulses
func (in SelectorCorrections) ByStarts(asc bool) SelectorCorrections {
	sorter := newCorrectionSorter(in, func(i, j Correction) bool {
		if asc {
			if i.Up() {
				return i.Starts() < j.Starts()
			}
			return i.Starts() > j.Starts()
		}

		if i.Up() {
			return i.Starts() > j.Starts()
		}
		return i.Starts() < j.Starts()
	})

	sort.Sort(sorter)
	return in
}

//ByTops sorting of impulses
func (in SelectorCorrections) ByTops(asc bool) SelectorCorrections {
	sorter := newCorrectionSorter(in, func(i, j Correction) bool {
		if asc {
			if i.Up() {
				return i.Tops() < j.Tops()
			}
			return i.Tops() < j.Tops()
		}

		if i.Up() {
			return i.Tops() > j.Tops()
		}
		return i.Tops() > j.Tops()
	})

	sort.Sort(sorter)
	return in
}

//BySlope sorting of impulses
func (in SelectorCorrections) BySlope(asc bool) SelectorCorrections {
	sorter := newCorrectionSorter(in, func(i, j Correction) bool {
		if asc {
			return i.Slope() < j.Slope()
		}

		return i.Slope() > j.Slope()
	})

	sort.Sort(sorter)
	return in
}
