package main
import "fmt"
// 变量相当于内存中的一个房间，变量使用注意事项：
// 1.指定变量的类型，声明后不赋值的话，使用默认值，int默认值是0
// 2.根据智自行判定变量类型，类型推到
// 3.第三种：省略var，使用：=，但是要注意等号左侧的变量不应该是已经声明过的，否则报错
// 4.多变量声明：同时声明多个变量，var n1， n2， n3 int
func main() {
	// 先定义后赋值的方式，var是variable，意思是变量，可以变的
	var name int
	name = 19
	// 定义的同时赋值
	var age int = 18
	// 类型推到
	var height = "178"
	// 海象赋值, 声明的同时赋值，并且冒号不能省略
	gender := "男孩"
	// 同时声明定义多个变量
	var n1,n2,n3 int  // 使用默认值0
	var a1,a2,a3 = 100,"xiaoming","男孩"  // 同时定义不同类型的变量
	
	// 打印输出格式
	fmt.Println("name=",name)
	fmt.Println("age=",age)
	fmt.Println("height=", height)
	fmt.Println("gender=", gender)
	fmt.Println("n1,n2,n3=",n1,n2,n3)
	fmt.Println("a1,a2,a3=",a1,a2,a3)
}




name= 19
age= 18
height= 178
gender= 男孩
n1,n2,n3= 0 0 0
a1,a2,a3= 100 xiaoming 男孩
[Finished in 0.7s]
