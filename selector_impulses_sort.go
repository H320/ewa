package ewa

import "sort"

type byImpulse func(m1, m2 Impulse) bool

type impulseSorter struct {
	waves  Impulses
	sortFn byImpulse
}

func (ws *impulseSorter) Len() int {
	return len(ws.waves)
}

func (ws *impulseSorter) Swap(i, j int) {
	ws.waves[i], ws.waves[j] = ws.waves[j], ws.waves[i]
}

func (ws *impulseSorter) Less(i, j int) bool {
	return ws.sortFn(ws.waves[i], ws.waves[j])
}

func newWaverSorter(selector SelectorImpulses, fn byImpulse) *impulseSorter {
	return &impulseSorter{
		waves:  selector.Imp,
		sortFn: fn,
	}
}

//ByDegree sorting of impulses
func (in SelectorImpulses) ByDegree(asc bool) SelectorImpulses {
	sorter := newWaverSorter(in, func(i, j Impulse) bool {
		if asc {
			return i.Degree() < j.Degree()
		}
		return i.Degree() > j.Degree()
	})

	sort.Sort(sorter)
	return in
}

//ByDuration sorting of impulses
func (in SelectorImpulses) ByDuration(asc bool) SelectorImpulses {
	sorter := newWaverSorter(in, func(i, j Impulse) bool {
		if asc {
			return i.Duration() < j.Duration()
		}

		return i.Duration() > j.Duration()
	})

	sort.Sort(sorter)
	return in
}

//ByLen sorting of impulses
func (in SelectorImpulses) ByLen(asc bool) SelectorImpulses {
	sorter := newWaverSorter(in, func(i, j Impulse) bool {
		if asc {
			return i.Len() < j.Len()
		}

		return i.Len() > j.Len()
	})

	sort.Sort(sorter)
	return in
}

//ByRetrace sorting of impulses
func (in SelectorImpulses) ByRetrace(val float64, asc bool) SelectorImpulses {
	sorter := newWaverSorter(in, func(i, j Impulse) bool {
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
func (in SelectorImpulses) ByBegins(asc bool) SelectorImpulses {
	sorter := newWaverSorter(in, func(i, j Impulse) bool {
		if asc {
			return j.Begins().After(i.Begins())
		}

		return i.Begins().After(j.Begins())
	})

	sort.Sort(sorter)
	return in
}

//ByEnds sorting of impulses
func (in SelectorImpulses) ByEnds(asc bool) SelectorImpulses {
	sorter := newWaverSorter(in, func(i, j Impulse) bool {
		if asc {
			return j.Ends().After(i.Ends())
		}

		return i.Ends().After(j.Ends())
	})

	sort.Sort(sorter)
	return in
}

//ByStarts sorting of impulses
func (in SelectorImpulses) ByStarts(asc bool) SelectorImpulses {
	sorter := newWaverSorter(in, func(i, j Impulse) bool {
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
func (in SelectorImpulses) ByTops(asc bool) SelectorImpulses {
	sorter := newWaverSorter(in, func(i, j Impulse) bool {
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
func (in SelectorImpulses) BySlope(asc bool) SelectorImpulses {
	sorter := newWaverSorter(in, func(i, j Impulse) bool {
		if asc {
			return i.Slope() < j.Slope()
		}

		return i.Slope() > j.Slope()
	})

	sort.Sort(sorter)
	return in
}
