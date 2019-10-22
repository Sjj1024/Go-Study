package main
import "fmt"


// 基本数据类型的默认值,
// int 默认为0，float默认为0.0，布尔默认为false，str默认为""


func main() {
	var num1 int
	var num2 float32
	var num3 bool
	var num4 string

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	// 按照原始值输出就用%v
	fmt.Printf("%v \n", num2)
}




0
0
false

0 
[Finished in 0.8s]
