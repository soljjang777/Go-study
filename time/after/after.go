package main //챗만들때 많이 씀

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++ // 1초간격으로 0부터 시작 1씩 증가하는 인티져(정수형)
	}
}

func main() {
	c := make(chan int) // channel1생성

	go push(c) // channel1을 push함수에 c인자로 넘기고 go Thred시작

	timerchan := time.After(10 * time.Second) // 특정시간 주기적 반환하는 타이머 채널을 만듬

	for {
		select {
		case v := <-c: // c에 입력값이 있으면 아래 v출력
			fmt.Println(v)
		case <-timerchan: //만든타임채널을 주기적으로 살펴본다
			fmt.Println("timed out")
			return
		}
	}
}
