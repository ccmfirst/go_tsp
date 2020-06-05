package ga

import (
	"math/rand"
	"sort"
)

func (s *GA) CreatePop() {
	pop := make([]Individual, s.PopSize)
	for i := 0; i < s.PopSize; i++ {
		pop[i].Ind = rand.Perm(s.CodeLen)
	}
	s.Pop = pop
}

func (s *GA) SolvePop() {
	for i := range s.Pop {
		s.SolveInd(&s.Pop[i])
	}

	sort.Sort(Pop(s.Pop))
}

func (s *GA) SolveInd(ind *Individual) {
	ind.Value = 0
	for i := 0; i < len(ind.Ind)-1; i++ {
		ind.Value += s.Data[ind.Ind[i]][ind.Ind[i+1]]
	}
	ind.Value += s.Data[ind.Ind[0]][ind.Ind[s.CodeLen-1]]
}

type Pop []Individual

var _ sort.Interface = Pop{}

func (s Pop) Len() int {
	return len(s)
}

func (s Pop) Less(i, j int) bool {
	return s[i].Value < s[j].Value
}

func (s Pop) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
