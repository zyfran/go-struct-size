package struct_size

import (
	"testing"
)

func TestCheck(t *testing.T) {
	items := [...]struct {
		Interface     interface{}
		InterfaceSize uint
		OptimalSize   uint
		Optimal       bool
	}{
		{Interface: nil, InterfaceSize: 8, OptimalSize: 8, Optimal: true},
		{Interface: true, InterfaceSize: 1, OptimalSize: 1, Optimal: true},
		{Interface: test1{}, InterfaceSize: 32, OptimalSize: 32, Optimal: true},
		{Interface: test2{}, InterfaceSize: 16, OptimalSize: 16, Optimal: true},
		{Interface: testBad{}, InterfaceSize: 28, OptimalSize: 24, Optimal: false},
		{Interface: testGood{}, InterfaceSize: 24, OptimalSize: 24, Optimal: true},
		{Interface: myStruct{}, InterfaceSize: 20, OptimalSize: 16, Optimal: false},
		{Interface: myStructOptimized1{}, InterfaceSize: 16, OptimalSize: 16, Optimal: true},
		{Interface: myStructOptimized2{}, InterfaceSize: 16, OptimalSize: 16, Optimal: true},
		{Interface: struct {
			a uint8
			b bool
		}{}, InterfaceSize: 2, OptimalSize: 2, Optimal: true},
	}

	for _, item := range items {
		current, best, optimal := CheckSize(item.Interface)
		if item.InterfaceSize != current || item.OptimalSize != best || item.Optimal != optimal {
			t.Errorf(
				"Interface %v Expected: %d, %d, %t Actual: %d, %d, %t",
				item.Interface,
				item.InterfaceSize, item.OptimalSize, item.Optimal,
				current, best, optimal,
			)
		}
	}
}
