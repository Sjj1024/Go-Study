package main
import "fmt"


// go在不同类型变量之间赋值时需要显式转换，也就是说数据类型不支持自动换换
// 基本类型转换语法：数据类型T（要被转换的变量V）  就可以将V转换成T类型


func main() {
	// 类型转换使用
	var num1 int = 10
	str1 := string(num1)
	float1 := float32(num1)
	num2 := int(float1)

	// 输出方式
	fmt.Println("nihao")
	fmt.Println(num1)
	fmt.Println(str1)
	fmt.Println(float1)
	fmt.Println(num2)
	fmt.Printf("%T \n", str1)
	fmt.Printf("%T \n", float1)
}







nihao
10


10
10
string 
float32 
[Finished in 0.8s]

