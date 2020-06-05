package main

import (
	"go_tsp/data"
	ga2 "go_tsp/ga"
)

func main() {
	ga := ga2.GA{}
	ga.Data = data.ReadData("data/citys")
	ga.CodeLen = len(ga.Data)
	ga.Run()
}
