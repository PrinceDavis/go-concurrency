package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	for i:=0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func (id int, wg *sync.WaitGroup)  {
			if b, ok := queryCache(id); ok {
				fmt.Println("From cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)

		go func (id int, wg *sync.WaitGroup)  {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("From database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)

		// time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCache(id int) (b Book, ok bool) {
	b, ok = cache[id]
	return
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}