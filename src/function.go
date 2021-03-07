package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
函数
func eval(a b int, op string) int
*/

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("不支持的操作：%s", op)
	}
}

/**
函数返回多个值时可以起名字
仅用于非常简单的函数
对于调用者而言没有区别
**/
// 13 / 3 =4 ...1
func div(a, b int) (q, r int) {
	return a / b, a % b
	// q = a / b
	// r = a % b
	// return
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/*
指针
指针不能运算
参数传递
值传递？引用传递？
go语言只有值传递一种方式
*/
func swap1(a, b int) {
	b, a = a, b
}

func swap2(a, b *int) {
	*b, *a = *a, *b
}

func swap3(a, b int) (int, int) {

	return b, a
}

func main() {
	fmt.Println(eval(3, 4, "/"))

	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3))
	a, b := 3, 4
	// swap1(a, b)
	// fmt.Println(a, b)
	// swap2(&a, &b)
	// fmt.Println(a, b)
	a, b = swap3(a, b)
	fmt.Println(a, b)
}

/*
函数语法要点
返回值类型写在最后面

可返回多个值

函数作为参数

没有默认参数，可选参数

*/
