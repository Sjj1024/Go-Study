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

// 猴子吃桃子的问题：猴子每天吃桃子的一半并多吃一个，当第十天的时候发现只有一个桃子了，
// 问最初有多少个桃子
func houzi(n int) int{
	if n == 10 {
		return 1
	}else{
		n++
		return 2*(houzi(n) +1)
	}
}

func main() {
	// 使用递归求出给你一个整数的斐波那契数列的和
	sum := feibo(4)
	fmt.Println(sum)

	// 使用求函数的值方程
	han := hanshu(2)
	fmt.Println(han)

	// 求出第一天一共有多少个桃子
	tao := houzi(1)
	fmt.Println(tao)
}



3
7
1534
[Finished in 0.8s]
