package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr [3]int8
	fmt.Println("arr值:", arr)
	fmt.Printf("arr:%p\n", &arr)    // 首元素的地址
	fmt.Println("arr[0]:", &arr[0]) // 相差一个字节
	fmt.Println("arr[1]:", &arr[1])
	fmt.Println("arr[2]:", &arr[2])
	// 取址
	var x int
	fmt.Printf("x的值:%d ", x)
	fmt.Printf("x的地址:%p ", &x)
	// 指针变量
	var p *int
	p = &x
	fmt.Printf("p的值:%p ", p)

	// 取值
	x = 10
	fmt.Printf("p指向的值:%d\n", *p)

	// 指针
	var a = 100
	var b = &a
	var c = &b
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c)

	p1 := 1
	p2 := &p1
	*p2++
	fmt.Println(p1)
	fmt.Println(*p2)
	// new
	var p3 *int   // 声明
	p3 = new(int) // 开辟空间
	fmt.Println(p3)
	fmt.Println(*p3)

}
