package algorithms

import "errors"

type QuickUnion struct {
	count int
	n     int
	field []int
}

func NewQuickUnion(n int) (*QuickUnion, error) {
	qu := &QuickUnion{}
	if n <= 0 {
		return nil, errors.New("n can't be less than 1")
	}
	qu.n = n
	qu.field = make([]int, n)
	qu.count = n

	for i := range qu.field {
		qu.field[i] = i
	}
	return qu, nil
}

func (qu *QuickUnion) Find(i int) int {

	for qu.field[i] != i {
		i = qu.field[i]
	}

	return i
}

func (qu *QuickUnion) Union(p int, q int) {

	i := qu.Find(p)
	j := qu.Find(q)

	qu.field[i] = j

	if i != j {
		qu.count--
	}
}
func (qu *QuickUnion) Connected(p int, q int) bool {
	return qu.Find(p) == qu.Find(q)
}
func (qu *QuickUnion) CountIslands() int {
	return qu.count
}
