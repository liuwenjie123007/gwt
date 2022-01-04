package main

import (
	"errors"
	"fmt"
)

func main() {
	// 定义变量
	// var variableName type
	// var vname1, vname2, vname3 type
	// var variableName type = value
	// var vname1, vname2, vname3 type= v1, v2, v3
	// var vname1, vname2, vname3 = v1, v2, v3
	// vname1, vname2, vname3 := v1, v2, v3
	// _, b := 34, 35
	//var i int

	// 常量
	// const constantName = value
	const Pi float32 = 3.1415926
	const i = 1000
	const MaxThread = 10
	const prefix = "astaxie_"

	// 内置基础类型
	//var isActive bool
	//var enabled, disabled = true, false
	//var available bool  // 一般声明
	//valid := false      // 简短声明
	//available = true    // 赋值操作
	//var c complex64 = 5 + 5i
	////output: (5+5i)
	//fmt.Printf("Value is: %v", c)

	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	// 错误类型
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

}
