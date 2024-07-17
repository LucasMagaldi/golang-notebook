package controller

import (
	"api/model"
	"api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "batata",
			Price: 15,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
