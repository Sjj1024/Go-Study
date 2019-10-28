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

	// 加加++和剪剪--， 只能独立使用，不能赋值给别人使用，也不能用在if中，规定死的
	// 而且只有后面的加加，减减（python中没有这些加加或减减操作）
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







// 运算符练习题
package main
import (
	"fmt"
)


// 加入还有97天放假，问：多少个星期零几天
// 定义一个变量保存华氏温度，华氏温度转换为摄氏维度的公式：5/9*(华氏温度-100)，请求出华氏温度对应的摄氏维度
func main() {
	// 计算放假天数
	days := 97
	monday := 97 / 7
	day := days % 7
	fmt.Println("还有", monday, day)
	fmt.Printf("%d星期%d天\n", monday,day)


	// 定义一个变量求维度
	huashi := 95.2
	sheshi := 5.0 / 9 * (huashi - 100)
	fmt.Println(sheshi)

}



还有 13 6
13星期6天
-2.666666666666665
[Finished in 0.7s]





