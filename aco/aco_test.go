package aco

import (
	"go_tsp/data"
	"testing"
)

func TestAco_Run(t *testing.T) {
	s := &Aco{}
	s.Data = data.ReadData("../data/cities")
	s.Run()
}
