package CodeTest

import "errors"

func SetMain(ina, inb int) error {
	if ina == 0 {
		return errors.New("is zero")
	}
	if inb == 0 {
		return errors.New("is zero")
	}

	var k int
	k = 10

	if (ina + inb) == k {
		return nil
	}else{
		return errors.New("sum error")
	}

}
