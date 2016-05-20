package ewa

import (
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var _ = Describe("Motivewave", func() {

	log.SetHandler(text.New(os.Stdout))
	log.SetLevel(log.InfoLevel)

	path := "/Users/andrewvorobyov/gdrive/MotiveWave" +
		"/OANDA/analysis/CFD/SPX500USD" +
		"/Primary Analysis.mwml"

	FDescribe("Selectors", func() {
		log.SetLevel(log.InfoLevel)
		mw := &mwQuery{}
		_ = mw.importMotiveWaveXML(path)
		markup, _ := mw.parse()
		waves := markup.Waves()

		It("Ongoing", func() {

			price := Point{T: time.Now()}

			helper := NewHelper(markup, price)

			degrees := []DegreeType{Minute, Minuette, Subminuette, Micro, Submicro}

			waves.Ongoing(price).Info("Ongoing")

			for _, degree := range degrees {
				log.Infof("In correction %s = %t", degree, helper.InCorrection(degree))
			}

		})
	})
})
