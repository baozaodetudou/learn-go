package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 数组
	//数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。因为数组的长度是固定的，所以在Go语言中很少直接使用数组
	var arr [2]float32
	arr[0] = 2
	arr[1] = 3.4
	fmt.Printf("%d 的类型是%T \n", arr, arr)
	fmt.Println(reflect.TypeOf(arr))
	var arr1 [3]int = [3]int{2, 3, 4}
	fmt.Println(arr1)
	fmt.Printf("%d 的类型是 %T \n", arr1, arr1)
	var arr2 = [3]int{4, 5, 6}
	fmt.Println(arr2)
	fmt.Printf("%d 的类型是%T \n", arr2, arr2)
	arr3 := [4]float32{3.2, 4.3, 5.4, 6.5}
	fmt.Println(arr3)
	fmt.Printf("%d 的类型是%T \n", arr3, arr3)

	arr4 := [...]int{6, 7, 8}
	arr5 := [...]int{7, 8, 9, 10}
	fmt.Printf("%d 的类型是%T \n", arr4, arr4)
	fmt.Printf("%d 的类型是%T \n", arr5, arr5)

	type arr6 [2]int
	myarr := arr6{1, 2}
	fmt.Printf("%d 的类型是%T \n", myarr, myarr)

	// [4]int{2:3}，4表示数组有4个元素，2 和 3 分别表示该数组索引为2（初始索引为0）的值为3，而其他没有指定值的，就是 int 类型的零值，即0
	artful := [4]int{2: 3}
	fmt.Printf("%d 的类型是%T \n", artful, artful)

	// 切片
	// 切片（Slice）与数组一样，也是可以容纳若干类型相同的元素的容器。与数组不同的是，无法通过切片类型来确定其值的长度。每个切片值都会将数组作为其底层数据结构。我们也把这样的数组称为切片的底层数组
	cutarr := [...]int{2, 3, 4}
	fmt.Printf("%d 的类型是 %T \n", cutarr, cutarr)
	fmt.Printf("%d 的类型是 %T \n", cutarr[0:2], cutarr[0:2])
	// 切片是对数组的一个连续片段的引用，所以切片是一个引用类型，这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集
	// 左闭右开的区间

	//构造方法
	cutarr1 := [...]int{1, 2, 3, 4, 5}
	cut1 := cutarr1[1:2]
	fmt.Println(cut1, "\n")
	// 在切片时，若不指定第三个数，那么切片终止索引会一直到原数组的最后一个数。
	// 而如果指定了第三个数，那么切片终止索引只会到原数组的该索引值
	cut2 := cutarr1[0:2:4]
	print(cut2, "\n")
	fmt.Printf("%d", cut2)

	// 声明字符串类型切片
	var strlist []string
	fmt.Printf("%d", strlist)
	// 声明整型切片
	var numList []int
	fmt.Print(numList)
	// 声明一个空切片
	var numListEmpty = []int{}
	fmt.Println(numListEmpty, "\n")

	// 使用 make 函数构造，make 函数的格式：make( []Type, size, cap )
	//
	//这个函数刚好指出了，一个切片具备的三个要素：类型（Type），长度（size），容量（cap）
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
	fmt.Println(cap(a), cap(b))

	a1 := []int{4: 2}
	fmt.Println(a1)
	fmt.Println(len(a1), cap(a1))

	//关于 len 和 cap 的概念，可能不好理解 ，这里举个例子：
	//
	//公司名，相当于字面量，也就是变量名。
	//
	//公司里的所有工位，相当于已分配到的内存空间
	//
	//公司里的员工，相当于元素。
	//
	//cap 代表你这个公司最多可以容纳多少员工
	//
	//len 代表你这个公司当前有多少个员工
	var myarr4 []int
	fmt.Println(myarr4 == nil)

	myarr5 := []int{1}
	// 追加一个元素
	myarr5 = append(myarr5, 2)
	// 追加多个元素
	myarr5 = append(myarr5, 3, 4)
	// 追加一个切片, ... 表示解包，不能省略
	myarr5 = append(myarr5, []int{7, 8}...)
	// 在第一个位置插入元素
	myarr5 = append([]int{0}, myarr5...)
	// 在中间插入一个切片(两个元素)
	myarr5 = append(myarr5[:5], append([]int{5, 6}, myarr5[5:]...)...)
	fmt.Println(myarr5)

	//
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, len(myslice))

	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])

}
