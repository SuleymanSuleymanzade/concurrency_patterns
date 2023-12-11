package main

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func matchRecorder(matchEvent *list.List, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		matchEvent.PushBack("MatchEvent " + strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(1 * time.Second)
		fmt.Println("appended match event")
	}
}

func clientHandler(mEvent *list.List, mutex *sync.RWMutex, st time.Time) {
	mutex.RLock()
	allEvents := copyAllEvents(mEvent)
	mutex.RUnlock()

	timeTaken := time.Since(st)
	fmt.Println(len(allEvents), "event copied in", timeTaken)
}

func copyAllEvents(matchEvents *list.List) []string {
	i := 0
	allEvents := make([]string, matchEvents.Len())
	for e := matchEvents.Front(); e != nil; e = e.Next() {
		allEvents[i] = e.Value.(string)
		i++
	}
	return allEvents
}

func main() {
	mutex := sync.RWMutex{}
	var matchEvents = list.New()
	for j := 0; j < 10000; j++ {
		matchEvents.PushBack("Match event")
	}
	go matchRecorder(matchEvents, &mutex)
	start := time.Now()

	for j := 0; j < 50000; j++ {
		go clientHandler(matchEvents, &mutex, start)
	}
	time.Sleep(100 * time.Second)

}
