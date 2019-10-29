package main
import "fmt"


// 小写转大写
// 如果判断的具体数值不多，而且复合整数，浮点数，字符，字符串这几种类型，建议使用switch语句，简洁高效

func main() {
	// 用户输入，定义char类型的话，是var char byte
	// var alpha byte
	// fmt.Println("请输入一个小写字母：")
	// // fmt.Scanf("%c", &alpha)
	// fmt.Scanln(&alpha)
	// // end := alpha - 32
	// // fmt.Printf("转换后的大字母是：%c \n", end)
	// switch alpha{
	// 	case 'a':
	// 		fmt.Println("A")
	// 	case 'b':
	// 		fmt.Println("B")	
	// 	default:
	// 		fmt.Println("输入错误！", alpha)
	// }

	// 输入学生成绩，判断优异程度,case后面跟范围的话，switch后就不能带参数了，不然报错
	var score int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)
	switch{
	case score < 60:
		fmt.Println("不合格")
	case score < 80:
		fmt.Println("合格")
	case score < 90:
		fmt.Println("良好")
	case score <= 100:
		fmt.Println("优秀")
	default:
		fmt.Println("不能大于100")
	}

}



请输入一个小写字母：
转换后的字母是：A 
[Finished in 0.7s]

