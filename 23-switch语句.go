package main

import (
	"fmt"
)

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
	// switch后面可以不带表达式，但是case后面要带上判断
	switch num1 {
	case 1:
		fmt.Println("这是周一！")
	case 2, 4, 6, 8:
		fmt.Println("这是周二！")
	default:
		fmt.Println("输入有误！")
	}
	// switch后不带表达式的用法
	switch {
	case num1>4:
		fmt.Println("真的可以")
		fallthrough
	case num1<4:
		fmt.Println("完美！")
	case num1 ==8:
		fmt.Println(num1)
	}
	// switch中穿透fallthrough，如果在case语句后增加fallthrough，
	// 则会继续执行后一个case，也叫switch穿透，默认只穿透一个
}



请输入一个整数：
8
这是周二！
真的可以
完美！
