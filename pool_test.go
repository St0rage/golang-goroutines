package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Pool
func TestPool(t *testing.T) {
	pool := sync.Pool{
		// Default Value
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Dani")
	pool.Put("Veronica")
	pool.Put("Dian")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
