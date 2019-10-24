package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	l := sync.RWMutex{}
	wg := sync.WaitGroup{}

	n := 10

	fmt.Printf("### Take RW Lock\n\n")

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			l.Lock()
			defer l.Unlock()

			fmt.Printf("num: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}(i)
	}

	wg.Wait()

	fmt.Printf("\n\n### Take R Lock\n\n")

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			l.RLock()
			defer l.RUnlock()

			fmt.Printf("num: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}(i)
	}

	wg.Wait()

	fmt.Printf("\n\n### Try to Take W Rock when R Lock is already aquired\n\n")

	wg.Add(1)
	go func() {
		defer wg.Done()

		l.RLock()
		defer l.RUnlock()

		fmt.Println("take R lock 1")

		time.Sleep(500 * time.Millisecond)
	}()

	time.Sleep(100 * time.Microsecond)
	wg.Add(1)
	go func() {
		defer wg.Done()

		l.Lock()
		defer l.Unlock()

		fmt.Println("take W lock")

		time.Sleep(500 * time.Millisecond)
	}()

	time.Sleep(100 * time.Microsecond)
	wg.Add(1)
	go func() {
		defer wg.Done()

		l.RLock()
		defer l.RUnlock()

		fmt.Println("take R lock 2")
	}()

	wg.Wait()
}
