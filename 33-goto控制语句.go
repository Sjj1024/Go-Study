package main
import "fmt"


// goto语句可以无条件的转移到程序指定的行
// goto语句通常与条件语句配合使用，可以实现条件转移，跳出循环体等功能
// 在go程序设计中一般不推荐使用goto语句，以免造成程序流程的混乱，使理解和调试程序都产生困难
// 基本语法：goto 标签
// goto label（标签名字可以随便取）
// return 使用在方法或者函数中，表示跳出所在的方法或函数，在讲解函数的时候，会详细讲解


func main() {
	// goto的基本使用
	goto lable1
	fmt.Println("这是第一句")

	lable1:
	for i := 1; i < 5; i++{
		fmt.Println(i)
	}
}


1
2
3
4
[Finished in 1.2s]
