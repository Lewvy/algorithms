package algorithms_test

import (
	"algorithms"
	"reflect"
	"testing"
)

func TestQuickFindUF_Init(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    *algorithms.QuickFindUF
		wantErr bool
	}{
		{
			name: "Valid initialization with n=5",
			n:    5,
			// The expected struct should match the implementation: []int slice, and count/n fields set.
			want: &algorithms.QuickFindUF{
				Field: []int{0, 1, 2, 3, 4},
				Count: 5,
				N:     5,
			},
			wantErr: false,
		},
		{
			name:    "Invalid initialization with n=0",
			n:       0,
			wantErr: true,
		},
		{
			name:    "Invalid initialization with n=-1",
			n:       -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var qf algorithms.QuickFindUF
			got, err := qf.Init(tt.n)

			// Check if an error occurred
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// If we expected an error, we can stop here.
			if tt.wantErr {
				return
			}
			// The comparison was inverted. We want to error if they are NOT equal.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestQuickFindUF_Union(t *testing.T) {
	// We define an initial state for the UF struct to run the Union method on.
	initialUF := func() *algorithms.QuickFindUF {
		qf := &algorithms.QuickFindUF{}
		qf.Init(10)
		return qf
	}

	tests := []struct {
		name      string
		initialQF *algorithms.QuickFindUF
		p         int
		q         int
		wantField []int
		wantCount int
	}{
		{
			name:      "Simple union of 4 and 3",
			initialQF: initialUF(),
			p:         4,
			q:         3,
			wantField: []int{0, 1, 2, 3, 3, 5, 6, 7, 8, 9}, // All 4's become 3's
			wantCount: 9,
		},
		{
			name: "Union on an already modified set",
			initialQF: &algorithms.QuickFindUF{
				Field: []int{0, 1, 8, 8, 1, 5, 6, 7, 8, 9}, // Start with 2->8, 3->8, 4->1
				Count: 7,
				N:     10,
			},
			p:         1,
			q:         8,
			wantField: []int{0, 8, 8, 8, 8, 5, 6, 7, 8, 9}, // All 1's (and what pointed to 1) become 8's
			wantCount: 6,
		},
		{
			name: "Union of already connected elements",
			initialQF: &algorithms.QuickFindUF{
				Field: []int{0, 1, 1, 3, 4, 5, 6, 7, 8, 9}, // 2 is connected to 1
				Count: 9,
				N:     10,
			},
			p:         2,
			q:         1,
			wantField: []int{0, 1, 1, 3, 4, 5, 6, 7, 8, 9}, // Should not change
			wantCount: 9,                                   // Count should not change
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qf := tt.initialQF
			qf.Union(tt.p, tt.q)
			if !reflect.DeepEqual(qf.Field, tt.wantField) {
				t.Errorf("Union() resulted in Field = %v, want %v", qf.Field, tt.wantField)
			}
			if qf.Count != tt.wantCount {
				t.Errorf("Union() resulted in Count = %v, want %v", qf.Count, tt.wantCount)
			}
		})
	}
}

func TestQuickFindUF_Connected(t *testing.T) {
	qf := &algorithms.QuickFindUF{
		Field: []int{0, 1, 1, 8, 8, 0, 0, 7, 8, 9}, // Components: {0,5,6}, {1,2}, {3,4,8}, {7}, {9}
		Count: 5,
		N:     10,
	}

	tests := []struct {
		name string
		p    int
		q    int
		want bool
	}{
		{name: "2 and 1 are connected", p: 2, q: 1, want: true},
		{name: "5 and 6 are connected", p: 5, q: 6, want: true},
		{name: "0 and 5 are connected", p: 0, q: 5, want: true},
		{name: "3 and 8 are connected", p: 3, q: 8, want: true},
		{name: "1 and 8 are NOT connected", p: 1, q: 8, want: false},
		{name: "7 and 9 are NOT connected", p: 7, q: 9, want: false},
		{name: "Same element is connected", p: 4, q: 4, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := qf.Connected(tt.p, tt.q); got != tt.want {
				t.Errorf("Connected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickFindUF_Count(t *testing.T) {
	qf, _ := (&algorithms.QuickFindUF{}).Init(10)

	if qf.CountIslands() != 10 {
		t.Errorf("Initial count should be 10, got %d", qf.CountIslands())
	}

	qf.Union(1, 2)
	if qf.CountIslands() != 9 {
		t.Errorf("Count after one union should be 9, got %d", qf.CountIslands())
	}

	qf.Union(3, 4)
	if qf.CountIslands() != 8 {
		t.Errorf("Count after two unions should be 8, got %d", qf.CountIslands())
	}

	qf.Union(1, 2) // Union on already connected components
	if qf.CountIslands() != 8 {
		t.Errorf("Count should not change on redundant union, got %d", qf.CountIslands())
	}

	qf.Union(1, 3) // Merge two components {1,2} and {3,4}
	if qf.CountIslands() != 7 {
		t.Errorf("Count after merging components should be 7, got %d", qf.CountIslands())
	}
}
