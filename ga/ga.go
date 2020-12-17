package ga

import (
	"fmt"
	"math/rand"
)

type Individual struct {
	Value float64
	Fit   float64
	Ind   []int
}

type GA struct {
	PopSize    int
	MaxGen     int
	PM         float64
	PC         float64
	Gap        float64
	Pop        []Individual
	SonPop     []Individual
	Data       [][]float64
	CodeLen    int
	BestInd    Individual
	Fits       []float64
	BestValues []int
}

func (s *GA) Run() {
	s.PopSize = 100
	s.MaxGen = 1
	s.PM = 0.25
	s.PC = 0.7
	s.Gap = 0.9
	s.CodeLen = len(s.Data)
	s.CreatePop()
	s.SolvePop()
	s.BestInd = s.Pop[0]
	for gen := 0; gen < s.MaxGen; gen++ {
		s.Ranking()
		s.Select()
		s.Cross()
		s.Mutate()
		s.Reverse()
		s.Rein()
		s.SolvePop()
		if s.Pop[0].Value < s.BestInd.Value {
			s.BestInd = s.Pop[0]
		}
		s.BestValues = append(s.BestValues, int(s.BestInd.Value))
	}

	s.BestInd.Ind = []int{9, 8, 7, 3, 15, 22, 10, 23, 18, 16, 2, 17, 21, 20, 19, 24, 25, 27, 26, 29, 30, 28, 0, 14, 13, 11, 12, 6, 5, 4, 1}
	s.SolveInd(&s.BestInd)
	fmt.Println(s.BestValues, "\n", s.BestInd)
}

// 计算适应度
func (s *GA) Ranking() {
	sumValue := 0.0
	for _, ind := range s.Pop {
		sumValue += 1 / ind.Value
	}

	s.Fits = make([]float64, s.PopSize)
	s.Pop[0].Fit = (1 / s.Pop[0].Value) / sumValue
	s.Fits[0] = s.Pop[0].Fit
	for i := 1; i < s.PopSize; i++ {
		s.Pop[i].Fit = (1 / s.Pop[i].Value) / sumValue
		s.Fits[i] = s.Fits[i-1] + s.Pop[i].Fit
	}
}

// 选择
func (s *GA) Select() {
	sonPopSize := int(float64(s.PopSize) * s.Gap)
	s.SonPop = make([]Individual, sonPopSize)
	for i := range s.SonPop {
		s.SonPop[i].Ind = make([]int, s.CodeLen)
		a := rand.Float64()
		for k, v := range s.Fits {
			if a <= v {
				copy(s.SonPop[i].Ind, s.Pop[k].Ind)
				break
			}
		}
	}
}

// 交叉
func (s *GA) Cross() {
	sonPopSize := len(s.SonPop)
	n := sonPopSize / 2
	index := rand.Perm(sonPopSize)
	for i := 0; i < n; i++ {
		a := rand.Float64()
		if a < s.PC {
			ind1 := s.SonPop[index[i]]
			ind2 := s.SonPop[index[n+i]]
			exchange(ind1.Ind, ind2.Ind)
		}
	}
}

// 变异
func (s *GA) Mutate() {
	for i := range s.SonPop {
		a := rand.Float64()
		if a < s.PC {
			r1 := rand.Intn(s.CodeLen)
			r2 := rand.Intn(s.CodeLen)
			if r1 > r2 {
				r1, r2 = r2, r1
			}
			s.SonPop[i].Ind[r1], s.SonPop[i].Ind[r2] = s.SonPop[i].Ind[r2], s.SonPop[i].Ind[r1]
		}
	}
}

// 逆转
func (s *GA) Reverse() {
	for j := range s.SonPop {
		var ind Individual
		ind.Ind = make([]int, s.CodeLen)
		copy(ind.Ind, s.SonPop[j].Ind)
		//s.SolveInd(&s.SonPop[j])				// 居家调度
		s.SolveInd(&s.SonPop[j]) // 机构调度
		r1 := rand.Intn(s.CodeLen)
		r2 := rand.Intn(s.CodeLen)
		if r1 > r2 {
			r1, r2 = r2, r1
		}

		gen := ind.Ind[r1 : r2+1]
		for i := 0; i < (r2-r1)/2; i++ {
			gen[i], gen[r2-r1-i] = gen[r2-r1-i], gen[i]
		}
		copy(ind.Ind[r1:r2+1], gen)
		//s.SolveInd(&s.SonPop[j])				// 居家调度
		s.SolveInd(&s.SonPop[j]) // 机构调度
		if ind.Value < s.SonPop[j].Value {
			copy(s.SonPop[j].Ind, ind.Ind)
		}
	}
}

// 基因重组
func (s *GA) Rein() {
	sonPopSize := len(s.SonPop)
	copy(s.Pop[s.PopSize-sonPopSize:], s.SonPop)
}

func exchange(a, b []int) {
	num := len(a)
	r1 := rand.Intn(num)
	r2 := rand.Intn(num)
	if r1 > r2 {
		r1, r2 = r2, r1
	}

	a1 := make([]int, num)
	b1 := make([]int, num)
	for i := r1; i < r2; i++ {
		copy(a1, a)
		copy(b1, b)
		a[i], b[i] = b[i], a[i]

		for j := 0; j < num; j++ {
			if a[j] == a[i] && j != i {
				a[j] = a1[i]
			}
			if b[j] == b[i] && j != i {
				b[j] = b1[i]
			}

			if a[j] == a1[i] && b[j] == b1[i] {
				break
			}
		}
	}
}
