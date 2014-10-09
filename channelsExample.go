package main

import(
	"fmt"
)

func write(chnl chan int){
	
	for i := 0; i < 10; i++ {
		chnl <- i
		fmt.Println("writer : ", i)
	}
}

func read(chnl chan int){
	i := 0
	for ;i < 10; {
		val, isDataAvailable := <- chnl
		if isDataAvailable {
			fmt.Println("Reader :", val)
			i++
		}
	}
}

func main() {
	
	chnl := make(chan int, 1)

	go read(chnl)
	go write(chnl)

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}