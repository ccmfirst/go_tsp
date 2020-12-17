package vns

import (
	"fmt"
	"math/rand"
)

type Vns struct {
	Gen     int // 邻域搜索次数
	CityNum int
	Solve   []int   // 最优解
	Value   float64 // 最优值
	Data    [][]float64
}

func (s *Vns) Run() {
	s.CityNum = len(s.Data)
	s.Solve = rand.Perm(s.CityNum)
	s.Value = s.CalculateDis(s.Solve)

	i := 0

	for i < 2 {
	Here:
		tempSolve := make([]int, s.CityNum)
		copy(tempSolve, s.Solve)
		for j := 0; j < s.Gen; j++ {
			if i == 0 {
				path := s.Reverse(tempSolve)
				//fmt.Println(tempSolve, "\n", path)
				tempValue := s.CalculateDis(path)

				if tempValue < s.Value {
					fmt.Println(i)
					s.Value = tempValue
					s.Solve = path
					fmt.Println(s.Value)
					goto Here
				}
			}

			if i == 1 {
				path := s.Insert(tempSolve)
				tempValue := s.CalculateDis(path)
				if tempValue < s.Value {
					fmt.Println(i)
					s.Value = tempValue
					s.Solve = path
					fmt.Println(s.Value)
					goto Here
				}
			}
		}
		i++
	}

	fmt.Println(s.Solve)
}

func (s *Vns) Reverse(path []int) []int {
	r1 := rand.Intn(s.CityNum)
	r2 := rand.Intn(s.CityNum)
	for r1 > r2 {
		r1, r2 = r2, r1
	}

	temp := path[r1 : r2+1]
	var tempSolve []int
	tempSolve = append(tempSolve, path[:r1]...)
	for i := r2 - r1; i >= 0; i-- {
		tempSolve = append(tempSolve, temp[i])
	}

	if r2 != s.CityNum-1 {
		tempSolve = append(tempSolve, path[r2+1:]...)
	}

	return tempSolve
}

func (s *Vns) Insert(path []int) []int {
	r1 := rand.Intn(s.CityNum)
	r2 := rand.Intn(s.CityNum)
	for r1 > r2 {
		r1, r2 = r2, r1
	}

	var temp []int
	if r1 < r2 {
		temp = []int{path[r1], path[r2]}
	} else {
		temp = []int{path[r1]}
	}

	if r2-r1 > 1 {
		temp = append(temp, path[r1+1:r2]...)
	}

	if r2 < s.CityNum-2 {
		temp = append(temp, path[r2+1:s.CityNum]...)
	}

	if r1 > 0 {
		temp = append(temp, path[:r1]...)
	}

	tempSolve := make([]int, s.CityNum)
	copy(tempSolve, temp)

	return tempSolve
}

func (s *Vns) Exchange(path []int) []int {
	r1 := rand.Intn(s.CityNum)
	r2 := rand.Intn(s.CityNum)
	if r1 == r2 {
		r2 = rand.Intn(s.CityNum)
	}

	//temp := make([]int, s.CityNum)
	//copy(temp, path)
	path[r1], path[r2] = path[r2], path[r1]

	//fmt.Println(temp, "\n", path)
	return path
}

func (s *Vns) CalculateDis(path []int) float64 {
	value := 0.0
	for i := 0; i < len(path)-1; i++ {
		value += s.Data[path[i]][path[i+1]]
	}
	value += s.Data[path[0]][path[s.CityNum-1]]

	return value
}
