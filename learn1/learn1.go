package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		name  string
		age   int
		adbot float32
		goole string = "goole_test"
	)
	name = "name"
	age = 45
	adbot = 0.98
	fang := "房子"
	zi := "子"
	shijian, jinqian := "时间", 233
	fmt.Println(name, age, adbot, goole, fang, shijian, jinqian, zi)
	ptr := new(int)
	fmt.Println("ptr address: ", ptr)
	fmt.Println("ptr value: ", *ptr) // * 后面接指针变量，表示从内存地址中取出值
	fmt.Println(newInt())
	fmt.Println(*newInt())
	fmt.Println(oldInt())
	fmt.Println(*oldInt())
	var num01 int = 0b1100
	var num02 int = 0o14
	var num03 int = 0xC

	/*
		%b    表示为二进制
		%c    该值对应的unicode码值
		%d    表示为十进制
		%o    表示为八进制
		%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
		%x    表示为十六进制，使用a-f
		%X    表示为十六进制，使用A-F
		%U    表示为Unicode格式：U+1234，等价于"U+%04X"
		%E    用科学计数法表示
		%f    用浮点数表示
	*/
	fmt.Printf("2进制数 %b 表示的是: %d \n", num01, num01)
	fmt.Printf("8进制数 %o 表示的是: %d \n", num02, num02)
	fmt.Printf("16进制数 %X 表示的是: %d \n", num03, num03)

	var a byte = 65
	// 8进制写法: var a byte = '\101'     其中 \ 是固定前缀
	// 16进制写法: var a byte = '\x41'    其中 \x 是固定前缀

	var b uint8 = 66
	fmt.Printf("a 的值: %c ,b 的值: %c \n", a, b)
	var a1 byte = 'A'
	var b1 uint8 = 'B'
	fmt.Printf("a 的值: %c ,b 的值: %c \n", a1, b1)

	// rune，占用4个字节，共32位比特位，所以它和 int32 本质上也没有区别。它表示的是一个 Unicode字符（Unicode是一个可以表示世界范围内的绝大部分字符的编码规范
	var a2 byte = 'A'
	var b2 rune = 'B'
	fmt.Printf("a 占用 %d 个字节数\nb 占用 %d 个字节数", unsafe.Sizeof(a2), unsafe.Sizeof(b2))

	// Go 中单引号与 双引号并不是等价的
	// 单引号用来表示字符，在上面的例子里，如果你使用双引号，就意味着你要定义一个字符串
	var bys rune = '中'
	fmt.Println(bys)

	var mystr01 string = "hello"
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr01: %s\n", mystr01)
	fmt.Printf("mystr02: %s", mystr02)
	// 以上虽然我都用双引号表示 一个字符串，但这并不是字符串的唯一表示方式。
	//
	//除了双引号之外 ，你还可以使用反引号。
	//
	//大多情况下，二者并没有区别，但如果你的字符串中有转义字符\ ，这里就要注意了，它们是有区别的
	var mystr03 string = "\\r\\n"
	var mystr04 string = `\r\n`
	fmt.Println(mystr03)
	fmt.Println(mystr04)
	fmt.Printf("%s 的解释型字符串是： %q \n", mystr04, mystr04)
	var txt = `是电瓶车手动出你的VS等你DVDVN的四渡赤水的基础上大V热皮uvbevn
    iodnv；大V你的搜是d第三次掉策划错愕吧我机电产品弄错而不被是uebrvdfvhdfvdfvndiv \n`
	fmt.Println(txt)

}

// 使用 new
func newInt() *int {
	return new(int)
}

// 使用传统的方式
func oldInt() *int {
	var dummy int
	return &dummy
}
