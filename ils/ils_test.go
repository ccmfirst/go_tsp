package ils

import (
	"testing"

	"go_tsp/data"
)

func TestILS_Run(t *testing.T) {
	ils := &ILS{
		MaxIter:      600,
		MaxNoImprove: 50,
		Data:         data.ReadData("../data/cities"),
	}

	ils.CitySize = len(ils.Data)
	ils.Run()
}
