package ewa

//SelectorCorrections struct
type SelectorCorrections struct {
	Label     string
	lastPrice *Point
	Corr      Corrections
}

func (in SelectorCorrections) outCopy() SelectorCorrections {
	return SelectorCorrections{
		Label:     in.Label,
		lastPrice: in.lastPrice,
	}
}

//ByWave finder
func (in SelectorCorrections) ByWave(w Waver) (Correction, bool) {

	for _, one := range in.Corr {
		if one.Starts() == w.Starts() && one.Tops() == w.Tops() {
			return one, true
		}
	}

	return Correction{}, false
}

//Ongoing waves
func (in SelectorCorrections) Ongoing() (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Begins().Before(in.lastPrice.T) && one.Ends().After(in.lastPrice.T) {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//FromTo finds waves that start and end at specified price
func (in SelectorCorrections) FromTo(from, to float64) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Starts() == from && one.Tops() == to {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//From finds waves that start there
func (in SelectorCorrections) From(from float64) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Starts() == from {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//To finds waves that end there
func (in SelectorCorrections) To(to float64) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Tops() == to {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//Sub finds waves that have subdivision
func (in SelectorCorrections) Sub(has bool) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Sub() == has {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

func (in SelectorCorrections) dir(dir bool) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Up() == dir {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//Up finds waves that go up
func (in SelectorCorrections) Up() SelectorCorrections {
	return in.dir(true)
}

//Down finds waves that go down
func (in SelectorCorrections) Down() SelectorCorrections {
	return in.dir(false)
}

//Degree finds waves that end there
func (in SelectorCorrections) Degree(degree DegreeType) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Degree() == degree {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//DegreeGreaterOr finds waves that degree GreaterOr
func (in SelectorCorrections) DegreeGreaterOr(degree DegreeType) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Degree() >= degree {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//DegreeGreater finds waves that degree GreaterOr
func (in SelectorCorrections) DegreeGreater(degree DegreeType) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Degree() > degree {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//DegreeLessOr finds waves that degree GreaterOr
func (in SelectorCorrections) DegreeLessOr(degree DegreeType) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Degree() <= degree {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//DegreeLess finds waves that degree GreaterOr
func (in SelectorCorrections) DegreeLess(degree DegreeType) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Degree() < degree {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//Info selector
func (in SelectorCorrections) Info() {
	for _, one := range in.Corr {
		one.Info(in.Label)
	}
}

//Len how many waves
func (in SelectorCorrections) Len() int {
	return len(in.Corr)
}

//First gets first correction
func (in SelectorCorrections) First() (Correction, bool) {
	if len(in.Corr) > 0 {
		return in.Corr[0], true
	}

	return Correction{}, false
}

//Last gets last correction
func (in SelectorCorrections) Last() (Correction, bool) {
	if len(in.Corr) > 0 {
		return in.Corr[len(in.Corr)-1], true
	}

	return Correction{}, false
}

// Specific to corrections

//Type gets corrections by type
func (in SelectorCorrections) Type(ct CorrectionType) (out SelectorCorrections) {
	out = in.outCopy()

	for _, one := range in.Corr {
		if one.Type() == ct {
			out.Corr = append(out.Corr, one)
		}
	}

	return
}

//Zigzag corrections only
func (in SelectorCorrections) Zigzag() (out SelectorCorrections) {
	return in.Type(CTZigzag)
}

//Flat corrections only
func (in SelectorCorrections) Flat() (out SelectorCorrections) {
	return in.Type(CTFlat)
}

//Triangle corrections only
func (in SelectorCorrections) Triangle() (out SelectorCorrections) {
	return in.Type(CTTriangle)
}

//Combo corrections only
func (in SelectorCorrections) Combo() (out SelectorCorrections) {
	return in.Type(CTCombo)
}

//Triple corrections only
func (in SelectorCorrections) Triple() (out SelectorCorrections) {
	return in.Type(CTTriple)
}
