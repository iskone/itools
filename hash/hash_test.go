package hash

import (
	"testing"
)

func TestHash_Diff(t *testing.T) {
	var h1 Hash
	h1.Set("1","1")
	h1.Set("2","2")
	h1.Set("3","3")
	h1.Set("4","4")
	t.Log("RAW  :",h1)
	h2:=NewHash()
	h2.Set("2","2")
	h2.Set("3","h")
	t.Log("TOB  :",h2)
	a,b:=h1.Diff(h2,func(i1 interface{},i2 interface{})bool {
		if i1.(string)==i2.(string) {
			return true
		}
		return false
	})
	if a.Len()!=3 {
		t.Fail()
	}
	t.Log("lack :",a)
	if b.Len()!=1 {
		t.Fail()
	}
	t.Log("Added:",b)
}
