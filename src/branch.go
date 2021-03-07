package main

import (
	"fmt"
	"io/ioutil"
)

func ifFunc() {
	const fileName = "abc.txt"
	// contents, err := ioutil.ReadFile(fileName)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// fmt.Printf("%s\n", contents) // contents出了if不存在
}

/**
 * switch 会自动break，除非使用fallthrough

 * switch 后面没有表达式
**/
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("警告 分数错误：%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	ifFunc()
	fmt.Println(
		"小明成绩等级：", grade(0),
		"小红成绩等级：", grade(53),
		"小黑成绩等级：", grade(63),
		"小蓝成绩等级：", grade(78),
		"小白成绩等级：", grade(88),
		"小紫成绩等级：", grade(98),
		"小绿成绩等级：", grade(-3),
	)
}
