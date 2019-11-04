package main
import "fmt"


// 交换两个变量的值
func swap(n1 int, n2 int){
	n1, n2 = n2, n1
	fmt.Println(n1, n2)
}

func main() {
	// 
	swap(1, 6)

}






6 1
[Finished in 0.9s]
