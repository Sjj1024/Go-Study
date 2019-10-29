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

	// 如果字符串中有中文，那么这种传统遍历会出现错误，传统遍历是按照字节来遍历的，但是汉字占3个字节
	// 解决办法就是将str转成切片
	// for的传统方式遍历字符串和数组
	var str1 = "woainia"
	for i:=0;i<len(str1);i++{
		fmt.Printf("%c \n", str1[i])
	}

	// for-range不会出现汉字错误的错误，因为这种默认是按照字符遍历的
	// 第二种for-range:index是索引，val是索引对应的值
	var str2 = "woainia祖国"
	for index, val:=range str2{
		fmt.Printf("%d, %c \n", index, val)
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
w 
o 
a 
i 
n 
i 
a 
0, w 
1, o 
2, a 
3, i 
4, n 
5, i 
6, a 
7, 祖 
10, 国 
[Finished in 1.3s]
