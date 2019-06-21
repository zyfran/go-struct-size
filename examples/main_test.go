package main

import (
	"testing"

	"gitlab.com/zyfran/go-struct-size"
)

func TestStructures(t *testing.T) {
	items := [...]interface{}{
		myStruct{},
		myStructOptimized1{},
		myStructOptimized2{},
	}

	for _, item := range items {
		if current, best, optimal := struct_size.CheckSize(item); !optimal {
			t.Errorf(
				`Structure can be optimized from %d to %d bytes
%s`,
				current,
				best,
				struct_size.Visualize(item),
			)
		}
	}
}
