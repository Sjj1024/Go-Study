package main
import "fmt"


// 使用for循环打印金字塔


func main() {
	// 打印直角金字塔
	var n = 6
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("*")
		}
		fmt.Println("")
	} 

	// 打印三角形金字塔
	fmt.Println("-------------------")
	var n2 = 6
	for i:=1;i<=n2;i++{
		for k:=1;k<=n2-i;k++{
			fmt.Printf(" ")
		}
		for j:=1;j<=(i*2-1);j++{
			fmt.Printf("*")
		}
		fmt.Println("")
	}

	// 水仙花数
	fmt.Println("---------------------")
	var num = 100
	for ;num<1000;num++{
		if 100< num && num < 1000{
			fmt.Println("这是一个三位数")
			var ge = num % 10
			var shi = num / 10 % 10
			var bai = num / 100
			var sum = ge*ge*ge + shi*shi*shi + bai*bai*bai
			fmt.Printf("个位是：%d 十位是：%d 百位是：%d 立方和是:%d \n", ge, shi, bai, sum)
			if num == sum{
				fmt.Println("恭喜，这是一个 水仙花数======================")
			}else{
				fmt.Println("不好意思，这不是一个水仙花数")
			}
		}else{
			fmt.Println("不是一个三位数")
		}
	}

}



*
**
***
****
*****
******
-------------------
     *
    ***
   *****
  *******
 *********
***********
[Finished in 1.1s]

