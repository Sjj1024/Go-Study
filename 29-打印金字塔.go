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

