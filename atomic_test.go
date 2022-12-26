package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// Contoh race condition dengan solusi Atomic
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				// x = x + 1
				/**
				kode dibawah sama dengan kita menggunakan mutex,
				untuk mengihandari race condition pada perubahan data primivite

				Mutex lebih cocok untuk data berupa struct
				*/
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}
