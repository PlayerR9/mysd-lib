package slices

import (
	"slices"
	"testing"

	test "github.com/PlayerR9/go-verify/test"
)

// TestMayInsert tests the MayInsert function.
func TestMayInsert(t *testing.T) {
	type args struct {
		Slice         []int
		Elem          int
		WantInserted  bool
		ExpectedSlice []int
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			ok, err := MayInsert(&args.Slice, args.Elem)
			if err != nil {
				t.Errorf("want no error, got %v", err)

				return
			}

			if ok != args.WantInserted {
				t.Errorf("want %t, got %t", args.WantInserted, ok)

				return
			}

			ok = slices.Equal(args.Slice, args.ExpectedSlice)
			if !ok {
				t.Errorf("want %v, got %v", args.ExpectedSlice, args.Slice)
			}
		}
	})

	_ = tests.AddTest("successful insert", args{
		Slice:         []int{1, 2},
		Elem:          3,
		WantInserted:  true,
		ExpectedSlice: []int{1, 2, 3},
	})

	_ = tests.AddTest("unsuccessful insert", args{
		Slice:         []int{1, 2},
		Elem:          1,
		WantInserted:  false,
		ExpectedSlice: []int{1, 2},
	})

	_ = tests.Run(t)
}

// TestUniquefy tests the Uniquefy function.
func TestUniquefy(t *testing.T) {
	type args struct {
		Slice         []int
		WantRemoved   int
		ExpectedSlice []int
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			n := Uniquefy(&args.Slice)
			if n != args.WantRemoved {
				t.Errorf("want %d, got %d", args.WantRemoved, n)

				return
			}

			ok := slices.Equal(args.Slice, args.ExpectedSlice)
			if !ok {
				t.Errorf("want %v, got %v", args.ExpectedSlice, args.Slice)
			}
		}
	})

	_ = tests.AddTest("unique slice", args{
		Slice:         []int{1, 2, 3},
		WantRemoved:   0,
		ExpectedSlice: []int{1, 2, 3},
	})

	_ = tests.AddTest("non-unique slice", args{
		Slice:         []int{1, 2, 3, 2},
		WantRemoved:   1,
		ExpectedSlice: []int{1, 2, 3},
	})

	_ = tests.Run(t)
}

// TestMerge tests the Merge function.
func TestMerge(t *testing.T) {
	type args struct {
		Slice         []int
		From          []int
		WantIgnored   int
		ExpectedSlice []int
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			n, err := Merge(&args.Slice, args.From)
			if err != nil {
				t.Errorf("want no error, got %v", err)

				return
			}

			if n != args.WantIgnored {
				t.Errorf("want %d, got %d", args.WantIgnored, n)

				return
			}

			ok := slices.Equal(args.Slice, args.ExpectedSlice)
			if !ok {
				t.Errorf("want %v, got %v", args.ExpectedSlice, args.Slice)
			}
		}
	})

	_ = tests.AddTest("successful merge", args{
		Slice:         []int{1, 2},
		From:          []int{3, 4},
		WantIgnored:   0,
		ExpectedSlice: []int{1, 2, 3, 4},
	})

	_ = tests.AddTest("identical merge", args{
		Slice:         []int{1, 2},
		From:          []int{1, 2},
		WantIgnored:   2,
		ExpectedSlice: []int{1, 2},
	})

	_ = tests.Run(t)
}
