package data

import (
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type city struct {
	lat int
	lng int
}

func ReadData(filename string) [][]float64 {
	var cities []city
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cs, err := ioutil.ReadAll(f)
	citys := string(cs)
	a := strings.Split(citys, "\n")
	for _, v := range a {
		var ct city
		b := strings.Split(v, "\t")
		ct.lat, _ = strconv.Atoi(b[0])
		ct.lng, _ = strconv.Atoi(b[1])
		cities = append(cities, ct)
	}
	data := make([][]float64, len(cities))
	for i := 0; i < len(cities); i++ {
		data[i] = make([]float64, len(cities))
		for j := 0; j < i; j++ {
			data[i][j] = math.Sqrt(math.Pow(float64(cities[i].lat-cities[j].lat), 2) + math.Pow(float64(cities[i].lng-cities[j].lng), 2))
			data[j][i] = data[i][j]
		}
	}
	return data
}
