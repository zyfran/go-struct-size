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
		if currentSize, optimalSize, ok := struct_size.CheckSize(item); !ok {
			t.Errorf(
				`Structure can be optimized from %d to %d bytes
%s`,
				currentSize,
				optimalSize,
				struct_size.Visualize(item),
			)
		}
	}
}
