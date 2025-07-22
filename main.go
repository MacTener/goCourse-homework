package main

import (
	"fmt"
	"time"
)

type bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type binList struct {
	binArray []bin
}

func newBin(
	id string,
	private bool,
	createdAt time.Time,
	name string,
) *bin {

	return &bin{

		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}

}

func newBinArray(arr []bin) *binList {
	return &binList{
		binArray: arr,
	}
}

func main() {

	bin1 := newBin("1111", true, time.Now(), "Odin")
	bin2 := newBin("2222", false, time.Now(), "Thor")
	bin3 := newBin("3333", true, time.Now(), "Loki")

	binsSlice := []bin{
		*bin1,
		*bin2,
		*bin3,
	}

	fmt.Println(binsSlice)

}
