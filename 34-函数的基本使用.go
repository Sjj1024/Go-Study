package main
import "fmt"


// 函数的定义和使用
/*
基本语法：func 函数名 (形参列表) (返回值列表){
	执行语句
	return 返回值列表
}
形参列表：表示函数的输入
函数中的语句：表示未来实现某一功能代码块
函数可以有返回值，也可以没有
*/

func cal(num1 int, num2 string, num3 int) int{
	var sum int
	switch num2{
	case "+":
		sum = num1 + num3
	case "-":
		sum = num1 - num3
	case "*":
		sum = num1 * num3
	case "/":
		sum = num1 / num3
	default:
		fmt.Println("输入有误！")
	}
}


func main() {
	// 使用函数
	res := cal(5, +, 5)
	fmt.Println("res=", res)
}







res= 10
[Finished in 1.3s]
