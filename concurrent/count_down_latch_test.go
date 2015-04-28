package common

import (
	"fmt"
	"testing"
	"time"
)

func TestCDL(t *testing.T) {
	count := 3
	cdl := NewCountDownLatch(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			time.Sleep(3 * time.Second)
			cdl.CountDown()
			fmt.Printf("%d count down\n", i)
		}(i)
	}
	cdl.Await()
	fmt.Println("finished")
}
