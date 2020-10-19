package lib

import (
	"fmt"
	"testing"
)

func TestPanand_GetDiskInfo(t *testing.T) {
	type Test23 struct {
		L  int `autoLen:"Cs"`
		Cs []int
	}
	type TestL struct {
		L   int `autoLen:"BB"`
		L2  int `autoLen:"BBc"`
		L3  int `autoLen:"Cs"`
		L4  int `autoLen:"OO"`
		Cs  []int
		LL  Test23
		BB  [3]Test23
		BBc *[]*Test23
		OO  map[string]*Test23
	}
	t23 := Test23{
		L:  0,
		Cs: []int{1, 2, 3, 48},
	}
	t24 := Test23{
		L:  0,
		Cs: []int{1, 2, 39},
	}
	t25 := Test23{
		L:  0,
		Cs: []int{1, 2},
	}
	var dw = []*Test23{&t23, &t23, &t24, &t25}
	var IOw = map[string]*Test23{
		"a": &t23,
		"b": &t25,
	}
	tl := TestL{
		L:   0,
		L2:  0,
		Cs:  []int{1, 2, 4},
		LL:  t23,
		BB:  [3]Test23{t23, t23, {Cs: []int{6, 5, 8, 1}}},
		BBc: &dw,
		OO:  IOw,
	}
	fmt.Printf("%p\n", tl.BBc)
	fmt.Printf("%p\n", dw)
	fmt.Println(*dw[0])
	fmt.Println((*tl.BBc)[0])

	fmt.Println(tl)
	fmt.Println(AutoSetLength(tl, "autoLen"))
	fmt.Println(tl)
	fmt.Println(AutoSetLength(&tl, "autoLen"))
	fmt.Println(tl)
	fmt.Println(tl.L2)
	for _, v := range tl.BB {
		fmt.Println(v)
	}
	fmt.Printf("%p\n", tl.BBc)
	fmt.Printf("%p\n", dw)
	(*tl.BBc)[0] = nil
	fmt.Println(dw[0])
}
