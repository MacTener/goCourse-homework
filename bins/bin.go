package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type binList struct {
	binArray []Bin
}

func NewBin(
	id string,
	private bool,
	createdAt time.Time,
	name string,
) *Bin {

	return &Bin{

		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}

}

func newBinArray(arr []Bin) *binList {
	return &binList{
		binArray: arr,
	}
}
