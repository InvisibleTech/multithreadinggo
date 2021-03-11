package main

import (
	"sync"
	"time"
)

var (
	money = 100
	lock  = sync.Mutex{}
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}

	println("Stingy done.")
}
func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money -= 10
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}

	println("Spendy done.")
}
func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	println(money)
}