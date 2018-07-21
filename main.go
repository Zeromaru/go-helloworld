package main

import (
	"fmt"

	"github.com/zeromaru/gotools"
	"github.com/zeromaru/gotools/stringhelper"
)

func init() {
	fmt.Println("Init!")
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(gotools.Add(10, 2))
	fmt.Println(stringhelper.Upper("Dog"))
}
