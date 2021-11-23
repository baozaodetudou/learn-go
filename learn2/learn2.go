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

}
