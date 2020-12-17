package aco

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"go_tsp/mymath"
)

type Path struct {
	Value float64 // 路径长度
	Route []int   // 访问路线
	Tabu  []int   // 已访问过城市
	Allow []int   // 未访问过城市
}

type Aco struct {
	PopSize    int         // 种群大小
	MaxGen     int         // 最大迭代次数
	Alpha      float64     // 信息素重要程度因子
	Beta       float64     // 启发函数重要程度因子
	Rho        float64     // 信息素挥发因子
	Q          float64     // 常系数
	Eta        [][]float64 // 启发函数
	Tau        [][]float64 // 信息素矩阵
	Table      []Path      // 路径记录
	PathBest   []Path      // 各代最佳路径
	LengthBest []float64   // 各代最佳路径的长度
	LengthAve  []float64   // 各代路径的平均长度
	Data       [][]float64 // 路网
	CityNum    int         // 城市数量
}

func (s *Aco) Run() {
	s.PopSize = 50
	s.MaxGen = 200
	s.Alpha = 1
	s.Beta = 5
	s.Rho = 0.1
	s.Q = 1
	s.CityNum = len(s.Data)
	s.CreateEtaTau()
	s.PathBest = make([]Path, s.MaxGen)
	s.LengthBest = make([]float64, s.MaxGen)
	for gen := 0; gen < s.MaxGen; gen++ {
		s.Table = make([]Path, s.PopSize)
		// 初始化起始点和待访问城市
		for i := 0; i < s.PopSize; i++ {
			s.Table[i].Allow = make([]int, s.CityNum)
			for j := 0; j < s.CityNum; j++ {
				s.Table[i].Allow[j] = j
			}
			s.Table[i].Route = make([]int, s.CityNum)
			s.Table[i].Route[0] = rand.Intn(s.CityNum)
		}

		// 路径搜索
		for i := 0; i < s.PopSize; i++ {
			for j := 1; j < s.CityNum; j++ {
				s.Table[i].Tabu = append(s.Table[i].Tabu, s.Table[i].Route[j-1])
				s.Table[i].Allow = mymath.DelInt(s.Table[i].Allow, s.Table[i].Route[j-1])
				p := make([]float64, s.CityNum-j) // 下一个城市被选择的概率
				for k := 0; k < len(p); k++ {
					p[k] = math.Pow(s.Tau[s.Table[i].Tabu[j-1]][s.Table[i].Allow[k]], s.Alpha) * math.Pow(s.Eta[s.Table[i].Tabu[j-1]][s.Table[i].Allow[k]], s.Beta)
				}
				sumP := mymath.SumFloat64(p)
				for k := 0; k < len(p); k++ {
					p[k] /= sumP
				}

				// 轮盘赌选择下一个城市
				pc := mymath.CumSumFloat(p)
				r := rand.Float64()
				index := mymath.FindFloat64(pc, r)
				target := s.Table[i].Allow[index]
				s.Table[i].Route[j] = target
			}
		}
		s.CalLength()
		sort.Sort(Paths(s.Table))
		s.PathBest[gen] = s.Table[0]
		s.LengthBest[gen] = s.Table[0].Value

		tmpTau := make([][]float64, s.CityNum)
		for i := 0; i < s.CityNum; i++ {
			tmpTau[i] = make([]float64, s.CityNum)
		}
		for i := 0; i < s.PopSize; i++ {
			for j := 0; j < s.CityNum-1; j++ {
				tmpTau[s.Table[i].Route[j]][s.Table[i].Route[j+1]] += s.Q / s.Table[i].Value
			}
			tmpTau[s.Table[i].Route[s.CityNum-1]][s.Table[i].Route[0]] += s.Q / s.Table[i].Value
		}

		for i := 0; i < s.CityNum; i++ {
			for j := 0; j < s.CityNum-1; j++ {
				s.Tau[i][j] = (1-s.Rho)*s.Tau[i][j] + tmpTau[i][j]
			}
		}
	}

	sort.Sort(Paths(s.PathBest))

	fmt.Println(s.PathBest[0].Value, len(s.PathBest[0].Route), s.PathBest[0].Route)
	for _, path := range s.PathBest {
		fmt.Println(path.Value)
	}
}

func (s *Aco) CreateEtaTau() {
	s.Tau = make([][]float64, s.CityNum)
	s.Eta = make([][]float64, s.CityNum)
	for i := 0; i < s.CityNum; i++ {
		s.Tau[i] = make([]float64, s.CityNum)
		s.Eta[i] = make([]float64, s.CityNum)
		for j := 0; j < s.CityNum; j++ {
			s.Tau[i][j] = 1
			if i != j && s.Data[i][j] != 0 {
				s.Eta[i][j] = 1 / s.Data[i][j]
			} else {
				s.Eta[i][j] = 0.0001
			}
		}
	}
}

func (s *Aco) CalLength() {
	for k, path := range s.Table {
		path.Value = 0
		for i := 0; i < s.CityNum-1; i++ {
			path.Value += s.Data[path.Route[i]][path.Route[i+1]]
		}
		path.Value += s.Data[path.Route[0]][path.Route[s.CityNum-1]]
		s.Table[k].Value = path.Value

	}
}

type Paths []Path

var _ sort.Interface = Paths{}

func (s Paths) Len() int {
	return len(s)
}

func (s Paths) Less(i, j int) bool {
	return s[i].Value < s[j].Value
}

func (s Paths) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
