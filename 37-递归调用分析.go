package main
import "fmt"


// 递归调用流程
func Gui(num int){

	if num > 4{
		// 当num--变化后，用到自身的地方就都变了,例如传进来的num是6，执行num--后，就变成5了，后面打印的时候也是5
		num--
		// 当递归调用自身时，后面的代码先不执行了。直到调用完自身后，再执行后面的代码
		Gui(num)
	}
	fmt.Println(num)
}

func main() {
	//  递归调用流程分析
	Gui(6)
}




4
4
5
[Finished in 1.0s]
