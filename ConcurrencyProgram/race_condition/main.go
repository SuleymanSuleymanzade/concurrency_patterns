package main

import (
	"fmt"
	"time"
)

func stringy(money *int) {
	for i := 0; i < 100_000; i++ {
		*money += 10
	}
	fmt.Println("Singy Done")
}

func spendy(money *int) {
	for i := 0; i < 100_000; i++ {
		*money -= 10
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	go stringy(&money)
	go spendy(&money)
	time.Sleep(2 * time.Second)
	fmt.Println("Money in bank account ", money)

}
