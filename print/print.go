package main

import (
	"fmt"
)

func main() {

	var num int = 10
	var name string = "hwan"
	var byte1 byte = 'A'
	var boolean bool = true
	var f float32 = 1.5
	var p *int = &num

	fmt.Printf("%d\n", num)     //output : 10 (정수)
	fmt.Printf("%f\n", f)       //output : 1.500000(실수)
	fmt.Printf("%s\n", name)    //output : hwan(문자열)
	fmt.Printf("%b\n", byte1)   //output : 1000001(2진수)
	fmt.Printf("%0o\n", byte1)  //output : 101(8진수)
	fmt.Printf("%x\n", byte1)   //output : 41(16진수)
	fmt.Printf("%c\n", byte1)   //output : A(char)
	fmt.Printf("%t\n", boolean) //output : true(boolean)
	fmt.Printf("%U\n", byte1)   //output : U+0041(유니코드)
	fmt.Printf("%p\n", &p)      //output : 0xc000006028(포인터가 가르키는 주소 값 = num의 주소 값)
	fmt.Printf("%p\n", p)       //output : 0xc000014088(포인터 주소 값)

	var name string = "hwan"
	var str1 string = "abc\n" +
		"def"
	var str2 string = `abc\n` + name
	var str3 string = `qqq\n
	zzzz`

	fmt.Println(`hello` + name) //output : hellohwan
	fmt.Println(str1)           //output : abc
	//		   def
	fmt.Println(str2) //output : abc\nhwan
	fmt.Println(str3) //output : qqq\n
	//				zzzz
}
