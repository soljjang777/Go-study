package main

import (
	"fmt"
)

type Bread struct {
	val string
}

type OrangeJam struct {
	opened bool
}

type SpoonOfOrangeberry struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	bread := make([]*Bread, num) // 포인터 배열타입을 만든다.
	for i := 0; i < num; i++ {
		bread[i] = &Bread{val: "빵"}
	}
	return bread
}

func OpenOrangeJam(jam *OrangeJam) {
	jam.opened = true
}

func GetOneSpoon(_ *OrangeJam) *SpoonOfOrangeberry {
	return &SpoonOfOrangeberry{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfOrangeberry) {
	bread.val += " + 오렌지잼"
}

func MakeSandwitch(bread []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(bread); i++ {
		sandwitch.val += bread[i].val + " + "
	}
	return sandwitch
}

func main() {
	// 빵을 두개 꺼낸다.
	breads := GetBreads(2)

	// 스트로베리 잼을 바른다.
	jam := &OrangeJam{}

	// 딸기잼 뚜겅을 연다
	OpenOrangeJam(jam)

	// 한스푼 뜬다
	spoon := GetOneSpoon(jam)

	// 딸기잼을 바른다
	PutJamOnBread(breads[0], spoon)

	// 빵을 덮는다.
	sandwitch := MakeSandwitch(breads)

	// 완성
	fmt.Println(sandwitch.val)
}
