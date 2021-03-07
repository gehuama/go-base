package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// var aa = 3
// var ss = "kkk"
// var bb = true
var (
	aa = 3
	ss = "kk"
	bb = true
)

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitailValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	// c := 3 + 4i
	// fmt.Println(cmplx.Abs(c)) // 5

	fmt.Println(
		cmplx.Exp(1i*math.Pi) + 1)
	// cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

// go 只有强制类型转换

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

/**
 * 常量的定义
 * const 数值可做各种类型使用
 * const a,b = 3,4
 *
 * 使用常量定义枚举
 * 普通枚举类型
 * const (
		cpp = iota
		_
		payton
		goLang
		javaScript
	)

 * 自增值枚举类型
 * const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
*/
func consts() {
	const fileName = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(fileName, c)

}

func enums() {
	const (
		cpp = iota
		_
		payton
		goLang
		javaScript
	)

	// b ,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javaScript, payton, goLang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello go")
	variable()
	variableInitailValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enums()
}

/**
1. fmt库 Println打印函数并且换行

2. go 没有全局变量 aa，ss，bb  仅仅是包内不变量

3. 定义变量：
var关键字
var a,b,c
可放在函数內或直接放在包内
使用var()集中定义变量

4. 编译器可以自动决定类型
eg： var a, b, c, s = 3, 4, true, "def"

5. 使用:定义变量
eg a, b, c, s := 3, 4, true, "def"
注意：只能在函数內使用

*/

/**
 * 变量定义的要点
 * 变量类型写在变量名之后
 * 编译器可以推测变量的类型
 * 没有char 只有rune（32位）
 * 原生支持复数类型
 **/
