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





package main
import (
	"fmt"
	"utils"
)


// 包的本质就是创建不同的文件夹存放不同的方法
// go的每一个文件都属于一个包，也就是说go是以包来区分文件的
// 控制函数，变量的使用范围，
// 调用包的方法：import 包路径，使用的时候包.方法名
// 相同的包下，不允许有相同的方法名，哪怕是在不同的文件中，

// 如果要将这个文件编译成可执行文件，就需要将这个文件声明为main，即package main
// 否则无法编译成可执行文件，如果你写一个库，包名可以自定义


func main() {
	// 包的介绍和使用
	// 使用函数
	res := utils.Cal(5, "+", 5)
	fmt.Println("res=", res)

}
