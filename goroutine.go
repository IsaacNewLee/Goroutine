package main

import ("fmt"
		"time"
		)

func calc()  {
	for i:=0;i<10;i++{
		time.Sleep(1*time.Second)
		fmt.Println("run",i,"times")
	}
	fmt.Println("calc finish")
}

func main(){
	go calc() //加上go可以并发执行此函数
	fmt.Println("i exited")
	time.Sleep(11*time.Second)
}