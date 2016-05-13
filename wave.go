package ewa

import "github.com/apex/log"

//Info corrections selector
func (w Wave) Info(label string) {
	log.WithFields(log.Fields{
		"M": w.Move,
		"D": w.Degree(),
	}).Info(label)
}

//Degree gets degree
func (w Wave) Degree() DegreeType {
	return w.WaveDegree
}

//ParentWave getter
func (w Wave) ParentWave() Waver {
	return w.Parent
}

//NextWave getter
func (w Wave) NextWave() Waver {
	return w.Next
}

//PrevWave getter
func (w Wave) PrevWave() Waver {
	return w.Prev
}
