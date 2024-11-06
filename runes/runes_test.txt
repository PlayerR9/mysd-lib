package runes

import (
	"slices"
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

// TestBytesToUtf8 tests the BytesToUtf8 function.
func TestBytesToUtf8(t *testing.T) {
	type args struct {
		b            []byte
		expected_err string
		expected_res []rune
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			chars, err := BytesToUtf8(args.b)

			err = test.CheckErr(args.expected_err, err)
			if err != nil {
				t.Error(err)
			} else {
				ok := slices.Equal(chars, args.expected_res)
				if !ok {
					t.Errorf("expected %v, got %v", args.expected_res, chars)
				}
			}
		}
	})

	_ = tests.AddTest("with valid bytes", args{
		b:            []byte("test"),
		expected_err: "",
		expected_res: []rune{'t', 'e', 's', 't'},
	})

	_ = tests.AddTest("with invalid bytes", args{
		b:            []byte{0xff, 0xff, 0xff, 0xff},
		expected_err: "byte 0 is not valid utf-8",
		expected_res: nil,
	})

	_ = tests.AddTest("with empty bytes", args{
		b:            []byte{},
		expected_err: "",
		expected_res: []rune{},
	})

	_ = tests.Run(t)
}

// TestStringToUtf8 tests the StringToUtf8 function.
func TestStringToUtf8(t *testing.T) {
	type args struct {
		s            string
		expected_err string
		expected_res []rune
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			chars, err := StringToUtf8(args.s)

			err = test.CheckErr(args.expected_err, err)
			if err != nil {
				t.Error(err)
			} else {
				ok := slices.Equal(chars, args.expected_res)
				if !ok {
					t.Errorf("expected %v, got %v", args.expected_res, chars)
				}
			}
		}
	})

	_ = tests.AddTest("with valid string", args{
		s:            "test",
		expected_err: "",
		expected_res: []rune{'t', 'e', 's', 't'},
	})

	_ = tests.AddTest("with invalid string", args{
		s:            string([]byte{0xff, 0xff, 0xff, 0xff}),
		expected_err: "byte 0 is not valid utf-8",
		expected_res: nil,
	})

	_ = tests.AddTest("with empty string", args{
		s:            "",
		expected_err: "",
		expected_res: []rune{},
	})

	_ = tests.Run(t)
}

// TestIndicesOf tests the IndicesOf function.
func TestIndicesOf(t *testing.T) {
	type args struct {
		data     []rune
		sep      rune
		expected []int
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			indices := IndicesOf(args.data, args.sep)

			ok := slices.Equal(indices, args.expected)
			if !ok {
				t.Errorf("expected %v, got %v", args.expected, indices)
			}
		}
	})

	_ = tests.AddTest("with empty data", args{
		data:     []rune{},
		sep:      'a',
		expected: []int{},
	})

	_ = tests.AddTest("with data and separator", args{
		data:     []rune("test"),
		sep:      't',
		expected: []int{0, 3},
	})

	_ = tests.AddTest("with data and no separator", args{
		data:     []rune("test"),
		sep:      'a',
		expected: nil,
	})

	_ = tests.Run(t)
}

// TestNormalizeRunes tests the NormalizeRunes function.
func TestNormalizeRunes(t *testing.T) {
	type args struct {
		data           []rune
		tab_size       int
		expected_err   string
		expected_slice []rune
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			normalized, err := NormalizeRunes(args.data, args.tab_size)

			err = test.CheckErr(args.expected_err, err)
			if err != nil {
				t.Error(err)
			} else {
				ok := slices.Equal(normalized, args.expected_slice)
				if !ok {
					t.Errorf("expected %v, got %v", args.expected_slice, normalized)
				}
			}
		}
	})

	_ = tests.AddTest("with empty data", args{
		data:           []rune{},
		tab_size:       3,
		expected_err:   "",
		expected_slice: []rune{},
	})

	_ = tests.AddTest("with valid '\\r\\n' data", args{
		data:           []rune{'t', 'e', 's', 't', '\r', '\n', 't', 'e', 's', 't'},
		tab_size:       3,
		expected_err:   "",
		expected_slice: []rune{'t', 'e', 's', 't', '\n', 't', 'e', 's', 't'},
	})

	_ = tests.AddTest("with valid invalid data", args{
		data:           []rune{'t', 'e', 's', 't', '\r', '\r', 'a'},
		tab_size:       3,
		expected_err:   "after '\\r': expected '\\n', got '\\r' instead",
		expected_slice: []rune{'t', 'e', 's', 't', '\r', '\r', 'a'},
	})

	_ = tests.AddTest("with invalid data", args{
		data:           []rune{'a', '\r'},
		tab_size:       3,
		expected_err:   "after '\\r': expected '\\n', got nothing instead",
		expected_slice: []rune{'a', '\r'},
	})

	_ = tests.AddTest("no tab size", args{
		data:           []rune{'t', 'e', 's', 't'},
		tab_size:       0,
		expected_err:   "(BadParameter) Parameter \"tab_size\" must be positive, got 0 instead",
		expected_slice: []rune{'t', 'e', 's', 't'},
	})

	_ = tests.AddTest("with tabs", args{
		data:           []rune{'t', 'e', 's', 't', '\t', '\t', 't', 'e', 's', 't'},
		tab_size:       3,
		expected_err:   "",
		expected_slice: []rune{'t', 'e', 's', 't', ' ', ' ', ' ', ' ', ' ', ' ', 't', 'e', 's', 't'},
	})

	_ = tests.Run(t)
}

// TestRepeat tests the Repeat function.
func TestRepeat(t *testing.T) {
	type args struct {
		char     rune
		count    int
		expected []rune
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			chars := Repeat(args.char, args.count)

			ok := slices.Equal(chars, args.expected)
			if !ok {
				t.Errorf("expected %v, got %v", args.expected, chars)
			}
		}
	})

	_ = tests.AddTest("with negative count", args{
		char:     'a',
		count:    -1,
		expected: []rune{},
	})

	_ = tests.AddTest("with zero count", args{
		char:     'a',
		count:    0,
		expected: []rune{},
	})

	_ = tests.AddTest("with positive count", args{
		char:     'a',
		count:    1,
		expected: []rune{'a'},
	})

	_ = tests.Run(t)
}
