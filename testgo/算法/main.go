package main

import "fmt"

//二维数组顺时针旋转90度
//var arry = [3][3]int {{1,2,3},
//	                    {4,5,6},
//	                    {7,8,9}}
//输出  7 4 1 ,
//     8 5 2 ,
//     9 6 3
func rotate(matrix *[3][3]int ) {
	n := len(matrix)
	for i :=0;i<n/2;i++{
		for j := i;j<n-1-i;j++{
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}

//自已写的数组旋转90度，把每一位的对应位置找出来做好位置比较，找出共同点再计算
func rotate2(arry [3][3]int) {
	arryBak := arry
	l := len(arry)

	for i:=0;i<l;i++{
		for j :=0;j<len(arry[0]) ;j++{
			arryBak[i][j] = arry[l-1-j][i]
		}
	}
	fmt.Println(arryBak)
}


func main(){
	var arry = [3][3]int {{1,2,3},
		                    {4,5,6},
		                    {7,8,9}}
	rotate2(arry)
}
