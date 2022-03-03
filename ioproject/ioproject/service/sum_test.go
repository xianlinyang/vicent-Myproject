package service_test

import (
	"fmt"
	"testing"
)

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}
func TestFunc(t *testing.T) {

}
