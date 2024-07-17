package controller

import (
	"api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
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
