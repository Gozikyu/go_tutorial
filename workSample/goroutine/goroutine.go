package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("task1のカウント:", i)
		time.Sleep(time.Microsecond * 500)
	}
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("task2のカウント:", i)
		time.Sleep(time.Microsecond * 300)
	}
}

func main() {
	fmt.Println("mainのスタート")

	var wg sync.WaitGroup

	wg.Add(2)

	go task1(&wg)
	go task2(&wg)

	wg.Wait()

	fmt.Println("mainの終了")
}
