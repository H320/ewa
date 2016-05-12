package ewa

//DegreeType - type of wave degrees
type DegreeType uint

// Degrees of the wave
const (
	Pico DegreeType = 1 << iota
	Subnano
	Nano
	Miniscule
	Submicro
	Micro
	Subminuette
	Minuette
	Minute
	Minor
	Intermediate
	Primary
)

//CorrectionType type
type CorrectionType string

// Correction types
const (
	CTUnknown  CorrectionType = "unknown"
	CTZigzag                  = "zigzag"
	CTFlat                    = "flat"
	CTTriangle                = "triangle"
	CTCombo                   = "combo"
	CTTriple                  = "triple"
)

//Degree - gets degree type from string
func Degree(degree string) DegreeType {
	switch degree {
	case "PRIMARY":
		return Primary
	case "INTERMEDIATE":
		return Intermediate
	case "MINOR":
		return Minor
	case "MINUTE":
		return Minute
	case "MINUETTE":
		return Minuette
	case "SUBMINUETTE":
		return Subminuette
	case "MICRO":
		return Micro
	case "SUBMICRO":
		return Submicro
	case "MINISCULE":
		return Miniscule
	case "NANO":
		return Nano
	case "SUBNANO":
		return Subnano
	case "PICO":
		return Pico
	}

	return Minuette
}

//Impulses slice
type Impulses []Impulse

//Corrections slice
type Corrections []Correction

//Waves - base type for selector
type Waves struct {
	Impulses    []Impulse
	Corrections []Correction
}
