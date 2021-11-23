package main

import (
	"fmt"
	"reflect"
)

func main() {
	//一定是双引号
	var name string = "测试生命"
	var name1 = "测试声明字符串"
	//var rotate1 = 123
	var rotate float32 = 0.34
	fmt.Println(typeof(rotate))
	fmt.Println("hello go!：", name, name1, rotate)
}

func typeof(key interface{}) reflect.Type {
	return reflect.TypeOf(key)
}
