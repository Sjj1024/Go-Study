package utils
import "fmt"



// 打包的基本结构 package 包名
// 使用的时候 import 包名
// 包内部的方法要大写才能被外部使用，该函数可以到处，不然小写的是私有的意思

// 文件包名通常和文件所在文件夹名字保持一致，一般为小写
// 当其他文件使用其他包函数或变量是，需要先引入对应的包
// package一定是在第一行的，
// 在别的文件中使用的时候，使用包名.函数名

func Cal(num1 int, num2 string, num3 int) int{
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
	return sum
}
