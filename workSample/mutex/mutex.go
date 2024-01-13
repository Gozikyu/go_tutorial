package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	data := make(map[string]string)
	var wg sync.WaitGroup

	wg.Add(2)

	// 2つのゴルーチンで並列で同時に書き込む
	go func() {
		defer wg.Done()
		fmt.Println("ゴルーチン1の開始")
		mu.Lock()
		data["key"] = "value"
		mu.Unlock()
		fmt.Println("ゴルーチン1の終了")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("ゴルーチン2の開始")
		mu.Lock()
		data["key"] = "another value"
		mu.Unlock()
		fmt.Println("ゴルーチン2の終了")
	}()

	wg.Wait()

	// 後に書き込まれた方が出力される（結果はランダム）
	fmt.Println(data["key"])
}
