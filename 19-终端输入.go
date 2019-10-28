package main
import "fmt"


// 终端输入方法
// fmt.Scanln():服务一行数据
// fmt.Scanf():


func main() {
	// 使用Scanln获取姓名，年龄，薪水，那女
	var name string
	var age int
	var sal float32
	var gendal bool

	// fmt.Println("请输入姓名:")
	// // 当程序执行到这里的时候，会等待用户的输入
	// fmt.Scanln(&name) // 传地址才能直接影响外面的name值

	// fmt.Println("请输入年龄:")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入性别真假:")
	// fmt.Scanln(&gendal)


	// 第二种方式
	fmt.Println("请输入：姓名 年龄 薪水 男女")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &gendal)
	fmt.Println(name, age, sal, gendal)
}




