package main
import (
	"fmt"
	"strconv"
)

// string类型转基础类型,需要使用strconv函数


func main() {
	
	// string类型转基础类型
	str1 := "true"
	// 开始转换,_下划线表示忽略的意思，因为这个ParseBool会返回两个值，一个bool，一个错误
	bool1,_ := strconv.ParseBool(str1)
	// 打印输出
	fmt.Println(bool1)
	fmt.Printf("%T \n", bool1)


	// 将str转成int类型数据
	str2 := "1243"
	// strconv.ParseInt(str2, 10, 0)中10表示十进制。0表示转成int，可以是8,16,32等
	int2,_ := strconv.ParseInt(str2, 10, 0)
	fmt.Println(int2)


	// 将str转成float类型
	str3 := "45.2"
	// strconv.ParseFloat(str3, 64)其中64表示float64
	float3,_ := strconv.ParseFloat(str3, 64)
	fmt.Println(float3)
}



true
bool 
1243
45.2
[Finished in 0.8s]


