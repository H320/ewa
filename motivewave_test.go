package ewa

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var _ = Describe("Motivewave", func() {

	log.SetHandler(text.New(os.Stdout))
	log.SetLevel(log.DebugLevel)

	path := "/Users/andrewvorobyov/gdrive/MotiveWave" +
		"/OANDA/analysis/CFD/SPX500USD" +
		"/Primary Analysis.mwml"

	Describe("Internals", func() {

		It("Import", func() {
			mw := &mwQuery{}

			err := mw.importMotiveWaveXML(path)
			Expect(err).Should(Succeed())
		})

		It("Parse", func() {
			mw := &mwQuery{}

			err := mw.importMotiveWaveXML(path)
			Expect(err).Should(Succeed())

			_, err = mw.parse()
			Expect(err).Should(Succeed())
		})
	})

	Describe("Selectors", func() {
		mw := &mwQuery{}
		_ = mw.importMotiveWaveXML(path)
		markup, _ := mw.parse()
		waves := markup.Waves()

		It("Degree", func() {
			waves.Degree(Micro).Print()
		})

		It("Without Sorting", func() {
			waves.Imp().Print()
		})

		It("To", func() {
			waves.To(1869.1).Print()
		})

		It("From", func() {
			waves.From(1944.3).Print()
		})

		Describe("Sorting", func() {
			Describe("Degree", func() {
				It("ASC", func() {
					waves.Imp().ByDegree(true).Print()
				})

				It("DESC", func() {
					waves.Imp().ByDegree(false).Print()
				})
			})

			Describe("Duration", func() {
				It("ASC", func() {
					waves.Imp().ByDuration(true).Print()
				})

				It("DESC", func() {
					waves.Imp().ByDuration(false).Print()
				})
			})

			Describe("Len", func() {
				It("ASC", func() {
					waves.Imp().ByLen(true).Print()
				})

				It("DESC", func() {
					waves.Imp().ByLen(false).Print()
				})
			})

			Describe("Retrace", func() {
				It("ASC", func() {
					waves.Up().Imp().ByRetrace(.5, true).Print()
				})

				It("DESC", func() {
					waves.Up().Imp().ByRetrace(.5, false).Print()
				})
			})

			Describe("Begins", func() {
				It("ASC", func() {
					waves.Imp().ByBegins(true).Print()
				})

				It("DESC", func() {
					waves.Imp().ByBegins(false).Print()
				})
			})

			Describe("Ends", func() {
				It("ASC", func() {
					waves.Imp().ByEnds(true).Print()
				})

				It("DESC", func() {
					waves.Imp().ByEnds(false).Print()
				})
			})

			Describe("Starts", func() {
				It("ASC", func() {
					waves.Up().Imp().ByStarts(true).Print()
				})

				It("DESC", func() {
					waves.Up().Imp().ByStarts(false).Print()
				})
			})

			Describe("Tops", func() {
				It("ASC", func() {
					waves.Down().Imp().ByTops(true).Print()
				})

				It("DESC", func() {
					waves.Down().Imp().ByTops(false).Print()
				})
			})

			Describe("Slope", func() {
				It("ASC", func() {
					waves.Imp().BySlope(true).Print()
				})

				It("DESC", func() {
					waves.Imp().BySlope(false).Print()
				})
			})

		})
	})

})
