package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.WaitGroup
func RunAsynchronous(group *sync.WaitGroup) {

	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	/**
	wait() akan menunggu sampai perulangan goroutine diatas selesai atau jika add(100)
	maka done() = 100, done akan dikurangi 1 persatu sampai ke 0, dan jika sudah sudah sampai 0
	maka wait() akan berhenti. jika tidak di done wait() akan terus dijalankan dan akan terjadi deadlock
	*/
	group.Wait()
	fmt.Println("Selesai")
}
