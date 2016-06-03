package ewa

import (
	"fmt"

	"github.com/apex/log"
)

//Markup struct
type Markup struct {
	Impulses    []*Impulse
	Corrections []*Correction
	Zigzags     []*Zigzag
	Flats       []*Flat
	Triangles   []*Triangle
	Combos      []*Combo
	Triples     []*Triple
}

func (m *Markup) addImpulse(impulse *Impulse) *Impulse {

	context := log.WithFields(log.Fields{
		"M": impulse.Move,
		"D": impulse.Degree,
	})

	for _, one := range m.Impulses {
		if one.Move.Base.P == impulse.Move.Base.P &&
			one.Move.End.P == impulse.Move.End.P {

			context.Debug("Impulse found")
			return one
		}
	}

	context.Debug("+Impulse")

	m.Impulses = append(m.Impulses, impulse)
	return impulse
}

func (m *Markup) addCorrection(correction *Correction) *Correction {
	context := log.WithFields(log.Fields{
		"M": correction.Move,
		"D": correction.Degree,
	})

	for _, one := range m.Corrections {
		if one.Move.Base.P == correction.Move.Base.P &&
			one.Move.End.P == correction.Move.End.P {

			context.Debug("Correction found")
			return one
		}
	}

	context.Debug("+Correction")

	m.Corrections = append(m.Corrections, correction)
	return correction
}

func (m *Markup) processImpulseWave(w mwImpulse) {

	degree := Degree(w.Degree)
	lessDegree := degree >> 1

	//Generating points

	ori := PointFromMW(w.Origin)
	pW1 := PointFromMW(w.Wave1)
	pW2 := PointFromMW(w.Wave2)
	pW3 := PointFromMW(w.Wave3)
	pW4 := PointFromMW(w.Wave4)
	pW5 := PointFromMW(w.Wave5)

	//Generating sub waves

	subW1 := &Wave{Move: &Move{ori, pW1}, WaveDegree: lessDegree}
	subW2 := &Wave{Move: &Move{pW1, pW2}, WaveDegree: lessDegree}
	subW3 := &Wave{Move: &Move{pW2, pW3}, WaveDegree: lessDegree}
	subW4 := &Wave{Move: &Move{pW3, pW4}, WaveDegree: lessDegree}
	subW5 := &Wave{Move: &Move{pW4, pW5}, WaveDegree: lessDegree}

	setImpulseWaveAdjecency([5]*Wave{subW1, subW2, subW3, subW4, subW5})

	// Generating impulse wave
	wave := &Wave{Move: &Move{ori, pW5}, WaveDegree: degree}

	setParentWave(wave, subW1, subW2, subW3, subW4, subW5)

	// Adding primitives
	impulse := m.addImpulse(&Impulse{Wave: wave})

	impulse.W1 = m.addImpulse(&Impulse{Wave: subW1})
	impulse.W3 = m.addImpulse(&Impulse{Wave: subW3})
	impulse.W5 = m.addImpulse(&Impulse{Wave: subW5})

	impulse.W2 = m.addCorrection(&Correction{Wave: subW2})
	impulse.W4 = m.addCorrection(&Correction{Wave: subW4})
}

func (m *Markup) processImpulses(mwQuery *mwQuery) {
	for _, w := range mwQuery.Impulses {
		m.processImpulseWave(w)
	}

	for _, w := range mwQuery.ImpulsesLeading {
		m.processImpulseWave(w)
	}

	for _, w := range mwQuery.ImpulsesEnding {
		m.processImpulseWave(w)
	}
}

func (m *Markup) processCorrections(mwQuery *mwQuery) {

	for _, w := range mwQuery.Corrections {

		degree := Degree(w.Degree)
		lessDegree := degree >> 1

		//Generating points

		ori := PointFromMW(w.Origin)
		pWA := PointFromMW(w.WaveA)
		pWB := PointFromMW(w.WaveB)
		pWC := PointFromMW(w.WaveC)

		//Generating sub waves

		wave := &Wave{Move: &Move{ori, pWC}, WaveDegree: degree}

		subWA := &Wave{Move: &Move{ori, pWA}, WaveDegree: lessDegree}
		subWB := &Wave{Move: &Move{pWA, pWB}, WaveDegree: lessDegree}
		subWC := &Wave{Move: &Move{pWB, pWC}, WaveDegree: lessDegree}

		setParentWave(wave, subWA, subWB, subWC)

		corr := m.addCorrection(&Correction{Wave: wave})

		wB := m.addCorrection(&Correction{Wave: subWB})
		wC := m.addImpulse(&Impulse{Wave: subWC})

		if subWA.Len()*.9 <= subWB.Len() {
			// Flat

			wA := m.addCorrection(&Correction{Wave: subWA})
			flat := &Flat{A: wA, B: wB, C: wC}
			corr.Flat = flat
			m.Flats = append(m.Flats, flat)
		} else {
			// Zigzag

			wA := m.addImpulse(&Impulse{Wave: subWA})
			zigzag := &Zigzag{A: wA, B: wB, C: wC}
			corr.Zigzag = zigzag
			m.Zigzags = append(m.Zigzags, zigzag)
		}
	}
}

func (m *Markup) processTriangles(mwQuery *mwQuery) {

	for _, w := range mwQuery.Triangles {

		degree := Degree(w.Degree)
		lessDegree := degree >> 1

		//Generating points

		ori := PointFromMW(w.Origin)
		pWA := PointFromMW(w.WaveA)
		pWB := PointFromMW(w.WaveB)
		pWC := PointFromMW(w.WaveC)
		pWD := PointFromMW(w.WaveD)
		pWE := PointFromMW(w.WaveE)

		//Generating sub waves

		wave := &Wave{Move: &Move{ori, pWE}, WaveDegree: degree}

		subWA := &Wave{Move: &Move{ori, pWA}, WaveDegree: lessDegree}
		subWB := &Wave{Move: &Move{pWA, pWB}, WaveDegree: lessDegree}
		subWC := &Wave{Move: &Move{pWB, pWC}, WaveDegree: lessDegree}
		subWD := &Wave{Move: &Move{pWC, pWD}, WaveDegree: lessDegree}
		subWE := &Wave{Move: &Move{pWD, pWE}, WaveDegree: lessDegree}

		setParentWave(wave, subWA, subWB, subWC, subWD, subWE)

		// Generating triangle wave
		triangleWave := &Wave{Move: &Move{ori, pWE}, WaveDegree: degree}

		triangleCorrection := m.addCorrection(&Correction{Wave: triangleWave})

		triangle := &Triangle{
			A: m.addCorrection(&Correction{Wave: subWA}),
			B: m.addCorrection(&Correction{Wave: subWB}),
			C: m.addCorrection(&Correction{Wave: subWC}),
			D: m.addCorrection(&Correction{Wave: subWD}),
			E: m.addCorrection(&Correction{Wave: subWE}),
		}

		log.WithField("Triangle", triangle).Debug("+Triangle")

		triangleCorrection.Triangle = triangle

		m.Triangles = append(m.Triangles, triangle)

		//TODO 9 wave triangle
	}
}

func (m *Markup) processTripleCombo(mwQuery *mwQuery) {
	for _, w := range mwQuery.TripleCombo {

		degree := Degree(w.Degree)
		lessDegree := degree >> 1

		//Generating points

		ori := PointFromMW(w.Origin)
		pWW := PointFromMW(w.WaveW)
		pWX := PointFromMW(w.WaveX)
		pWY := PointFromMW(w.WaveY)
		pWX2 := PointFromMW(w.WaveX2)
		pWZ := PointFromMW(w.WaveZ)

		//Generating sub waves

		wave := &Wave{Move: &Move{ori, pWZ}, WaveDegree: degree}

		subWW := &Wave{Move: &Move{ori, pWW}, WaveDegree: lessDegree}
		subWX := &Wave{Move: &Move{pWW, pWX}, WaveDegree: lessDegree}
		subWY := &Wave{Move: &Move{pWX, pWY}, WaveDegree: lessDegree}
		subWX2 := &Wave{Move: &Move{pWY, pWX2}, WaveDegree: lessDegree}
		subWZ := &Wave{Move: &Move{pWX2, pWZ}, WaveDegree: lessDegree}

		setParentWave(wave, subWW, subWX, subWY, subWX2, subWZ)

		// Generating triangle wave
		tripleWave := &Wave{Move: &Move{ori, pWZ}, WaveDegree: degree}

		tripleCorrection := m.addCorrection(&Correction{Wave: tripleWave})

		triple := &Triple{
			W:  m.addCorrection(&Correction{Wave: subWW}),
			X:  m.addCorrection(&Correction{Wave: subWX}),
			Y:  m.addCorrection(&Correction{Wave: subWY}),
			X2: m.addCorrection(&Correction{Wave: subWX2}),
			Z:  m.addCorrection(&Correction{Wave: subWZ}),
		}

		tripleCorrection.Triple = triple

		log.WithField("Triple", triple).Debug("+Triple")

		m.Triples = append(m.Triples, triple)
	}
}

func (m *Markup) processCombo(mwQuery *mwQuery) {

	for _, w := range mwQuery.Combo {
		degree := Degree(w.Degree)
		lessDegree := degree >> 1

		//Generating points

		ori := PointFromMW(w.Origin)
		pWW := PointFromMW(w.WaveW)
		pWX := PointFromMW(w.WaveX)
		pWY := PointFromMW(w.WaveY)

		//Generating sub waves

		wave := &Wave{Move: &Move{ori, pWY}, WaveDegree: degree}

		subWW := &Wave{Move: &Move{ori, pWW}, WaveDegree: lessDegree}
		subWX := &Wave{Move: &Move{pWW, pWX}, WaveDegree: lessDegree}
		subWY := &Wave{Move: &Move{pWX, pWY}, WaveDegree: lessDegree}

		setParentWave(wave, subWW, subWX, subWY)

		// Generating triangle wave
		comboWave := &Wave{Move: &Move{ori, pWY}, WaveDegree: degree}

		comboCorrection := m.addCorrection(&Correction{Wave: comboWave})

		combo := &Combo{
			W: m.addCorrection(&Correction{Wave: subWW}),
			X: m.addCorrection(&Correction{Wave: subWX}),
			Y: m.addCorrection(&Correction{Wave: subWY}),
		}

		comboCorrection.Combo = combo

		log.WithField("Combo", combo).Debug("+Combo")

		m.Combos = append(m.Combos, combo)
	}
}

func (m *Markup) processTree(mwQuery *mwQuery) {
	log.Debug("Processing Tree")

	setNext := func(curr *Wave) {
		for _, n := range m.Impulses {
			if curr.Ends() == n.Begins() && curr.Degree() == n.Degree() {
				curr.Next = n.Wave
				n.Prev = curr
				return
			}
		}

		for _, n := range m.Corrections {
			if curr.Ends() == n.Begins() && curr.Degree() == n.Degree() {
				curr.Next = n.Wave
				n.Prev = curr
				return
			}
		}
	}

	for _, i := range m.Impulses {
		if i.Next == nil {
			setNext(i.Wave)
		}
	}

	for _, i := range m.Corrections {
		if i.Next == nil {
			setNext(i.Wave)
		}
	}
}

func (m *Markup) printStackTree() {

	for _, one := range m.Impulses {
		log.WithFields(log.Fields{
			"Time": fmt.Sprintf("%s->%s", one.Base.T, one.End.T),
			"Move": fmt.Sprintf("%.2f->%.2f", one.Base.P, one.End.P),
			"Prev": one.PrevWave(),
			"Next": one.NextWave(),
		}).Debug("Imp")
	}

	for _, one := range m.Corrections {
		log.WithFields(log.Fields{
			"Time": fmt.Sprintf("%s->%s", one.Base.T, one.End.T),
			"Move": fmt.Sprintf("%.2f->%.2f", one.Base.P, one.End.P),
			"Prev": one.PrevWave(),
			"Next": one.NextWave(),
		}).Debug("Cor")
	}
}

func (m *Markup) printStack() {

	for _, one := range m.Impulses {
		log.WithFields(log.Fields{
			"D":  one.Duration(),
			"P":  fmt.Sprintf("%.2f->%.2f", one.Base.P, one.End.P),
			"W1": one.W1,
			"W2": one.W2,
			"W3": one.W3,
			"W4": one.W4,
			"W5": one.W5,
		}).Debug("Imp")
	}

	for _, one := range m.Corrections {
		log.WithFields(log.Fields{
			"D": one.Duration(),
			"P": fmt.Sprintf("%.2f->%.2f", one.Base.P, one.End.P),
			"T": one.Type(),
		}).Debug("Corr")
	}
}
