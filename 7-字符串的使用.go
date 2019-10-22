package main
import "fmt"


// go中字符串是由单个字符拼接而成的,使用utf-8编码的
// 字符串的两种表示形式：1，双引号：可以识别转义字符，2.反引号，

func main() {

	// 字符串的使用
	str1 := "我爱你，中国！"
	// 字符串使用细节
	str1 = "我和你"
	// 使用双引号表示
	str2 := "我爱你，\n 母亲！"
	// 使用原始字符串输出,原始内容是啥，就输出啥。反引号是esc下面的按键
	str3 := `\t\nwoaini`
	// 字符串的拼接方式
	str4 := str1 + str2
	str4 += str3
	// 加好连接的字符比较多的时候，可以换行，但是加号要留在上面,不然报错b
	// 因为go默认会在后面加一个分号，但是看到有加号，就不加这个分号了。
	str4 = "woaini" + "woaini" +
	"woaini" + "woaini" + "zhongguo"


	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)
}









我和你
我爱你，
 母亲！
\t\nwoaini
woainiwoainiwoainiwoainizhongguo
[Finished in 0.8s]

