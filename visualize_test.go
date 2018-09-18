package struct_size

import (
	"testing"
)

type (
	testGood struct {
		ID    string
		Val   int32
		Val1  int32
		test1 uint16
		test2 uint16
		test  bool
	}
	testBad struct {
		test  bool
		test1 uint16
		test2 uint16
		Val1  int32
		ID    string
		Val   int32
	}

	test1 struct {
		testGood
		d *int
		f bool
	}

	test2 struct {
		*testGood
		f int
		a interface{}
	}
)

func TestVisualize(t *testing.T) {
	t.Logf(
		`
%s`,
		Visualize(
			nil,
			[...]string{"1", "2"},
			[...]bool{true, false},
			[...]interface{}{true, nil, true},
			testBad{},
			testGood{},
			test1{},
			test2{},
		),
	)
}

func benchmarkVisualize(b *testing.B, v interface{}) {
	b.SetBytes(2)
	for i := 0; i < b.N; i++ {
		Visualize(v)
	}
}

func BenchmarkVisualizeNil(b *testing.B) {
	benchmarkVisualize(b, nil)
}

func BenchmarkVisualizeBase(b *testing.B) {
	benchmarkVisualize(b, [...]bool{true, false})
}

func BenchmarkVisualizeStructure(b *testing.B) {
	benchmarkVisualize(b, testGood{})
}
