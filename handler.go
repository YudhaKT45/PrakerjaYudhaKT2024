package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHdl struct {
	Repository *ProductRepo
}

func (u *ProductHdl) GetGorm(ctx *gin.Context) {
	products, err := u.Repository.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (u *ProductHdl) CreateGorm(ctx *gin.Context) {
	product := &Product{}
	if err := ctx.Bind(product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "invalid body request",
		})
		return
	}
	if err := product.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": err,
		})
		return
	}
	err := u.Repository.Create(product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusCreated, product)
}

func (u *ProductHdl) UpdateGorm(ctx *gin.Context) {
	products, err1 := u.Repository.Get()
	if err1 != nil {
		panic(err1)
	}
	productID, err2 := strconv.Atoi(ctx.Param("id"))
	if err2 != nil {
		panic(err2)
	}
	product := &Product{}
	if err := ctx.Bind(product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "invalid body request",
		})
		return
	}
	if err := product.Validate(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": err,
		})
		return
	}

	isExist := false
	for _, pdr := range products {
		fmt.Println(pdr.ID)
		if pdr.ID == uint64(productID) {
			isExist = true
			break
		}
	}
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]any{
			"message": "product with id not found",
		})
		return
	}

	product.ID = uint64(productID)
	err := u.Repository.Update(product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong???",
		})
		return
	}
	ctx.JSON(http.StatusCreated, product)
}
func (u *ProductHdl) DeleteGorm(ctx *gin.Context) {
	products, err1 := u.Repository.Get()
	if err1 != nil {
		panic(err1)
	}
	productID, err2 := strconv.Atoi(ctx.Param("id"))
	if err2 != nil {
		panic(err2)
	}
	product := &Product{}
	isExist := false
	for _, pdr := range products {
		fmt.Println(pdr.ID)
		if pdr.ID == uint64(productID) {
			isExist = true
			break
		}
	}
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]any{
			"message": "product with id not found",
		})
		return
	}
	product.ID = uint64(productID)
	err := u.Repository.Delete(product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusCreated, product)
}
