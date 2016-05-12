package ewa

import (
	"os"

	. "github.com/onsi/ginkgo"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var _ = Describe("Motivewave", func() {

	log.SetHandler(text.New(os.Stdout))
	log.SetLevel(log.DebugLevel)

	path := "/Users/andrewvorobyov/gdrive/MotiveWave" +
		"/OANDA/analysis/CFD/SPX500USD" +
		"/Primary Analysis.mwml"

	Describe("Sorting Correction", func() {

		mw := &mwQuery{}
		_ = mw.importMotiveWaveXML(path)
		markup, _ := mw.parse()
		waves := markup.Waves()

		Describe("Degree", func() {
			It("ASC", func() {
				waves.Corr().ByDegree(true).Print()
			})

			It("DESC", func() {
				waves.Corr().ByDegree(false).Print()
			})
		})

		Describe("Duration", func() {
			It("ASC", func() {
				waves.Corr().ByDuration(true).Print()
			})

			It("DESC", func() {
				waves.Corr().ByDuration(false).Print()
			})
		})

		Describe("Len", func() {
			It("ASC", func() {
				waves.Corr().ByLen(true).Print()
			})

			It("DESC", func() {
				waves.Corr().ByLen(false).Print()
			})
		})

		Describe("Retrace", func() {
			It("ASC", func() {
				waves.Up().Corr().ByRetrace(.5, true).Print()
			})

			It("DESC", func() {
				waves.Up().Corr().ByRetrace(.5, false).Print()
			})
		})

		Describe("Begins", func() {
			It("ASC", func() {
				waves.Corr().ByBegins(true).Print()
			})

			It("DESC", func() {
				waves.Corr().ByBegins(false).Print()
			})
		})

		Describe("Ends", func() {
			It("ASC", func() {
				waves.Corr().ByEnds(true).Print()
			})

			It("DESC", func() {
				waves.Corr().ByEnds(false).Print()
			})
		})

		Describe("Starts", func() {
			It("ASC", func() {
				waves.Up().Corr().ByStarts(true).Print()
			})

			It("DESC", func() {
				waves.Up().Corr().ByStarts(false).Print()
			})
		})

		Describe("Tops", func() {
			It("ASC", func() {
				waves.Down().Corr().ByTops(true).Print()
			})

			It("DESC", func() {
				waves.Down().Corr().ByTops(false).Print()
			})
		})

		Describe("Slope", func() {
			It("ASC", func() {
				waves.Corr().BySlope(true).Print()
			})

			It("DESC", func() {
				waves.Corr().BySlope(false).Print()
			})
		})

	})
})
