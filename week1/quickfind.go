package algorithms

import "errors"

type UF interface {
	Union(p, q int)
	Connected(p, q int) bool
	Count() int
	Find() int
}

type QuickFind struct {
	Count int
	N     int
	Field []int
}

func NewQuickFind(n int) (*QuickFind, error) {
	qf := &QuickFind{}
	if n <= 0 {
		return nil, errors.New("n can't be less than 1")
	}
	qf.N = n
	qf.Field = make([]int, n)
	qf.Count = n

	for i := range qf.Field {
		qf.Field[i] = i
	}
	return qf, nil
}

func (qf *QuickFind) Find(i int) int {
	return qf.Field[i]
}

func (qf *QuickFind) Union(p int, q int) {
	if qf.Field[p] == qf.Field[q] {
		return
	}

	prev_parent := qf.Field[p]
	new_parent := qf.Field[q]

	for i := 0; i < qf.N; i++ {
		if qf.Field[i] == prev_parent {
			qf.Field[i] = new_parent
		}
	}
	qf.Count--
}
func (qf *QuickFind) Connected(p int, q int) bool {
	return qf.Field[p] == qf.Field[q]
}
func (qf *QuickFind) CountIslands() int {
	return qf.Count
}
