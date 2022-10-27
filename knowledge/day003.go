package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 1. 整型转整型
	var num int32 = 200
	num1 := int16(num)
	fmt.Println(num1, reflect.TypeOf(num1))
	// 2. 整型转字符串
	num2 := 100
	string2 := strconv.Itoa(num2)
	fmt.Println(string2, reflect.TypeOf(string2))
	// 3. 字符串转整型
	ageStr := "22"
	age, err := strconv.Atoi(ageStr)
	fmt.Println("err:", err)
	fmt.Println(age, reflect.TypeOf(age))

	age1, err := strconv.ParseInt("44", 10, 8)
	fmt.Println("err1:", err)
	fmt.Println(age1, reflect.TypeOf(age1))
	// 4. 字符串转浮点型
	float1, err := strconv.ParseFloat("3.1415926", 64)
	fmt.Print(float1, reflect.TypeOf(float1))
	fmt.Println("err:", err)
	// 5. 字符串转布尔型
	bool1, err := strconv.ParseBool("1")
	fmt.Print(bool1, reflect.TypeOf(bool1))
	fmt.Println(" err3:", err)
	bool2, err := strconv.ParseBool("0")
	fmt.Print(bool2, reflect.TypeOf(bool2))
	fmt.Println(" err4:", err)
	bool3, err := strconv.ParseBool("true")
	fmt.Print(bool3, reflect.TypeOf(bool3))
	fmt.Println(" err5:", err)
	bool4, err := strconv.ParseBool("false")
	fmt.Print(bool4, reflect.TypeOf(bool4))
	fmt.Println(" err6:", err)
	// Scan
	var name string
	fmt.Scanln(&name, &age) // 换行即结束
	fmt.Println("姓名ln:", name, "年龄ln:", age)
	fmt.Scan(&name, &age)
	fmt.Println("姓名:", name, "年龄:", age)
	var number1, number2 int
	fmt.Print("按指定格式输入:")
	fmt.Scanf("%d+%d", &number1, &number2)
	fmt.Println(number1 + number2)

}
