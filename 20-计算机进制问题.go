package main
import "fmt"

// 进制：二进制：满二进一，十进制：满十进一，八进制：满八进一，0-7表示，以数字0开头表示，
// 十六进制：满16进一，0-9和A-F表示，字母不区分大小写，以0X或0x表示，
// 



func main() {
	// 二进制输入可以使用%b输出
	var num = 8
	fmt.Printf("%b \n", num)

	// 八进制可以直接输出,满8进一
	var num2 int = 014
	fmt.Println(num2)

	// 十六进制：以0x开头，满16进一
	var num3 int = 0x11
	fmt.Println(num3)


}


1000 
12
17
[Finished in 0.6s]
