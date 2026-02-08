package main

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func newBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	if id == "" {
		return nil, errors.New("Идентификатор не может быть пустой строкой")
	}

	if name == "" {
		return nil, errors.New("Имя не может быть пустой строкой")
	}

	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}, nil
}

func newBinList(bins []Bin) (*BinList, error) {
	if bins == nil {
		return nil, errors.New("Список bin'ов не должен быть nil")
	}

	return &BinList{
		bins: bins,
	}, nil
}

func main() {

}
