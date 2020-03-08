package main

import (
	"math/rand"
	"time"
)

var cache = map[int]Book
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i:=0; i < 10; i++ {
		id = rnd.Intn(10) + 1
	}
}

func queryCache(id int) (b Book, ok bool) {
	b, ok := cache[id]
	return
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			return b, true
		}
	}
	return Book{}, false
}