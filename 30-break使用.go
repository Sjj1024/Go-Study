package main
import (
	"fmt"
	"math/rand"
	"time"
)


// 

func main() {
	var count int 
	for {
		// 生成一个随机数
		r := rand.New(rand.NewSource(time.Now().UnixNano()))  // 先使用rand生成一个种子，这个种子是由1970年到现在的秒数组成的
		n := r.Intn(100) + 1  // 再有这个rand种子生成随机数
		count++
		fmt.Printf("生成的随机数是：%d \n", n)
		if n == 99{
			fmt.Printf("一共随机了：%d次 \n", count)
			break  // 退出for循环
		}
	}
}



生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：14 
生成的随机数是：99 
一共随机了：208次 
[Finished in 1.6s]
