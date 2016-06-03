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

	Describe("Sorting Corection", func() {

		mw := &mwQuery{}
		_ = mw.importMotiveWaveXML(path)
		markup, _ := mw.parse()

		selector := NewWavesSelector("Tester", markup, &Point{})

		Describe("Degree", func() {
			It("ASC", func() {
				selector.Cor().ByDegree(true).Info()
			})

			It("DESC", func() {
				selector.Cor().ByDegree(false).Info()
			})
		})

		Describe("Duration", func() {
			It("ASC", func() {
				selector.Cor().ByDuration(true).Info()
			})

			It("DESC", func() {
				selector.Cor().ByDuration(false).Info()
			})
		})

		Describe("Len", func() {
			It("ASC", func() {
				selector.Cor().ByLen(true).Info()
			})

			It("DESC", func() {
				selector.Cor().ByLen(false).Info()
			})
		})

		Describe("Retrace", func() {
			It("ASC", func() {
				selector.Up().Cor().ByRetrace(.5, true).Info()
			})

			It("DESC", func() {
				selector.Up().Cor().ByRetrace(.5, false).Info()
			})
		})

		Describe("Begins", func() {
			It("ASC", func() {
				selector.Cor().ByBegins(true).Info()
			})

			It("DESC", func() {
				selector.Cor().ByBegins(false).Info()
			})
		})

		Describe("Ends", func() {
			It("ASC", func() {
				selector.Cor().ByEnds(true).Info()
			})

			It("DESC", func() {
				selector.Cor().ByEnds(false).Info()
			})
		})

		Describe("Starts", func() {
			It("ASC", func() {
				selector.Up().Cor().ByStarts(true).Info()
			})

			It("DESC", func() {
				selector.Up().Cor().ByStarts(false).Info()
			})
		})

		Describe("Tops", func() {
			It("ASC", func() {
				selector.Down().Cor().ByTops(true).Info()
			})

			It("DESC", func() {
				selector.Down().Cor().ByTops(false).Info()
			})
		})

		Describe("Slope", func() {
			It("ASC", func() {
				selector.Cor().BySlope(true).Info()
			})

			It("DESC", func() {
				selector.Cor().BySlope(false).Info()
			})
		})

	})
})
