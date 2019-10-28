package main
import (
	"fmt"
)


// 赋值运算符：等号，加等，除等，乘等，减等
// 运算符优先级：单目运单和赋值运算都是从右向左，其他都是从左向右
// 括号++-- > 单目运算*&等 > 算数运算符 > 移位运算 > 关系运算 > 位运算 > 逻辑运算 > 赋值运算 > 逗号
// go中不支持三元运算符，只能使用ifelse语句，python支持：num = 5 if 4>3 else 8

func main() {
	// 加减乘除等
	a := 4
	b := 5
	// 相当于b = b+a
	b += a
	fmt.Println(b)

	// 相当于b = b-a
	b -= a
	fmt.Println(b)

	// 相当于b = b-a
	b *= a
	fmt.Println(b)

	// 不使用中间变量交换两个变量的值
	b,a = a,b
	fmt.Println(a, b)

	// 求出两个数的最大值
	var max int
	if b > a {
		max = b
	} else {
		max = a
	}
	fmt.Println(max)

	

}



9
5
20
20 4
20
[Finished in 0.9s]
