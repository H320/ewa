package ewa

//Helper - helper for advanced EWA functions
type Helper struct {
	sel       SelectorWaves
	lastPrice *Point
}

//NewHelper creates new helper
func NewHelper(label string, m *Markup, lastPrice *Point) Helper {
	return Helper{
		sel:       NewWavesSelector(label, m, lastPrice),
		lastPrice: lastPrice,
	}
}

//InCorrection returns if there is ongoing correction of this degree
func (h Helper) InCorrection(degree DegreeType) bool {
	return h.sel.Degree(degree).Ongoing().Cor().Len() > 0
}
