package main
import "fmt"


// 函数的注意事项和细节讨论：
/*
函数的形参和返回值可以是多个
返回的舒蕾可以是值类型和引用类型
函数名大写表示共有的，小写表示私有的
函数的变量是局部的，函数外不能生效
基本数据类型和数组默认都是值传递的，即进行值拷贝，在函数内修改，不会影响到原来的值
如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量，从效果上看类似引用传递
go函数不支持函数重载（不允许函数名重名）

2、go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
函数及人事一种数据类型，函数也可以作为形参，并且调用；
为了简化数据类型，go支持自定义数据类型，基本语法：
type 自定义数据类型名字 数据类型  理解：相当于一个别名
例子：type myint int // 这是myint 就等价int来使用了，但是go中还是认为这两个是不同的类型
如果需要这两种类型赋值的话， 还是需要进行数据类型转换的。

支持对函数返回值命名：优点：return后面不用再写返回值名字了，sum不用重新定义了，也就是说不用写：= 号了
func cal(num int) (sum int){
	sum = num + 10
	return 
}

go中支持可变参数：
案例演示：func sum(n1 int, args... int) int{
	函数体
}


*/

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


// 使用指针进行引用类型传递的效果
func yin(n *int){
	*n = *n + 10
	fmt.Println(*n)
}


// 通过函数中调用函数名
func diao(myfunc func(int) int, num3 int) int {
	return myfunc(num3)
}


// 函数自定义数据类型使用在函数名传递的时候，方便
type myFunc func(int) int
// 通过自定义函数类型，进行函数名调用
func dia2(myfun myFunc, num3 int) int {
	return myfun(num3)
}


// 可变参数的函数,args的名字可以随便改，相当于列表的名字
// 可变参数必须写在形参列表最后，否则无法通过
func sumn(num8 int, args... int) int {
	end := num8
	// 遍历args这个切片
	for i:=0; i<len(args); i++{
		end += args[i]  // args[i] 表示取出切片的第一个元素值，切片是可变的数组列表
	}
	return end
}

// 数组传递都是值传递，
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

	// 使用引用累心进行引用传递
	num := 10
	yin(&num)
	fmt.Println(num)

	num2 := 20
	bieming := yin
	bieming(&num2)
	fmt.Println(num2)

	// 调用函数的引用，在函数中使用别的函数
	num4 := 7
	num5 := diao(hanshu, num4)
	fmt.Println(num5)

	// 函数自定义数据类型使用在函数名传递的时候，方便
	// type myFunc func(int) int
	num6 := 7
	num7 := dia2(hanshu, num6)
	fmt.Println(num7)

	// 支持可变参数
	end := sumn(1, 4, 5, 6)
	fmt.Println(end)


}







3
7
1534
20
20
30
30
255
255
16
[Finished in 3.5s]
