package main
import "fmt"

// 闭包的理解
/*
闭包的介绍和案例演示
闭包：闭包就是一个函数和与其相关的引用环境组合的一个整体
注意：返回的是一个匿名函数，匿名函数使用到了外函数的n，因此匿名函数就和n形成了一个整体，构成闭包
可以理解为：闭包是一个类，内函数是一个方法，n是一个属性
当我们反复调用f的时候，因为n只初始化一次，因此每调用一次就进行累加
我们要分析出

*/

func Addnum() func(int) int {
	var n int = 10  // 只会初始化一次
	return func (x int) int {
		n = n + x
		return n
	}
}

func main() {
	// 闭包的理解和使用
	f := Addnum()
	fmt.Println(f(1))  // 11
	fmt.Println(f(1))  // 12
}





11
12
[Finished in 0.9s]
