package algorithms

import "errors"

type UF interface {
	Init(n int) *UF
	Union(p, q int)
	Connected(p, q int) bool
	Count() int
}

type QuickFindUF struct {
	Count int
	N     int
	Field []int
}

func (qf *QuickFindUF) Init(n int) (*QuickFindUF, error) {
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
func (qf *QuickFindUF) Union(p int, q int) {
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
func (qf *QuickFindUF) Connected(p int, q int) bool {
	return qf.Field[p] == qf.Field[q]
}
func (qf *QuickFindUF) CountIslands() int {
	return qf.Count
}
