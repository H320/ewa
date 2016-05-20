package ewa

//Selector struc
type Selector struct {
	Label     string
	lastPrice *Point
}

//SelectorWaves struct
type SelectorWaves struct {
	Label       string
	lastPrice   *Point
	Impulses    []Impulse
	Corrections []Correction
}

//NewWavesSelector creates new selector
func NewWavesSelector(label string, m *Markup, lastPrice *Point) SelectorWaves {
	s := SelectorWaves{
		Label:     label,
		lastPrice: lastPrice,
	}

	for _, i := range m.Impulses {
		s.Impulses = append(s.Impulses, *i)
	}

	for _, c := range m.Corrections {
		s.Corrections = append(s.Corrections, *c)
	}

	return s
}

func (in SelectorWaves) outCopy() SelectorWaves {
	return SelectorWaves{
		Label:     in.Label,
		lastPrice: in.lastPrice,
	}
}

//ByWave finder
func (in SelectorWaves) ByWave(w Waver) (Impulse, Correction, bool) {

	for _, one := range in.Corrections {
		if one.Starts() == w.Starts() && one.Tops() == w.Tops() {
			return Impulse{}, one, true
		}
	}

	for _, one := range in.Impulses {
		if one.Starts() == w.Starts() && one.Tops() == w.Tops() {
			return one, Correction{}, true
		}
	}

	return Impulse{}, Correction{}, false
}

//Ongoing waves
func (in SelectorWaves) Ongoing() (out SelectorWaves) {
	out = in.outCopy()

	for _, one := range in.Impulses {
		if one.Begins().Before(in.lastPrice.T) && one.Ends().After(in.lastPrice.T) {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Begins().Before(in.lastPrice.T) && one.Ends().After(in.lastPrice.T) {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//FromTo finds waves that start and end at specified price
func (in SelectorWaves) FromTo(from, to float64) (out SelectorWaves) {
	out = in.outCopy()

	for _, one := range in.Impulses {
		if one.Starts() == from && one.Tops() == to {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Starts() == from && one.Tops() == to {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//From finds waves that start there
func (in SelectorWaves) From(from float64) (out SelectorWaves) {
	out = in.outCopy()

	for _, one := range in.Impulses {
		if one.Starts() == from {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Starts() == from {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//To finds waves that end there
func (in SelectorWaves) To(to float64) (out SelectorWaves) {
	out = in.outCopy()

	for _, one := range in.Impulses {
		if one.Tops() == to {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Tops() == to {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//Sub finds waves that have subdivision
func (in SelectorWaves) Sub(has bool) (out SelectorWaves) {
	out = in.outCopy()

	for _, one := range in.Impulses {
		if one.Sub() == has {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Sub() == has {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

func (in SelectorWaves) dir(dir bool) (out SelectorWaves) {
	out = in.outCopy()

	for _, one := range in.Impulses {
		if one.Up() == dir {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Up() == dir {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//Up finds waves that go up
func (in SelectorWaves) Up() SelectorWaves {
	return in.dir(true)
}

//Down finds waves that go down
func (in SelectorWaves) Down() SelectorWaves {
	return in.dir(false)
}

//Degree finds waves that end there
func (in SelectorWaves) Degree(degree DegreeType) (out SelectorWaves) {
	out = in.outCopy()

	for _, one := range in.Impulses {
		if one.Degree() == degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree() == degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//DegreeGreaterOr finds waves that degree GreaterOr
func (in SelectorWaves) DegreeGreaterOr(degree DegreeType) (out SelectorWaves) {
	out = in.outCopy()

	for _, one := range in.Impulses {
		if one.Degree() >= degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree() >= degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//DegreeGreater finds waves that degree GreaterOr
func (in SelectorWaves) DegreeGreater(degree DegreeType) (out SelectorWaves) {
	out = in.outCopy()
	for _, one := range in.Impulses {
		if one.Degree() > degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree() > degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//DegreeLessOr finds waves that degree GreaterOr
func (in SelectorWaves) DegreeLessOr(degree DegreeType) (out SelectorWaves) {
	out = in.outCopy()

	for _, one := range in.Impulses {
		if one.Degree() <= degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree() <= degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//DegreeLess finds waves that degree GreaterOr
func (in SelectorWaves) DegreeLess(degree DegreeType) (out SelectorWaves) {
	out = in.outCopy()

	for _, one := range in.Impulses {
		if one.Degree() < degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree() < degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//Info selector
func (in SelectorWaves) Info() {
	for _, one := range in.Impulses {
		one.Info(in.Label)
	}

	for _, one := range in.Corrections {
		one.Info(in.Label)
	}
}

//Len how many waves
func (in SelectorWaves) Len() int {
	return len(in.Corrections) + len(in.Impulses)
}

//Imp - getting impulses selector
func (in SelectorWaves) Imp() SelectorImpulses {
	return SelectorImpulses{
		lastPrice: in.lastPrice,
		Label:     in.Label,
		Imp:       in.Impulses,
	}
}

//Cor - getting impulses selector
func (in SelectorWaves) Cor() SelectorCorrections {
	return SelectorCorrections{
		lastPrice: in.lastPrice,
		Label:     in.Label,
		Corr:      in.Corrections,
	}
}
