package slices

import (
	"slices"
	"testing"

	test "github.com/PlayerR9/go-verify/test"
)

// TestApplyFilter tests the ApplyFilter function.
func TestApplyFilter(t *testing.T) {
	const (
		MAX int = 100
	)

	expected_evens := make([]int, 0, MAX/2)
	for i := 0; i < MAX; i += 2 {
		expected_evens = append(expected_evens, i)
	}

	nums := make([]int, 0, 100)
	for i := 0; i < MAX; i++ {
		nums = append(nums, i)
	}

	type args struct {
		elems     []int
		predicate Predicate[int]
		expected  []int
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			actual := ApplyFilter(args.elems, args.predicate)

			ok := slices.Equal(actual, args.expected)
			if !ok {
				t.Errorf("expected %v, got %v", args.expected, actual)
			}
		}
	})

	_ = tests.AddTest("success", args{
		elems:     nums,
		predicate: func(x int) bool { return x%2 == 0 },
		expected:  expected_evens,
	})

	_ = tests.AddTest("no filter", args{
		elems:     nums,
		predicate: nil,
		expected:  nil,
	})

	_ = tests.Run(t)
}

// TestApplyReject tests the ApplyReject function.
func TestApplyReject(t *testing.T) {
	const (
		MAX int = 100
	)

	expected_odds := make([]int, 0, MAX/2)
	for i := 1; i < MAX; i += 2 {
		expected_odds = append(expected_odds, i)
	}

	nums := make([]int, 0, 100)
	for i := 0; i < MAX; i++ {
		nums = append(nums, i)
	}

	type args struct {
		elems     []int
		predicate Predicate[int]
		expected  []int
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			actual := ApplyReject(args.elems, args.predicate)

			ok := slices.Equal(actual, args.expected)
			if !ok {
				t.Errorf("expected %v, got %v", args.expected, actual)
			}
		}
	})

	_ = tests.AddTest("success", args{
		elems:     nums,
		predicate: func(x int) bool { return x%2 == 0 },
		expected:  expected_odds,
	})

	_ = tests.AddTest("no filter", args{
		elems:     nums,
		predicate: nil,
		expected:  nil,
	})

	_ = tests.Run(t)
}

type MockStruct struct {
}

func (ms *MockStruct) IsNil() bool {
	return ms == nil
}

// TestRejectNils tests the RejectNils function.
func TestRejectNils(t *testing.T) {
	type args struct {
		tokens   []*MockStruct
		expected []*MockStruct
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			result := RejectNils(args.tokens)

			if len(result) != len(args.expected) {
				t.Errorf("want %d elements, got %d", len(args.expected), len(result))
			} else {
				for i := 0; i < len(result); i++ {
					if result[i] == nil {
						t.Errorf("at index %d, want %p, got nil", i, args.expected[i])
					}
				}
			}
		}
	})

	_ = tests.AddTest("no elems", args{
		tokens:   []*MockStruct{},
		expected: []*MockStruct{},
	})

	_ = tests.AddTest("one elem", args{
		tokens:   []*MockStruct{{}},
		expected: []*MockStruct{{}},
	})

	_ = tests.AddTest("two elems", args{
		tokens:   []*MockStruct{nil, {}},
		expected: []*MockStruct{{}},
	})

	_ = tests.AddTest("all nil", args{
		tokens:   []*MockStruct{nil, nil},
		expected: nil,
	})

	_ = tests.Run(t)
}
