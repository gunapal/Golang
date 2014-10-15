package main

import(
	"fmt"
)

type thread struct {
	threaId int
}

func printThreadID(t thread){
	
	fmt.Println("Thread ID :", t.threaId)
	for i := 0; i < 10; i++ {
		fmt.Println("T :", t.threaId, "on iteration : ", i)
	}
}

func main() {
	thread1 := thread{1}
	thread2 := thread{2}

	go printThreadID(thread1)
	go printThreadID(thread2)

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}