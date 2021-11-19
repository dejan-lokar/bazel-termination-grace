package tools

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestTrap(t *testing.T) {
	c := make(chan os.Signal, 10)
	signal.Notify(c)
	go func() {
		fmt.Printf("Waiting for a signal... PID: %d", os.Getpid())
		for s := range c {
			fmt.Printf("Trapped: %s\n", s)
		}
	}()

	time.Sleep(1 * time.Hour)
}
