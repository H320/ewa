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
	if w.Parent != nil {
		return w.Parent
	}
	return nil
}

//NextWave getter
func (w Wave) NextWave() Waver {
	if w.Next != nil {
		return w.Next
	}
	return nil
}

//PrevWave getter
func (w Wave) PrevWave() Waver {
	if w.Prev != nil {
		return w.Prev
	}
	return nil
}
