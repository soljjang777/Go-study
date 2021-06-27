package main

import (
	"fmt"
	"goproject/usepkg/custompkg"
	"goproject/usepkg/exinit"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()

	s := custompkg.Student{}
	s.Name = "화랑"
	s.Age = 32
	fmt.Println(s.Name, s.Age)
	fmt.Println(custompkg.PI)
	custompkg.Ratio = 10
	fmt.Println(custompkg.Ratio)
	custompkg.PrintCustom2()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)

	custompkg.PrintCustom()
	exinit.PrintD()
}
