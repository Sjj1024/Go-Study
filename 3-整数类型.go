package main
import "fmt"


// 整数类型数据:有符号和无符号，有符号可以表示负值，无符号只能表示正值，int和uint的大小和系统位数有关
// fmt.Printf(可以用于格式化输出)
// int8 -128---127之间
// int16 -2de15---2de15-1之间
// int32,int68


func main() {
	num1 := -129
	// 使用int8，超过范围就会包超出范围
	var num2 int8 = 127
	// 无符号的类型uint8，表示0---255，uint16：0---2de15


	fmt.Println(num1)
	fmt.Println(num2)
	// 格式化输出的时候%T 中t就是type
	fmt.Printf("nihaoya,niduoda %T", num2)
}




-129
127
[Finished in 1.0s]
