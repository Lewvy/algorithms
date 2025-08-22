package algorithms

import (
	"errors"
)

type WeightedUnion struct {
	count int
	n     int
	size  []int
	field []int
}

func NewWeightedUnion(n int) (*WeightedUnion, error) {
	if n <= 0 {
		return nil, errors.New("n can't be less than 1")
	}
	wu := &WeightedUnion{}
	wu.n = n
	wu.field = make([]int, n)
	wu.size = make([]int, n)
	wu.count = n

	for i := range wu.field {
		wu.field[i] = i
		wu.size[i] = 1
	}
	return wu, nil
}

func (wu *WeightedUnion) Find(i int) int {

	for wu.field[i] != i {
		wu.field[i] = wu.field[wu.field[i]]
		i = wu.field[i]
	}
	return i
}

func (wu *WeightedUnion) Union(p int, q int) {

	i := wu.Find(p)
	j := wu.Find(q)

	if i == j {
		return
	}
	if wu.size[i] > wu.size[j] {
		wu.field[j] = i
		wu.size[i] += wu.size[j]
	} else {
		wu.field[i] = j
		wu.size[j] += wu.size[i]
	}

	wu.count--
}
func (wu *WeightedUnion) Connected(p int, q int) bool {
	return wu.Find(p) == wu.Find(q)
}
func (wu *WeightedUnion) CountIslands() int {
	return wu.count
}
