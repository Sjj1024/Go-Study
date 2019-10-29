package main
import "fmt"


// for循环控制:

// 使用细节:循环条件返回值是一个布尔值；
// 第二种写法：循环变量初始化 for 循环条件{ 循环变量迭代 循环体}


func main() {
	// for循环使用,初始变量；for 条件控制；变量变化{循环语句}
	for i :=1; i<=10; i++ {
		fmt.Println("我爱你")
	}

	// 第二种for写法
	j := 3
	for j < 10{
		fmt.Println("祖国")
		j++
	}
	 
	// 第三种：死循环,通常配合break语句使用
	// for {
	// 	fmt.Println("真的")
	// }
	n := 5
	for ;;{
		if n<10{
			fmt.Println("zhende")
		}else{
			break
		}
		n++
		
	}
}





我爱你
我爱你
我爱你
我爱你
我爱你
我爱你
我爱你
我爱你
我爱你
我爱你
祖国
祖国
祖国
祖国
祖国
祖国
祖国
zhende
zhende
zhende
zhende
zhende
[Finished in 0.6s]

