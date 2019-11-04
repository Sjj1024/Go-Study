package main
import "fmt"


// 递归调用流程
func Gui(num int){

	if num > 4{
		// 当num--变化后，用到自身的地方就都变了
		num--
		// 当递归调用自身时，后面的代码先不执行了。直到调用完自身后，再执行后面的代码
		Gui(num)
	}
	fmt.Println(num)
}


// 递归调用流程2使用
func Gui2(num int){
	if num > 4{
		// 当num--变化后，用到自身的地方就都变了
		// 递归必须向退出递归的条件逼近，否则就是无限递归，死鬼了
		num--
		// 当递归调用自身时，后面的代码先不执行了。直到调用完自身后，再执行后面的代码
		Gui(num)
	} else {
		fmt.Println(num)
	}
	
}

func main() {
	//  递归调用流程分析
	// Gui(6)

	fmt.Println("----------------")

	// 递归调用2使用
	Gui2(6)
}




4
4
5
[Finished in 1.0s]
