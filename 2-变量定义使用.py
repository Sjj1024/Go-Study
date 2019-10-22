package main
import "fmt"
// 变量相当于内存中的一个房间，变量使用注意事项：
// 1.指定变量的类型，声明后不赋值的话，使用默认值，int默认值是0
// 2.根据智自行判定变量类型，类型推到
func main() {
	// 先定义后赋值的方式，var是variable，意思是变量，可以变的
	var name int
	name = 19
	// 定义的同时赋值
	var age int = 18
	// 类型推到
	var height = "178"
	// 打印输出格式
	fmt.Println("name=",name)
	fmt.Println("age=",age)
	fmt.Println("height=", height)
}






name= 19
age= 18
height= 178
[Finished in 0.7s]
