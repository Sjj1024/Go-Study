package main
import "fmt"
import "unsafe"


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

	// 如何查看程序中某个变量的占用字节大小和数据类型
	// unsafe.Sizeof(num1) 是unsafe的一个函数，可以返回num1变量占用的字节数
	fmt.Printf("数据类型：%T 占用的字节数：%d", num1, unsafe.Sizeof(num1))
}






-129
127
nihaoya,niduoda int8数据类型：int 占用的字节数：8[Finished in 0.8s]
