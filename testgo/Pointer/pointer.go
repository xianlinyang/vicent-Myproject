package main

// beego go常用框架
// &取地址 *解引用
// 值在传递给函数或者方法的时候会被复制
// 通道、函数、方法、映射、切片是引用变量，它们持有的都是引用，也即保存指针的变量。

import (
	"fmt"
)

// 解构赋值，无需使用第三方临时变量交换
// 声明三个持有指针的变量
func swap1(x, y, p *int) {
	if *x > *y {
		*x, *y = *y, *x
	}
	*p = *x * *y
}

// 持有值变量交换
// 类似于C语言的函数声明
// 支持多值返回
func swap2(x, y int) (int, int, int) {
	if x > y {
		x, y = y, x
	}
	return x, y, x * y
}

func main() {
	// 类型推导声明赋值
	i := 10
	j := 5
	product := 0

	// 取内存地址交换
	swap1(&i, &j, &product)
	fmt.Println(i, j, product)

	a := 64
	b := 23

	// 复制值交换
	a, b, p := swap2(a, b)
	fmt.Println(a, b, p)

}