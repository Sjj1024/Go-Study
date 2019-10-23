package main
import (
	"fmt"
)


// 使用指针写一个程序，获取int变量num的地址，并显示在终端
// 将num的地址赋值给指针ptr，并通过指针ptr去修改num的值
// 指针使用细节，使用类型：*类型

func main() {
	// 定义变量num

	var num = 10

	// 使用指针获取地址：下面三种都可以
	// var ptr *int = &num
	// var ptr = &num
	ptr := &num

	// 通过指针修改num的值,那么这个指针指向的num值就改变了
	*ptr = 22

	// 显示在终端
	fmt.Println(num)
	fmt.Println(ptr)
}







22
0xc0420381d0
[Finished in 0.8s]
