package main
import (
	"fmt"

)

/*
1.算术运算符：加减乘除/取余，取模%,如果参与运算的数都是整数，得到的结果还是整数。


*/


func main() {
	// 除法运算中如果参与运算的都是整数，得到的也是整数，
	num1 := 10 / 4
	fmt.Println(num1)

	// 带小数的得到的就是小数
	num2 := 10.0/4
	fmt.Println(num2)

	// 取模运算：% 特殊用法公式：a%b = a-a/b*b(python里也通用)
	num3 := -10 / 3
	fmt.Println(num3)

	// 加加++和剪剪--
	num4 := 100
	num4 ++  // 自增1，相当于num4 = num4 +1
	num4 --  // 自减1，相当于num4 = num4 -1

	fmt.Println(num4)


}









2
2.5
-3
100
[Finished in 0.7s]







