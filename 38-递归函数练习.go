package main
import "fmt"


// 斐波那契数列
func feibo(n int) int{
	if n == 1{
		return 1
	}else if n == 2{
		return 1
	}else{
		// 返回前两个数的和
		return feibo(n-1) + feibo(n-2)
	}
}


// 求函数的值
func hanshu(n int) int{
	if n == 1{
		return 3
	}else{
		return 2*hanshu(n-1)+1
	}

}

func main() {
	// 使用递归求出给你一个整数的斐波那契数列的和
	sum := feibo(4)
	fmt.Println(sum)

	// 使用求函数的值方程
	han := hanshu(2)
	fmt.Println(han)
}




3
7
[Finished in 0.8s]
