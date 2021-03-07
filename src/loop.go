package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * 循环
 * for 语句
 * for的条件里不需要括号
 * for的条件里可以省略初始条件，结束条件，递增表达式

 * 1. convertToBin
 * 省略初始条件，相当于while
 * 2. printFile
 * 省略初始条件，相当于while
 * 3. forEver
 * for {
		fmt.Println("abc")
	}
 * 无限循环
*/
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forEver() {
	i := 0
	for {
		fmt.Println("abc")
		i++
		if i > 3 {
			return
		}
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011
		convertToBin(1024))
	printFile("abc.txt")
	forEver()
}

/**
基本语法
for if 后面的条件没有括号
if 条件也可定义变量
没有while
switch不需要break，也可以直接switch 多条件
*/
