package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func task1(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		// メッセージをチャネルに送信
		ch <- fmt.Sprintf("task1のカウント: %d", i)
		time.Sleep(time.Millisecond * 500)
	}

	fmt.Println("task1終了")
}

func task2(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("task2のカウント: %d", i)
		time.Sleep(time.Millisecond * 300)
	}

	fmt.Println("task2終了")
}

func getGoroutineNum() {
	numGoroutines := runtime.NumGoroutine()
	fmt.Printf("実行されているゴルーチンの数: %d\n", numGoroutines)
}

func main() {
	fmt.Println("mainの開始")
	getGoroutineNum()

	var wg sync.WaitGroup
	wg.Add(2)

	messageChannel := make(chan string)

	fmt.Println("tasksのゴルーチン開始")
	go task1(messageChannel, &wg)
	go task2(messageChannel, &wg)
	getGoroutineNum()

	// ゴルーチンとして実行しないとwg.Waitでmainゴルーチンが待機状態になった時点で、デットロックのエラーになってしまう
	go func() {
		fmt.Println("waitのゴルーチン開始")
		getGoroutineNum()

		// ゴルーチンが終了するまで待機
		wg.Wait()
		fmt.Println("全てのtask終了")
		getGoroutineNum()

		close(messageChannel)
	}()

	for message := range messageChannel {
		fmt.Println(message)
	}

	fmt.Println("waitのゴルーチン終了")
	getGoroutineNum()

	fmt.Println("mainの終了")
}
