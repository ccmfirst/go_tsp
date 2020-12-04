package ga

import (
	"testing"

	"go_tsp/data"
)

func TestGA_Run(t *testing.T) {
	s := &GA{}
	s.Data = data.ReadData("../data/cities")
	s.Run()
}
