package ewa

//Helper - helper for advanced EWA functions
type Helper struct {
	markup    Markup
	lastPrice Point
}

//NewHelper creates new helper
func NewHelper(m *Markup, p Point) Helper {
	return Helper{*m, p}
}

//InCorrection returns if there is ongoing correction of this degree
func (h Helper) InCorrection(degree DegreeType) bool {
	return h.markup.Waves().Degree(degree).Ongoing(h.lastPrice).Corr().Len() > 0
}
