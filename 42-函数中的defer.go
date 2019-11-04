package main
import "fmt"


// 函数中的defer
/*
在函数中，程序员经常需要创建资源（比如数据库连接，文件句柄，锁等）为了在函数执行完毕后，
及时的释放资源，Go的设计者提供defer延时机制。

*/

func sum(n1, n2 int) int{
	// defer 会将后面的内容压入defer栈中，保持先进后出的原则
	defer fmt.Println("ok1, n1======")
	defer fmt.Println("ok2, n2========")

	res := n1 + n2
	fmt.Println(res)
	return res
}

func main() {
	// 
	res := sum(10, 20)
	fmt.Println(res)
}






30
ok2, n2========
ok1, n1======
30
[Finished in 1.2s]
