package handler

import (
	"net/http"
	"strconv"

	"github.com/ahdaan98/pkg/config"
	interfaces "github.com/ahdaan98/pkg/usecase/interface"
	"github.com/ahdaan98/pkg/utils/models"
	"github.com/ahdaan98/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	usecase interfaces.InventoryUseCase
}

func NewInventoryHandler(usecase interfaces.InventoryUseCase) *InventoryHandler {
	return &InventoryHandler{
		usecase: usecase,
	}
}

func (i *InventoryHandler) AddInventory(c *gin.Context) {
	var inventory models.AddInventory

	if err := c.ShouldBindJSON(&inventory); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error binding json format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	_, err := i.usecase.AddInventory(inventory)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to add a new product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added new inventory...", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) ListProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	per_product, _ := strconv.Atoi(c.Query("per_product"))

	k, err := i.usecase.ListProducts(page, per_product)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Failed to retrieve products", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "product list", k, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) EditInventory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var editInventory models.EditInventory

	if err := c.ShouldBindJSON(&editInventory); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error binding json format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	inv, err := i.usecase.EditInventory(editInventory, id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Failed to edit product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "edited product", inv, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) UpdateInventory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Updatedinventory models.UpdateInventory

	if err := c.ShouldBindJSON(&Updatedinventory); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error binding json format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	inv,err:=i.usecase.UpdateInventory(Updatedinventory,id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Failed to update product stock", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "updated product", inv, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) ShowIndividualProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	inv,err:=i.usecase.ShowIndividualProduct(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Failed to get product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "sucessfully retrieved the product by id", inv, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) CheckStock(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	inv,err:=i.usecase.CheckStock(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Failed to get product stock", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "sucessfully retrieved the product by id", inv, nil)
	c.JSON(http.StatusOK, successRes)
}


func (a *InventoryHandler) UploadProductImage(c *gin.Context) {
	cfg,_:=config.LoadEnvVariables()
	productID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
	}

	form, err := c.MultipartForm()
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse form"})
			return
	}

	files := form.File["image"]
	var urls []string

	for _, file := range files {
			filename := file.Filename

			// Save the uploaded file
			if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
					return
			}

			// Construct the URL for the uploaded image
			url := "http://localhost:"+cfg.PORT+"/admin/uploads/" + filename
			urls = append(urls, url)

			// Insert the image URL into the database
			if err := a.usecase.AddImage(productID, filename); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add image to database"})
					return
			}
	}

	c.JSON(http.StatusOK, gin.H{"success": "image added","data":urls})
}

func (h *InventoryHandler) ListProductsWithImages(c *gin.Context) {
    productList, err := h.usecase.ListProductsWithImages()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, productList)
}