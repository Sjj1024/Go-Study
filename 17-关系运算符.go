package main
import (
	"fmt"
)


// 关系运算符的使用
/*
运算结构都是布尔类型
逻辑运算符：多条件判断：&& || ! 与或非


*/


func main() {
	// 关系运算符的使用

	n1 := 2
	n2 := 6

	res1 := n2 > n1
	res2 := n2 < n1
	res3 := res1 || ! res2

	fmt.Println(n2 > n1 && res3)

	if n2 > n1 && res3 {
		fmt.Println("你猜")
	}

	if n2 > n1 || res3 {
		fmt.Println("你猜2")
	}

	if n2 > n1 && ! res3 {
		fmt.Println("你猜3")
	}
}





true
你猜
你猜2
[Finished in 0.8s]
