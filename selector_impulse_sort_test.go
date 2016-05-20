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

	Describe("Sorting Impulse", func() {

		mw := &mwQuery{}
		_ = mw.importMotiveWaveXML(path)
		markup, _ := mw.parse()

		selector := NewWavesSelector("Tester", markup, &Point{})

		Describe("Degree", func() {
			It("ASC", func() {
				selector.Imp().ByDegree(true).Info()
			})

			It("DESC", func() {
				selector.Imp().ByDegree(false).Info()
			})
		})

		Describe("Duration", func() {
			It("ASC", func() {
				selector.Imp().ByDuration(true).Info()
			})

			It("DESC", func() {
				selector.Imp().ByDuration(false).Info()
			})
		})

		Describe("Len", func() {
			It("ASC", func() {
				selector.Imp().ByLen(true).Info()
			})

			It("DESC", func() {
				selector.Imp().ByLen(false).Info()
			})
		})

		Describe("Retrace", func() {
			It("ASC", func() {
				selector.Up().Imp().ByRetrace(.5, true).Info()
			})

			It("DESC", func() {
				selector.Up().Imp().ByRetrace(.5, false).Info()
			})
		})

		Describe("Begins", func() {
			It("ASC", func() {
				selector.Imp().ByBegins(true).Info()
			})

			It("DESC", func() {
				selector.Imp().ByBegins(false).Info()
			})
		})

		Describe("Ends", func() {
			It("ASC", func() {
				selector.Imp().ByEnds(true).Info()
			})

			It("DESC", func() {
				selector.Imp().ByEnds(false).Info()
			})
		})

		Describe("Starts", func() {
			It("ASC", func() {
				selector.Up().Imp().ByStarts(true).Info()
			})

			It("DESC", func() {
				selector.Up().Imp().ByStarts(false).Info()
			})
		})

		Describe("Tops", func() {
			It("ASC", func() {
				selector.Down().Imp().ByTops(true).Info()
			})

			It("DESC", func() {
				selector.Down().Imp().ByTops(false).Info()
			})
		})

		Describe("Slope", func() {
			It("ASC", func() {
				selector.Imp().BySlope(true).Info()
			})

			It("DESC", func() {
				selector.Imp().BySlope(false).Info()
			})
		})

	})
})
