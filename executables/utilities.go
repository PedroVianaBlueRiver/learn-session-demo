package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func serializeJson(v any) ([]byte, error) {
	contentJson, err := json.Marshal(v)
	if err != nil {

		log.Fatal(err)
		return nil, err
	}
	return contentJson, nil
}

func deserializeJson(content []byte, v any) error {
	model := v
	err := json.Unmarshal([]byte(content), &model)
	if err != nil {

		log.Fatal(err)
		return err
	}
	return nil
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

// func main() {

// 	done := make(chan bool)
// 	fmt.Println("Main going to call hello go goroutine")
// 	go hello(done)
// 	<-done
// 	fmt.Println("Main received data")

// 	// number := 589
// 	// sqrch := make(chan int)
// 	// cubech := make(chan int)
// 	// go calcSquares(number, sqrch)
// 	// go calcCubes(number, cubech)
// 	// squares, cubes := <-sqrch, <-cubech
// 	// fmt.Println("Final output", squares+cubes)
// }
