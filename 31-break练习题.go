package main
import "fmt"


func main() {
	// 一百以内的数求和，求出第一次大于20的当前数
	var sum int
	for i := 1; i <= 100; i++{
		sum += i
		if sum > 20{
			fmt.Println("总和是：", sum)
			fmt.Println("最后一个数是：", i)
			break
		}
	}

	// 实验验证登陆，有三次机会，如果用户名为宋江江，密码是123456，提示登陆成功，否则提示还有几次机会
	var name string
	var password string
	// fmt.Println("请输入用户名：")
	// fmt.Scanln(&name)
	// fmt.Println("请输入密码：")
	// fmt.Scanln(&password)
	name = "宋江江"
	password = "123456"
	for i := 1; i <= 3; i++{
		if name == "宋江江" && password == "123456"{
			fmt.Println("登陆成功")
			break
		}else{
			fmt.Printf("还有次%d机会 \n", 3-i)
		}
	}
}



总和是： 21
最后一个数是： 6
登陆成功
[Finished in 1.2s]

