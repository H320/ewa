package ewa

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/joho/godotenv"
)

var _ = Describe("Motivewave", func() {

	log.SetHandler(text.New(os.Stdout))
	log.SetLevel(log.DebugLevel)

	godotenv.Load()
	path := os.Getenv("PATH_MWML")

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
})
