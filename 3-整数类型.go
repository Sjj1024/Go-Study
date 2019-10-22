package main
import "fmt"


// 整数类型数据
// int8 -128---127之间
// int16 -2de15---2de15-1之间
// int32,int68


func main() {
	num1 := -129
	// 使用int8，超过范围就会包超出范围
	var num2 int8 = 127


	fmt.Println(num1)
	fmt.Println(num2)
}





-129
127
[Finished in 1.0s]
