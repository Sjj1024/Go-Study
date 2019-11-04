package main
import "fmt"

// init函数是在main函数执行之前执行的，做初始化使用
/*
1、如果一个文件中有全局变量，init函数和main函数，则执行顺序是先执行全局变量，init，main
2、



*/

var num = test()


func test() int {
	fmt.Println("test函数生成全局变量")  // 1
	return 100
}

func init(){
	fmt.Println("初始化函数init")  // 2
}

func main() {
	// 初始化init使用
	fmt.Println("入口函数", num)  // 3
}






test函数生成全局变量
初始化函数init
入口函数 100
[Finished in 0.9s]
