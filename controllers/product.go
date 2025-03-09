package controllers

import (
	"go-laris/lib"
	"go-laris/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAllProduct(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 9999
	}

	listProduct := repository.FindAllProduct(search, page, limit)
	lib.HandlerOK(ctx, "Find All Product Success", nil, listProduct)
}

// func CreateProduct(ctx *gin.Context) {
// 	id := ctx.GetInt("userId")

// 	if id == 0 {
// 		lib.HandlerBadReq(ctx, "User Id Not Found")
// 		return
// 	}

// 	var form dtos.Product
// 	if err := ctx.ShouldBind(&form); err != nil {
// 		lib.HandlerBadReq(ctx, "Invalid Input Data")
// 		return
// 	}

// 	file, err := ctx.FormFile("image")
// 	var img *string
// 	if err == nil {
// 		allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
// 		fileExt := strings.ToLower(filepath.Ext(file.Filename))

// 		if !allowExt[fileExt] {
// 			lib.HandlerBadReq(ctx, "Invalid file extension")
// 			return
// 		}

// 		image := uuid.New().String() + fileExt
// 		root := "./img/product"

// 		if err := ctx.SaveUploadedFile(file, root+image); err != nil {
// 			lib.HandlerBadReq(ctx, "Upload Image Failed")
// 			return
// 		}
// 		baseURL := "http://localhost:8080"
// 		imgUrl := baseURL + "/img/profile/" + image
// 		img = &imgUrl

// 	}

// 	productData := dtos.Product{
// 		Image:        img,
// 		NameProduct:  form.NameProduct,
// 		Price:        form.Price,
// 		Discount:     form.Discount,
// 		CategoriesId: &id,
// 	}

// 	_, err = repository.CreateProduct(productData, id)
// 	if err != nil {
// 		fmt.Println(err)
// 		lib.HandlerBadReq(ctx, "Create Failed")
// 		return
// 	}

// 	lib.HandlerOK(ctx, "Create product success", productData, nil)
// }

func DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		lib.HandlerBadReq(ctx, "Invalid Product ID")
		return
	}

	dataProduct, err := repository.FindOneProductById(id)

	if err != nil {
		lib.HandlerBadReq(ctx, "Invalid Product ID")
		return
	}

	err = repository.DeleteProduct(id)
	if err != nil {
		lib.HandlerNotfound(ctx, "Id Not Found")
		return
	}

	lib.HandlerOK(ctx, "Product deleted successfully", nil, dataProduct)

}

func ListProduct(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	products, err := repository.SeeAllProduct(page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Respont{
			Success: false,
			Message: "Failed to find Products",
		})
		return
	}
	log.Println("Produk yang diambil:", products)
	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "List Products",
		Result:  products,
	})
}

func FindProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dataProduct := repository.FindOneProduct(id)
	if dataProduct.Id != 0 {

		ctx.JSON(http.StatusOK, lib.Respont{
			Success: true,
			Message: "Product Found",
			Result:  dataProduct,
		})
	} else {
		ctx.JSON(http.StatusNotFound, lib.Respont{
			Success: false,
			Message: "Product Not Found",
		})
	}
}
