package algorithms_test

import (
	"algorithms/week1"
	"testing"
)

func TestQuickUnion_New(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    *algorithms.QuickFind
		wantErr bool
	}{
		{
			name:    "Valid initialization with n=5",
			n:       5,
			wantErr: false,
		},
		{
			name:    "Invalid initialization with n=0",
			n:       0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q, err := algorithms.NewQuickUnion(tt.n)

			if err != nil && tt.wantErr == false {
				t.Errorf("%q\n", err)
			}

			if q.CountIslands() != tt.n {
				t.Errorf("%q\n", err)
			}
		})
	}
}

func TestQuickUnion_Behaviors(t *testing.T) {
	qu, _ := algorithms.NewQuickUnion(10)

	if qu.CountIslands() != 10 {
		t.Fatalf("Initial count should be 10, got %d", qu.CountIslands())
	}
	if qu.Connected(0, 1) {
		t.Fatal("0 and 1 should not be connected initially")
	}

	qu.Union(0, 1)
	if !qu.Connected(0, 1) {
		t.Error("Connected(0, 1) should be true after union")
	}
	if qu.CountIslands() != 9 {
		t.Errorf("Count should be 9 after one union, got %d", qu.CountIslands())
	}

	qu.Union(1, 2)
	if !qu.Connected(0, 2) {
		t.Error("Connected(0, 2) should be true after transitive union")
	}
	if qu.CountIslands() != 8 {
		t.Errorf("Count should be 8 after two unions, got %d", qu.CountIslands())
	}

	qu.Union(5, 6)
	qu.Union(6, 7)
	if qu.Connected(0, 5) {
		t.Error("Components {0,1,2} and {5,6,7} should not be connected")
	}
	qu.Union(0, 7)
	if !qu.Connected(2, 5) {
		t.Error("Connected(2, 5) should be true after merging components")
	}
	if qu.CountIslands() != 5 {
		t.Errorf("Count should be 5 after merging components, got %d", qu.CountIslands())
	}

	t.Run("Redundant union on connected components", func(t *testing.T) {
		initialCount := qu.CountIslands()
		qu.Union(0, 2)
		if qu.CountIslands() != initialCount {
			t.Errorf("Count should not change on redundant union. Was %d, now %d",
				initialCount, qu.CountIslands())
		}
	})
}

func TestQuickUnion_Find_WorstCase(t *testing.T) {
	qu, _ := algorithms.NewQuickUnion(10)

	qu.Union(0, 1)
	qu.Union(1, 2)
	qu.Union(2, 3)
	qu.Union(3, 4)

	expectedRoot := 4
	t.Run("Find root from bottom of the chain", func(t *testing.T) {
		if root := qu.Find(0); root != expectedRoot {
			t.Errorf("Find(0) should be %d, got %d", expectedRoot, root)
		}
	})

	t.Run("Find root from middle of the chain", func(t *testing.T) {
		if root := qu.Find(2); root != expectedRoot {
			t.Errorf("Find(2) should be %d, got %d", expectedRoot, root)
		}
	})

	t.Run("Find root of an element that is already a root", func(t *testing.T) {
		if root := qu.Find(4); root != expectedRoot {
			t.Errorf("Find(4) should be %d, got %d", expectedRoot, root)
		}
	})

	t.Run("Find root of an unconnected element", func(t *testing.T) {
		if root := qu.Find(9); root != 9 {
			t.Errorf("Find(9) should be 9, got %d", root)
		}
	})
}
