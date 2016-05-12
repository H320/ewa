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

	})
})
