package bins

import "time"

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
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

		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}

}

func newBinArray(arr []Bin) *binList {
	return &binList{
		binArray: arr,
	}
}
