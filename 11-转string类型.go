package main
import (
	"fmt"
	"strconv"
)


// 数据转换成string类型
// 转换一：fmt.Sprintf（"%参数", 表达式），会返回转换后的字符串
// 格式二：使用

func main() {
	// 转化方法
	num1 := 100
	num2 := 10.42
	num3 := true
	num4 := 'h'
	var str1 string
	var str2 string


	// 将上述变量转换成string,%d是专业转换int类型的，%f是专业转换浮点型的
	// %t是专业转换布尔类型的，%c是专业转换字符类型的
	str1 = fmt.Sprintf("%d", num1)
	str1 = fmt.Sprintf("%f", num2)
	str1 = fmt.Sprintf("%t", num3)
	str1 = fmt.Sprintf("%c", num4)

	// 第二种方式：使用strconv函数
	str2 = strconv.FormatInt(int64(num1), 10)
	// 或者使用Itoa函数将int转成string
	str2 = strconv.Itoa(num1)
	// strconv.FormatFloat(num1,'f',10,64)中f表示输出格式是原格式，10表示小数点后保留几位，64表示float64类型
	// str2 = strconv.FormatFloat(num2,'f',10,64)
	// str2 = strconv.FormatBool(num3)
	// str2 = strconv.FormatInt(int64(num1), 10)

	// 打印出第一种转换方式的输出
	fmt.Println(str1)
	// %q表示输出一个引号括起来的原数据字面值
	fmt.Printf("%q \n",str1)

	// 打印出第二种转换方式的输出
	fmt.Println(str2)
	// %q表示输出一个引号括起来的原数据字面值
	fmt.Printf("%q \n",str2)

}





h
"h" 
true
"true" 
[Finished in 0.7s]


