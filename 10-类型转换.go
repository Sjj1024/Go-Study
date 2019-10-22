package main
import "fmt"


// go在不同类型变量之间赋值时需要显式转换，也就是说数据类型不支持自动转换
// 基本类型转换语法：数据类型T（要被转换的变量V）  就可以将V转换成T类型
// 数据类型转换可以表示范围小的到范围大的，也可以是范围大的到范围小的
// 被转换的是变量存储的数据，变量本身的数据类型并没有变化
// 如果从大的转换到小的，编译时不会出错，只是转换后的结果按溢出处理
// 浮点类型不能转换成字符串类型！

func main() {
	// 类型转换使用
	var num1 int = 10
	str1 := string(num1)
	float1 := float32(num1)
	num2 := int(float1)
	// 浮点类型转换
	float2 := 10.56464331651
	int1 := int(float2)
	str2 := string(float2)

	// 输出方式
	fmt.Println("nihao")
	fmt.Println(num1)
	fmt.Println(str1)
	fmt.Println(float1)
	fmt.Println(num2)
	fmt.Println(float2)
	fmt.Println(int1)
	fmt.Println(str2)
	fmt.Printf("%T \n", str1)
	fmt.Printf("%T \n", float1)
}





nihao
10


10
10
10.56464331651
10


string 
float32 
[Finished in 0.7s]



