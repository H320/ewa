package ewa

//SelectorImpulses struct
type SelectorImpulses struct {
	Label     string
	lastPrice *Point
	Imp       Impulses
}

func (in SelectorImpulses) outCopy() SelectorImpulses {
	return SelectorImpulses{
		Label:     in.Label,
		lastPrice: in.lastPrice,
	}
}

//ByWave finder
func (in SelectorImpulses) ByWave(w Waver) (Impulse, bool) {

	for _, one := range in.Imp {
		if one.Starts() == w.Starts() && one.Tops() == w.Tops() {
			return one, true
		}
	}

	return Impulse{}, false
}

//Ongoing waves
func (in SelectorImpulses) Ongoing() (out SelectorImpulses) {
	out = in.outCopy()

	for _, one := range in.Imp {
		if one.Begins().Before(in.lastPrice.T) && one.Ends().After(in.lastPrice.T) {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//FromTo finds waves that start and end at specified price
func (in SelectorImpulses) FromTo(from, to float64) (out SelectorImpulses) {
	out = in.outCopy()

	for _, one := range in.Imp {
		if one.Starts() == from && one.Tops() == to {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//From finds waves that start there
func (in SelectorImpulses) From(from float64) (out SelectorImpulses) {
	out = in.outCopy()

	for _, one := range in.Imp {
		if one.Starts() == from {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//To finds waves that end there
func (in SelectorImpulses) To(to float64) (out SelectorImpulses) {
	out = in.outCopy()

	for _, one := range in.Imp {
		if one.Tops() == to {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//Sub finds waves that have subdivision
func (in SelectorImpulses) Sub(has bool) (out SelectorImpulses) {
	out = in.outCopy()

	for _, one := range in.Imp {
		if one.Sub() == has {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

func (in SelectorImpulses) dir(dir bool) (out SelectorImpulses) {
	out = in.outCopy()

	for _, one := range in.Imp {
		if one.Up() == dir {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//Up finds waves that go up
func (in SelectorImpulses) Up() SelectorImpulses {
	return in.dir(true)
}

//Down finds waves that go down
func (in SelectorImpulses) Down() SelectorImpulses {
	return in.dir(false)
}

//Degree finds waves that end there
func (in SelectorImpulses) Degree(degree DegreeType) (out SelectorImpulses) {
	out = in.outCopy()

	for _, one := range in.Imp {
		if one.Degree() == degree {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//DegreeGreaterOr finds waves that degree GreaterOr
func (in SelectorImpulses) DegreeGreaterOr(degree DegreeType) (out SelectorImpulses) {
	out = in.outCopy()

	for _, one := range in.Imp {
		if one.Degree() >= degree {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//DegreeGreater finds waves that degree GreaterOr
func (in SelectorImpulses) DegreeGreater(degree DegreeType) (out SelectorImpulses) {
	out = in.outCopy()
	for _, one := range in.Imp {
		if one.Degree() > degree {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//DegreeLessOr finds waves that degree GreaterOr
func (in SelectorImpulses) DegreeLessOr(degree DegreeType) (out SelectorImpulses) {
	out = in.outCopy()

	for _, one := range in.Imp {
		if one.Degree() <= degree {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//DegreeLess finds waves that degree GreaterOr
func (in SelectorImpulses) DegreeLess(degree DegreeType) (out SelectorImpulses) {
	out = in.outCopy()

	for _, one := range in.Imp {
		if one.Degree() < degree {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//Info selector
func (in SelectorImpulses) Info() {
	for _, one := range in.Imp {
		one.Info(in.Label)
	}
}

//Len how many waves
func (in SelectorImpulses) Len() int {
	return len(in.Imp)
}

//First gets first impulse
func (in SelectorImpulses) First() (Impulse, bool) {
	if len(in.Imp) > 0 {
		return in.Imp[0], true
	}

	return Impulse{}, false
}

//Last gets last impulse
func (in SelectorImpulses) Last() (Impulse, bool) {
	if len(in.Imp) > 0 {
		return in.Imp[len(in.Imp)-1], true
	}

	return Impulse{}, false
}

// Specific to impulses

//Extended gets only extended impulses
func (in SelectorImpulses) Extended(extended bool) (out SelectorImpulses) {
	for _, one := range in.Imp {
		if one.Extended() == extended {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}

//Diagonal gets only diagonal impulses
func (in SelectorImpulses) Diagonal(diagonal bool) (out SelectorImpulses) {
	for _, one := range in.Imp {
		if one.Diagonal() == diagonal {
			out.Imp = append(out.Imp, one)
		}
	}

	return
}
