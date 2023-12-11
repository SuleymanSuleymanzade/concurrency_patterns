package main

import (
	"fmt"
	"sync"
	"time"
)

func stringy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 100_000; i++ {
		mutex.Lock()
		*money += 10
		mutex.Unlock()
	}
	fmt.Println("Singy Done")
}

func spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 100_000; i++ {
		mutex.Lock()
		*money -= 10
		mutex.Unlock()
	}
	fmt.Println("Spendy Done")
}

func main() {
	money := 100
	mutex := sync.Mutex{}
	go stringy(&money, &mutex)
	go spendy(&money, &mutex)
	time.Sleep(2 * time.Second)
	fmt.Println("Money in bank account ", money)

}
