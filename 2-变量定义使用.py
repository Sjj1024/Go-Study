package main
import "fmt"
// 变量相当于内存中的一个房间，
func main() {
	// 先定义后赋值的方式
	var name int
	name = 19
	// 定义的同时赋值
	var age int = 18
	fmt.Println("name=",name)
	fmt.Println("age=",age)
}

name= 19
[Finished in 0.8s]
