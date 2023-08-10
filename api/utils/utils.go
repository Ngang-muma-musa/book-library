package utils

import (
	"book-library/api/types"
	"errors"
)

func GetById(id string, books []types.Book) (*types.Book,error){
	for i, b:= range books {
		if b.Id == id {
			return &books[i], nil
		}
		
	}
	return nil,errors.New("book bot found")
}

func GetItemIndex (id string, books []types.Book) (int,error){
	for i, b:= range books {
		if b.Id == id {
			return i,nil
		}
	}
	return 0,errors.New("book bot found")
}