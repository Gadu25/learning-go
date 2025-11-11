package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wait group
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

// concurrency
func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// dbCall(i)
		// spawn a new go routine
		wg.Add(1)
		go dbCall(i)
	}
	// this waits for the counter down to 0, all the tasks were completed.
	// and the rest of the code will execute
	wg.Wait()
	fmt.Printf("\nTotal execution time %v", time.Since(t0))
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	// decrements the counter/wg
	wg.Done()
}
