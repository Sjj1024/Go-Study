package main
import "fmt"

// 字符串是由字节组成的:
// 字符常量是用单引号括起来的单个字符，例如'a','1','我'
// 允许使用\表示转义字符，go中的编码格式是utf-8编码的，再也不会遇到中文乱码问题
// 英文是一个字节保存，汉字占3个字节
// 字符类型相当于一个整数，可以进行数学运算。



func main() {
	// 保存一个字符,如果使用byte保存，只能使用单引号，双引号表示str类型，byte表示范围是0-255
	var alpha byte = 'a'
	// 想用byte输出对应的字符，可以使用%c格式化输出
	// 如果保存的字节码超过了byte的范围，可以使用int类型曹村这个字节码值
	var signle int = '我'
	// 字符是可以进行数学运算的
	var sum = 10+signle

	fmt.Println(alpha)
	fmt.Printf("%c \n",alpha)
	fmt.Println(signle)
	fmt.Printf("%c \n",signle)
	fmt.Println(sum)

}






97
a 
25105
我 
25115
[Finished in 0.7s]
