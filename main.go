package main

import (
	"fmt"
	"time"

	"github.com/climanager/bins"
)

func main() {

	bin1 := bins.NewBin("1111", true, time.Now(), "Odin")
	bin2 := bins.NewBin("2222", false, time.Now(), "Thor")
	bin3 := bins.NewBin("3333", true, time.Now(), "Loki")

	binsSlice := []bins.Bin{
		*bin1,
		*bin2,
		*bin3,
	}

	fmt.Println(binsSlice)
	fmt.Println("Hello")
}
