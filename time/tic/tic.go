package main

//복잡스케쥴표 짤 때 씀(게임프로그램 할 때 많이씀)

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

	go push(c) // channel1을 push함수에  c인자로 넘기고 go Thred 시작

	timerchan := time.After(10 * time.Second)   // afrer 한번만 알려준다
	tickTimerchan := time.Tick(2 * time.Second) //tick 2초 간격 주기적으로 알려준다
	for {
		select {
		case v := <-c: // c에 입력값이 있으면 아래 v출력
			fmt.Println(v)
		case <-timerchan:
			fmt.Println("timed out")
			return
		case <-tickTimerchan: //tick을 시간주기적 알림
			fmt.Println("Tick")
		}
	}
}
