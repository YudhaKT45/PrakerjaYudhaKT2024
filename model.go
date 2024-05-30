package main

import (
	"errors"
)

var Products = []*Product{}

type Product struct {
	ID    uint64 `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Price uint64 `json:"price" gorm:"column:price"`
}

func (u *Product) Validate() error {
	if u.Name == "" {
		return errors.New("invalid Name input")
	}

	if u.Price == 0 {
		return errors.New("invalid Price input")
	}
	return nil
}
