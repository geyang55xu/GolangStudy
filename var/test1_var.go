package main

import "fmt"

//定义全局变量，全局变量不可以用:=去定义
var hh int = 0
var ii = "123"

var (
	pp int    = 123
	qq string = "123"
)

func main() {
	//定义int类型的变量
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//定义int类型的变量同时赋值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	//定义string类型的变量同时赋值
	var bb string = "123"
	fmt.Printf("bb = %s,type of bb = %T\n", bb, bb)

	//省略变量类型定义变量
	var cc = "123"
	fmt.Printf("cc = %s,type of cc = %T\n", cc, cc)
	var dd = 123
	fmt.Printf("dd = %d,type of dd = %T\n", dd, dd)

	//省略var关键字定义变量
	ee := 123
	fmt.Printf("ee = %d,type of ee = %T\n", ee, ee)
	ff := "123"
	fmt.Printf("ff = %s,type of ff = %T\n", ff, ff)
	gg := 3.14
	fmt.Printf("gg = %f,type of gg = %T\n", gg, gg)

	//全局变量
	fmt.Printf("hh = %d,type of hh = %T\n", hh, hh)
	fmt.Printf("ii = %s,type of ii = %T\n", ii, ii)

	//一次声明多个变量
	var jj, kk = 123, "123"
	fmt.Printf("jj = %d,type of jj = %T\n", jj, jj)
	fmt.Printf("kk = %s,type of kk = %T\n", kk, kk)
	var ll, mm int = 1, 2
	fmt.Printf("ll = %d,type of ll = %T\n", ll, ll)
	fmt.Printf("mm = %d,type of mm = %T\n", mm, mm)

	//多行的多变量声明
	var (
		nn int    = 123
		oo string = "123"
	)
	fmt.Printf("nn = %d,type of nn = %T\n", nn, nn)
	fmt.Printf("oo = %s,type of oo = %T\n", oo, oo)
	fmt.Printf("pp = %d,type of pp = %T\n", pp, pp)
	fmt.Printf("qq = %s,type of qq = %T\n", qq, qq)
}
