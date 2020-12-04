package vns

import (
	"go_tsp/data"
	"testing"
)

func TestVns_Run(t *testing.T) {
	s := &Vns{
		Gen: 50,
	}
	s.Data = data.ReadData("../data/cities")
	s.Run()
}
