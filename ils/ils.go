package ils

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	Route []int
	Cost  float64
}

type ILS struct {
	MaxIter      int
	MaxNoImprove int
	Cities       [][]int
	Data         [][]float64
	CitySize     int
	BestSolution Solution
	Delta        [][]float64
}

func (s *ILS) Run() {
	s.Delta = make([][]float64, s.CitySize)
	for i := 0; i < s.CitySize; i++ {
		s.Delta[i] = make([]float64, s.CitySize)
	}

	s.IteratedLocalSearch()

	fmt.Println(s.BestSolution.Cost)
	fmt.Println(s.BestSolution.Route)
}

func (s *ILS) IteratedLocalSearch() {
	//var curSolution Solution

	s.BestSolution.Route = rand.Perm(s.CitySize)
	s.BestSolution.Cost = s.CostTotal(s.BestSolution.Route)
	s.BestSolution = s.LocalSearch(s.BestSolution)

	for i := 0; i < s.MaxIter; i++ {
		curSolution := s.Perturbation()
		curSolution = s.LocalSearch(curSolution)
		if curSolution.Cost < s.BestSolution.Cost {
			copy(s.BestSolution.Route, curSolution.Route)
			s.BestSolution.Cost = curSolution.Cost
		}
	}
}

func (s *ILS) CostTotal(route []int) float64 {
	var total float64
	for i := 0; i < s.CitySize-1; i++ {
		total += s.Data[route[i]][route[i+1]]
	}

	total += s.Data[route[s.CitySize-1]][route[0]]

	return total
}

func (s *ILS) LocalSearch(solution Solution) Solution {
	initalCost := solution.Cost
	var nowCost float64
	var curSolution Solution
	for i := 0; i < s.CitySize-1; i++ {
		for j := i + 1; j < s.CitySize; j++ {
			s.Delta[i][j] = s.CalCostDelta(i, j, solution.Route)
		}
	}

	count := 0
	for count < s.MaxNoImprove {
		for i := 0; i < s.CitySize; i++ {
			for j := 0; j < s.CitySize; j++ {
				curSolution.Route = s.TwoOptSwap(i, j, solution.Route)
				nowCost = initalCost + s.Delta[i][j]
				curSolution.Cost = nowCost
				if curSolution.Cost < solution.Cost {
					count = 0
					copy(solution.Route, curSolution.Route)
					solution.Cost = curSolution.Cost
					initalCost = curSolution.Cost
					s.Update(i, j, solution.Route)
				}
			}
		}
		count++
	}

	return solution
}

func (s *ILS) CalCostDelta(i, j int, tmp []int) float64 {
	var delta float64
	if i == 0 {
		if j == s.CitySize-1 {
			delta = 0
		} else {
			delta = 0 - s.Data[tmp[j]][tmp[j+1]] + s.Data[tmp[i]][tmp[j+1]] -
				s.Data[tmp[s.CitySize-1]][tmp[i]] + s.Data[tmp[s.CitySize-1]][tmp[j]]
		}
	} else {
		if j == s.CitySize-1 {
			delta = 0 - s.Data[tmp[i-1]][tmp[i]] + s.Data[tmp[i-1]][tmp[j]] -
				s.Data[tmp[0]][tmp[j]] + s.Data[tmp[i]][tmp[0]]
		} else {
			delta = 0 - s.Data[tmp[i-1]][tmp[i]] + s.Data[tmp[i-1]][tmp[j]] -
				s.Data[tmp[j]][tmp[j+1]] + s.Data[tmp[i]][tmp[j+1]]
		}
	}
	return delta
}

func (s *ILS) TwoOptSwap(i, j int, bestRoute []int) []int {
	newRoute := make([]int, s.CitySize)
	copy(newRoute, bestRoute)
	s.SwapElement(i, j, newRoute)
	return newRoute
}

func (s *ILS) SwapElement(i, j int, route []int) {
	for i < j {
		route[i], route[j] = route[j], route[i]
		i++
		j--
	}
}

func (s *ILS) Update(i, j int, route []int) {
	if i > 0 && j != s.CitySize-1 {
		i--
		j++
		for k := i; k <= j; k++ {
			for l := j + 1; l < s.CitySize; l++ {
				s.Delta[k][l] = s.CalCostDelta(k, l, route)
			}
		}

		for k := 0; k < j; k++ {
			for l := i; l <= j; l++ {
				if k >= l {
					continue
				}

				s.Delta[k][l] = s.CalCostDelta(k, l, route)
			}
		}
	} else {
		for i := 0; i < s.CitySize-1; i++ {
			for j := i + 1; j < s.CitySize; j++ {
				s.Delta[i][j] = s.CalCostDelta(i, j, route)
			}
		}
	}
}

// 扰动
func (s *ILS) Perturbation() Solution {
	var curSolution Solution
	curSolution.Route = s.DoubleBridgeMove()
	curSolution.Cost = s.CostTotal(curSolution.Route)
	return curSolution
}

//将城市序列分成4块，然后按块重新打乱顺序。
//用于扰动函数
func (s *ILS) DoubleBridgeMove() []int {
	var tempRoute []int
	pos1 := 1 + rand.Intn(s.CitySize/4)
	pos2 := pos1 + 1 + rand.Intn(s.CitySize/4)
	pos3 := pos2 + 1 + rand.Intn(s.CitySize/4)

	tempRoute = append(tempRoute, s.BestSolution.Route[:pos1]...)
	tempRoute = append(tempRoute, s.BestSolution.Route[pos3:]...)
	tempRoute = append(tempRoute, s.BestSolution.Route[pos2:pos3]...)
	tempRoute = append(tempRoute, s.BestSolution.Route[pos1:pos2]...)

	return tempRoute
}
