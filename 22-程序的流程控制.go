package main
import "fmt"

// 流程控制：顺序控制，分支控制，循环控制
// 分支控制：if else 
// 多分支：if else if else if 


// 循环控制:


func main() {
	// 单分支语句
	if age:=18; age>16{
		fmt.Println("我喜欢你")
	}

	num2 := 10
	if num2 > 10 {  // 前面的逻辑判断可以写括号，也可以不写括号，官方推荐不写
		fmt.Println("假的")
	} else {  // else只能放到这里，不然报错
		fmt.Println("真的")
	}

	// 多分支
	if num2 > 10{
		fmt.Println("1")
	} else if num2 > 20{
		fmt.Println("2")
	} else if num2 < 30{
		fmt.Println("3")
	} else{
		fmt.Println("4")
	}

	num3 := 4
	fmt.Println(num3 * 2)

	// switch分支结构:每一个case都是唯一的
	var kay int
	fmt.Println("请输入一个数字：")
	fmt.Scanln(&kay)

	switch kay{
		case 1:
			fmt.Println("这是周一!")
		case 2:
			fmt.Println("这是周二！")
		case 3:
			fmt.Println("这是周三！")
		case 4:
			fmt.Println("这是周四！")
		default:
			fmt.Println("太累了不想写了！")
	}

}




我喜欢你
真的
3
[Finished in 0.9s]
