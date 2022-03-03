package main

import (
	"fmt"
	"math/big"
)

func main() {
	price := new(big.Int).SetInt64(int64(-100))
	total := new(big.Int).SetInt64(int64(101))
	//k := new(big.Int).Sub(total,price) //total-price
	//k = new(big.Int).Add(total,price) //total+price

	//expectedDebt := new(big.Int).Neg(price)
	//fmt.Println(price)
	//expectedDebt :=price.Neg(total) //把total变成负数
	expectedDebt := total.Cmp(price) //0是等于 1是大于 -1是小于
	//expectedDebt := total.Sub(total,price)

	//f := new(big.Int).Mul(price,new(big.Int).SetInt64(10))
	//
	fmt.Println(new(big.Int).Neg(price))
	fmt.Println(price)
	fmt.Println(expectedDebt)
}
