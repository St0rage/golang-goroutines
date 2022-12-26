package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// NewTicker
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	// tidak ideal
	for time := range ticker.C {
		fmt.Println(time)
	}

}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	// tidak ideal
	for time := range channel {
		fmt.Println(time)
	}
}
