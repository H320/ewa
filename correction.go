package ewa

import (
	"fmt"

	"github.com/apex/log"
)

//Info prints impulse with label
func (c Correction) Info(label string) {
	log.WithFields(log.Fields{
		"M": c.Move,
		"D": fmt.Sprintf("%3d", c.Degree()),
		"T": c.Type(),
	}).Info(label)
}

//Print impulse
func (c Correction) Print() {
	c.Info("Corr")
}

//Sub - does it has subwaves
func (c Correction) Sub() bool {
	return c.Zigzag != nil || c.Flat != nil ||
		c.Triangle != nil || c.Combo != nil || c.Triple != nil
}

//Type - gets correction type
func (c Correction) Type() CorrectionType {
	if c.Zigzag != nil {
		return CTZigzag
	}

	if c.Flat != nil {
		return CTFlat
	}

	if c.Triangle != nil {
		return CTTriangle
	}

	if c.Combo != nil {
		return CTCombo
	}

	if c.Triple != nil {
		return CTTriple
	}

	return CTUnknown
}
