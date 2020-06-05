package main

import (
	"go_tsp/data"
	//ga2 "go_tsp/ga"
	aco2 "go_tsp/aco"
)

func main() {
	//ga := ga2.GA{}
	//ga.Data = data.ReadData("data/citys")
	//ga.Run()

	aco := aco2.Aco{}
	aco.Data = data.ReadData("data/citys")
	aco.Run()
}
