package ewa

import (
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/joho/godotenv"
)

var _ = Describe("Motivewave", func() {

	log.SetHandler(text.New(os.Stdout))
	log.SetLevel(log.InfoLevel)

	godotenv.Load()
	path := os.Getenv("PATH_MWML")

	Describe("Selectors", func() {
		log.SetLevel(log.InfoLevel)
		mw := &mwQuery{}
		_ = mw.importMotiveWaveXML(path)
		markup, _ := mw.parse()

		It("Ongoing", func() {

			price := &Point{T: time.Now()}

			helper := NewHelper("My", markup, price)

			degrees := []DegreeType{Minute, Minuette, Subminuette, Micro, Submicro}

			sel := NewWavesSelector("Selector", markup, price)
			sel.Ongoing().Info()

			for _, degree := range degrees {
				log.Infof("In correction %s = %t", degree, helper.InCorrection(degree))
			}

		})
	})
})
