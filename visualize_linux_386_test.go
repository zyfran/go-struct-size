package struct_size

import (
	"bytes"
	"testing"
)

func TestVisualize(t *testing.T) {
	items := [...]struct {
		Interface interface{}
		Result    []byte
	}{
		{Interface: nil, Result: []byte(`sizeof(nil)=8
`)},
		{Interface: [...]string{"1", "2"}, Result: []byte(`sizeof([2]string)=16
`)},
		{Interface: [...]bool{true, false}, Result: []byte(`sizeof([2]bool)=2
`)},
		{Interface: [...]interface{}{true, nil, true}, Result: []byte(`sizeof([3]interface {})=24
`)},
		{Interface: testBad{}, Result: []byte(`sizeof(struct_size.testBad)=28 with alignment=4
    test  bool   [x][ ]
    test1 uint16       [x][x]
    test2 uint16 [x][x][ ][ ]
    Val1  int32  [x][x][x][x]
    Val2  int16  [x][x][ ][ ]
    ID    string [x][x][x][x]
                 [x][x][x][x]
    Val   int32  [x][x][x][x]
`)},
		{Interface: testGood{}, Result: []byte(`sizeof(struct_size.testGood)=24 with alignment=4
    ID    string [x][x][x][x]
                 [x][x][x][x]
    Val   int32  [x][x][x][x]
    Val1  int32  [x][x][x][x]
    Val2  int16  [x][x]
    test1 uint16       [x][x]
    test2 uint16 [x][x]
    test  bool         [x][ ]
`)},
		{Interface: test1{}, Result: []byte(`sizeof(struct_size.test1)=32 with alignment=4
   ~testGood struct_size.testGood [x][x][x][x]
                                  [x][x][x][x]
                                  [x][x][x][x]
                                  [x][x][x][x]
                                  [x][x][x][x]
                                  [x][x][x][x]
    d        *int                 [x][x][x][x]
    f        bool                 [x][ ][ ][ ]
`)},
		{Interface: test2{}, Result: []byte(`sizeof(struct_size.test2)=16 with alignment=4
   ~testGood *struct_size.testGood [x][x][x][x]
    f        int                   [x][x][x][x]
    a        interface {}          [x][x][x][x]
                                   [x][x][x][x]
`)},
		{Interface: myStruct{}, Result: []byte(`sizeof(struct_size.myStruct)=20 with alignment=4
    myBool  bool    [x][ ][ ][ ]
    myFloat float64 [x][x][x][x]
                    [x][x][x][x]
    myInt   int32   [x][x][x][x]
    Int     int16   [x][x][ ][ ]
`)},
		{Interface: myStructOptimized1{}, Result: []byte(`sizeof(struct_size.myStructOptimized1)=16 with alignment=4
    myFloat float64 [x][x][x][x]
                    [x][x][x][x]
    myInt   int32   [x][x][x][x]
    Int     int16   [x][x]
    myBool  bool          [x][ ]
`)},
		{Interface: myStructOptimized2{}, Result: []byte(`sizeof(struct_size.myStructOptimized2)=16 with alignment=4
    myFloat float64 [x][x][x][x]
                    [x][x][x][x]
    myInt   int32   [x][x][x][x]
    myBool  bool    [x][ ]
    Int     int16         [x][x]
`)},
		{Interface: struct {
			a uint8
			b bool
		}{}, Result: []byte(`sizeof(struct { a uint8; b bool })=2 with alignment=1
    a uint8 [x]
    b bool  [x]
`)},
		{Interface: struct {
			a string
			b struct {
				a uint8
				b bool
				c bool
				d bool
				e bool
				f bool
				g bool
			}
			c int16
		}{}, Result: []byte(`sizeof(struct { a string; b struct { a uint8; b bool; c bool; d bool; e bool; f bool; g bool }; c int16 })=20 with alignment=4
    a string                                                             [x][x][x][x]
                                                                         [x][x][x][x]
    b struct { a uint8; b bool; c bool; d bool; e bool; f bool; g bool } [x][x][x][x]
                                                                         [x][x][x][ ]
    c int16                                                              [x][x][ ][ ]
`)},
		{Interface: struct {
			a string
			b struct {
				a uint8
				b bool
				c bool
				d bool
				e bool
			}
			c int16
		}{}, Result: []byte(`sizeof(struct { a string; b struct { a uint8; b bool; c bool; d bool; e bool }; c int16 })=16 with alignment=4
    a string                                             [x][x][x][x]
                                                         [x][x][x][x]
    b struct { a uint8; b bool; c bool; d bool; e bool } [x][x][x][x]
                                                         [x][ ]
    c int16                                                    [x][x]
`)},
		{Interface: struct {
			a string
			c int16
			b struct {
				a uint8
				b bool
				c bool
				d bool
				e bool
				f bool
				g bool
			}
		}{}, Result: []byte(`sizeof(struct { a string; c int16; b struct { a uint8; b bool; c bool; d bool; e bool; f bool; g bool } })=20 with alignment=4
    a string                                                             [x][x][x][x]
                                                                         [x][x][x][x]
    c int16                                                              [x][x]
    b struct { a uint8; b bool; c bool; d bool; e bool; f bool; g bool }       [x][x]
                                                                         [x][x][x][x]
                                                                         [x][ ][ ][ ]
`)},
		{Interface: struct {
			a string
			c int16
			b struct {
				a uint8
				b bool
				c bool
				d bool
				e bool
			}
		}{}, Result: []byte(`sizeof(struct { a string; c int16; b struct { a uint8; b bool; c bool; d bool; e bool } })=16 with alignment=4
    a string                                             [x][x][x][x]
                                                         [x][x][x][x]
    c int16                                              [x][x]
    b struct { a uint8; b bool; c bool; d bool; e bool }       [x][x]
                                                         [x][x][x][ ]
`)},
	}

	for _, item := range items {
		result := Visualize(item.Interface)
		if !bytes.Equal(item.Result, result) {
			t.Errorf(`Interface %v
Expected:
%s
Actual:
%s`,
				item.Interface,
				item.Result,
				result,
			)
		}
	}
}
