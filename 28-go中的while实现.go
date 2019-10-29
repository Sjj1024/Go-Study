package main
import "fmt"


// go中没有while和dowhile语法，可以使用for循环实现while
// for{if 循环条件{break} 循环体语句}


func main() {
	// 使用for实现while循环
	var num = 1
	for{
		if num > 10{
			break
		}
		num++
		fmt.Println("woaini")
	}

	// 使用for实现do while
	var num2 = 5
	for{
		fmt.Println("我爱你")
		num2++
		if num2 > 10{
			break  // break就是跳出for循环
		}
	}

}




woaini
woaini
woaini
woaini
woaini
woaini
woaini
woaini
woaini
woaini
我爱你
我爱你
我爱你
我爱你
我爱你
我爱你
[Finished in 1.1s]
