package main
import "fmt"


// continue也可以跟标签，跳转指定的循环
// 结束本次循环，继续下次循环


func main() {
	// continue的课堂练习
	for i := 1; i < 10; i++{
		if i == 6{
			continue
		}
		fmt.Println(i)
	}
}




1
2
3
4
5
7
8
9
[Finished in 1.3s]
