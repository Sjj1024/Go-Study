package main
import (
	"fmt"
	// "math/rand"
	// "time"
)


// break使用注意事项：
// break可以通过标签指明要终止的是哪一层语句块：break 标签


func main() {
	// var count int 
	// for {
	// 	// 生成一个随机数
	// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))  // 先使用rand生成一个种子，这个种子是由1970年到现在的秒数组成的
	// 	n := r.Intn(100) + 1  // 再有这个rand种子生成随机数
	// 	count++
	// 	fmt.Printf("生成的随机数是：%d \n", n)
	// 	if n == 99{
	// 		fmt.Printf("一共随机了：%d次 \n", count)
	// 		break  // 退出for循环
	// 	}
	// }

	// 通过break终止指定标签,默认跳出最近的for循环
	// lable2: // 这是一个外层标签
	for i := 0; i < 4; i++{
		lable1:  // 这是一个标签，名字可以随意
		for j := 0; j < 10; j++{
			if j == 2{
				break lable1
			}
			fmt.Println("j=", j)
		}
	}
}









j= 0
j= 1
j= 0
j= 1
j= 0
j= 1
j= 0
j= 1
[Finished in 1.7s]
