package main
import "fmt"

func main(){
	var num1 int
	fmt.Println("请输入一个整数：")
	fmt.Scanln(&num1)
	// switch的使用注意事项：
	// switch和case后面可以是一个常量，变量，或者是一个有返回值的表达式
	// case后面各个表达式的值要和switch后的数据类型保持一致
	// case后面可以跟多个表达式：使用，号隔开
	// case后面的数据不能重复
	// case后面不需要带break，匹配到后就执行代码，然后退出
	// default可以没有，
	switch num1 {
	case 1:
		fmt.Println("这是周一！")
	case 2, 4, 6, 8:
		fmt.Println("这是周二！")
	default:
		fmt.Println("输入有误！")

	}
}



请输入一个整数：
4
这是周二！
