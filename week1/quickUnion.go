package algorithms

import "errors"

type QuickUnion struct {
	Count int
	N     int
	Field []int
}

func (qu *QuickUnion) New(n int) (*QuickUnion, error) {
	if n <= 0 {
		return nil, errors.New("n can't be less than 1")
	}
	qu.N = n
	qu.Field = make([]int, n)
	qu.Count = n

	for i := range qu.Field {
		qu.Field[i] = i
	}
	return qu, nil
}

func (qu *QuickUnion) Find(i int) int {

	for qu.Field[i] != i {
		i = qu.Field[i]
	}

	return i
}

func (qu *QuickUnion) Union(p int, q int) {
	if qu.Field[p] == qu.Field[q] {
		return
	}

	i := qu.Find(p)
	j := qu.Find(q)

	qu.Field[i] = j

	if i != j {
		qu.Count--
	}
}
func (qu *QuickUnion) Connected(p int, q int) bool {
	return qu.Find(p) == qu.Find(q)
}
func (qu *QuickUnion) CountIslands() int {
	return qu.Count
}
