package main

import (
	"log"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func (u *ProductRepo) Migrate() {

	err := u.DB.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal(err)
	}

}

func (u *ProductRepo) Get() ([]*Product, error) {
	products := []*Product{}
	err := u.DB.Debug().Model(&Product{}).Find(&products).Error
	return products, err
}

func (u *ProductRepo) Create(product *Product) error {
	err := u.DB.Debug().Model(&Product{}).Create(product).Error
	return err
}

func (u *ProductRepo) Update(product *Product) error {
	err := u.DB.Debug().Model(&Product{}).Where("id = ?", product.ID).Updates(product).Error
	return err
}

func (u *ProductRepo) Delete(product *Product) error {
	err := u.DB.Debug().Model(&Product{}).Delete(product, product.ID).Error
	return err
}
