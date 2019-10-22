package main
import "fmt"


// 浮点类型的使用，即小数类型使用介绍
// 小数类型分类：单精度32位占用4字节存储空间，双精度64位占8字节空间，都可以表示负数，64表示数据更多
// 32位精度的话可能会造成精度损失，go中浮点型默认声明为float64类型

func main() {
	num1 := 9.4
	num2 := -9.34342323252

	// 保存数值高的请用大的
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Printf("num的数据类型%T",num1)
}


9.4
-9.34342323252
num的数据类型float64[Finished in 0.6s]
