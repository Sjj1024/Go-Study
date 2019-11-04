package main
import "fmt"

// init函数是在main函数执行之前执行的，做初始化使用
/*
1、如果一个文件中有全局变量，init函数和main函数，则执行顺序是先执行全局变量，init，main
2、go支持匿名函数，匿名函数就是没有名字的函数，或者是把函数交给一个变量



*/

var num = test()

func test() int {
	fmt.Println("test函数生成全局变量")  // 1
	return 100
}

func init(){
	fmt.Println("初始化函数init")  // 2
}

// 全局匿名函数的定义方式：不支持类型推到：=这种赋值
var myfunc = func (n1 int, n2 int) int {
		return n1 * n2
	}
// 或者使用var （）
// var (
// 	myfunc2 = func () {
		
// 	}
// )

func main() {
	// 初始化init使用
	fmt.Println("入口函数", num)  // 3

	// 求两个数的和，使用匿名函数的方式完成
	sum := func (n1 int, n2 int) int {
		return n1 + n2
	}(20, 30)
	fmt.Println(sum)

	// 第二种匿名函数的使用方式:可以在main中使用这种方式定义函数
	sub := func(n1 int, n2 int) int{
		return n1 - n2
	}
	res := sub(30, 40)
	fmt.Println(res)

	// 全局匿名函数
	res2 := myfunc(5, 6)
	fmt.Println(res2)
}






test函数生成全局变量
初始化函数init
入口函数 100
50
-10
30
[Finished in 1.0s]
