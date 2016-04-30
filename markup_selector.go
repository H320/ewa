package ewa

import "github.com/apex/log"

//Waves makes waves structure
func (m *Markup) Waves() Waves {
	return Waves{Impulses: m.Impulses, Corrections: m.Corrections}
}

//FromTo finds waves that start and end at specified price
func (in Waves) FromTo(from, to float64) (out Waves) {
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
func (in Waves) From(from float64) (out Waves) {

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
func (in Waves) To(to float64) (out Waves) {

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

//Degree finds waves that end there
func (in Waves) Degree(degree DegreeType) (out Waves) {

	for _, one := range in.Impulses {
		if one.Degree == degree {
			out.Impulses = append(out.Impulses, one)
		}
	}

	for _, one := range in.Corrections {
		if one.Degree == degree {
			out.Corrections = append(out.Corrections, one)
		}
	}

	return
}

//Print selector
func (in Waves) Print() {
	for _, one := range in.Impulses {
		log.WithFields(log.Fields{
			"M": one.Move,
			"D": one.Degree,
		}).Info("Impulse")
	}

	for _, one := range in.Corrections {
		log.WithFields(log.Fields{
			"M": one.Move,
			"D": one.Degree,
			"T": one.Type(),
		}).Info("Correction")
	}

	return
}
