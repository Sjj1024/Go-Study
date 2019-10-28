package main
import (
	"fmt"
)


// 标识符命名规范：凡是可以自己起名字的地方都叫标识符
/*
1.由数字字母下划线组成，字母不能开头，
2.严格区分大小写，也就是说num和Num是两个不同的变量，
3.标识符不能包含空格
4.下划线_在go中是一个特殊的标识符，成为空标识符。可以代表任何其它的标识符，但是
它对应的值会被忽略，所以仅能被作为占位符使用，不能作为标识符使用。
5.不能以系统保留关键字作为标识符

标识符命名规范：
1.包名：保持和package的名字和目录文件夹名字保持一致；尽量采取有意义的名字，
2.变量名、函数名、常量名：采用驼峰法命名：var stuNmame string = " "
3.如果变量名、函数名、常量名首字符大写，则可以被其他的包访问，如果首字母小写，则只能在本包里使用
（可以理解为，首字母大写是公开的，首字母小写是私有的），在go中没有pblic，private等关键字
4.系统保留关键字25个：和预定义标识符36个

*/



func main() {
	Num_a := 15

	fmt.Println(Num_a)
}
