package main

import(
	"fmt"
	"time"
)

//The arrow after the 'chan' keyword
//indicates write can only perform send/write
//to the channel
func write(chnl chan<- int){
	
	for i := 0; i < 10; i++ {
		chnl <- i
		fmt.Println("writer : ", i)
		time.Sleep(time.Second )
	}

	close(chnl)
}

//The arrow before the 'chan' keyword
//indicates read can only recieve from 
//channel
func read(chnl <-chan int){
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

	//Make a channel with buffering capacity of one.
	chnl := make(chan int, 1)

	go read(chnl)
	write(chnl)

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}