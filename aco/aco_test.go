package aco

import (
	"testing"

	"go_tsp/data"
)

func TestAco_Run(t *testing.T) {
	s := &Aco{}
	s.Data = data.ReadData("../data/cities")
	s.Run()
}
