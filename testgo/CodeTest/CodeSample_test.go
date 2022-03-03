package CodeTest

import (
	"testing"
)

func TestSetMain(t *testing.T){
	k := 6
	d := 4
	err := SetMain(0,d)
	if err == nil{
		t.Fatal(err)
	}
	err = SetMain(1,0)
	if err == nil{
		t.Fatal(err)
	}
	err = SetMain(k,d)
	if err != nil{
		t.Fatal(err)
	}
	err = SetMain(k,11)
	if err == nil{
		t.Fatal(err)
	}
}
