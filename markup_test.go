package ewa

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var _ = Describe("Markup", func() {

	log.SetHandler(text.New(os.Stdout))
	log.SetLevel(log.DebugLevel)

	path := "/Users/andrewvorobyov/gdrive/MotiveWave" +
		"/OANDA/analysis/CFD/SPX500USD" +
		"/Primary Analysis.mwml"

	Describe("Tree", func() {

		It("Dev", func() {
			mw := &mwQuery{}

			err := mw.importMotiveWaveXML(path)
			Expect(err).Should(Succeed())

			markup := &Markup{}

			markup.processImpulses(mw)
			markup.processCorrections(mw)
			markup.processTriangles(mw)
			markup.processTripleCombo(mw)
			markup.processTree(mw)
			markup.printStackTree()
		})
	})
})
