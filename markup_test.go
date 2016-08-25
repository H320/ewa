package ewa

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/joho/godotenv"
)

var _ = Describe("Markup", func() {

	log.SetHandler(text.New(os.Stdout))
	log.SetLevel(log.DebugLevel)

	godotenv.Load()
	path := os.Getenv("PATH_MWML")

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
