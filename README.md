[![coverage report](https://gitlab.com/zyfran/go-struct-size/badges/master/coverage.svg)](https://gitlab.com/zyfran/go-struct-size/commits/master)
[![pipeline status](https://gitlab.com/zyfran/go-struct-size/badges/master/pipeline.svg)](https://gitlab.com/zyfran/go-struct-size/commits/master)

#### Description

#### Requirements
- go version `1.2.0+`

#### Installation
```bash
go get gitlab.com/zyfran/go-struct-size
```

#### Usage
```go
// main.go
package main

type (
	myStruct struct {
		myBool  bool    // 1 byte
		myFloat float64 // 8 bytes
		myInt   int32   // 4 bytes
	}

	myStructOptimized1 struct {
		myFloat float64 // 8 bytes
		myBool  bool    // 1 byte
		myInt   int32   // 4 bytes
	}

	myStructOptimized2 struct {
		myFloat float64 // 8 bytes
		myInt   int32   // 4 bytes
		myBool  bool    // 1 byte
	}
)

func main() {
	// do something
}
```

```go
// main_test.go
package main

import (
	"testing"

	"gitlab.com/zyfran/go-struct-size"
)

func TestStructures(t *testing.T) {
	t.Logf(
		`
%s`,
		struct_size.Visualize(
			myStruct{},
			myStructOptimized1{},
			myStructOptimized2{},
		),
	)
}
```

##### Check structures for 64-bit OS
```bash
$ GOOS=linux GOARCH=amd64 go test -v
```
```text
=== RUN   TestStructures
--- PASS: TestStructures (0.00s)
    main_test.go:9:
        sizeof(main.myStruct)=24
            myBool  bool    [x][ ][ ][ ][ ][ ][ ][ ]
            myFloat float64 [x][x][x][x][x][x][x][x]
            myInt   int32   [x][x][x][x][ ][ ][ ][ ]
        sizeof(main.myStructOptimized1)=16
            myFloat float64 [x][x][x][x][x][x][x][x]
            myBool  bool    [x][ ][ ][ ]
            myInt   int32               [x][x][x][x]
        sizeof(main.myStructOptimized2)=16
            myFloat float64 [x][x][x][x][x][x][x][x]
            myInt   int32   [x][x][x][x]
            myBool  bool                [x][ ][ ][ ]
```
##### Check structures for 32-bit OS
```bash
$ GOOS=linux GOARCH=386 go test -v
```
```text
=== RUN   TestStructures
--- PASS: TestStructures (0.00s)
    main_test.go:9:
        sizeof(main.myStruct)=16
            myBool  bool    [x][ ][ ][ ]
            myFloat float64 [x][x][x][x]
                            [x][x][x][x]
            myInt   int32   [x][x][x][x]
        sizeof(main.myStructOptimized1)=16
            myFloat float64 [x][x][x][x]
                            [x][x][x][x]
            myBool  bool    [x][ ][ ][ ]
            myInt   int32   [x][x][x][x]
        sizeof(main.myStructOptimized2)=16
            myFloat float64 [x][x][x][x]
                            [x][x][x][x]
            myInt   int32   [x][x][x][x]
            myBool  bool    [x][ ][ ][ ]
```
