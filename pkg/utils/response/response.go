package response

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type CategoryResponse struct {
	CategoryID int         `json:"category_id"`
	Category   string      `json:"category"`
	Data       interface{} `json:"data"`
}

type BrandResponse struct {
	BrandID int         `json:"brand_id"`
	Brand   string      `json:"brand"`
	Data    interface{} `json:"data"`
}

func ClientResponse(statusCode int, message string, data interface{}, err interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
		Error:   err,
	}
}

func CategoryResponseWithProduct(statusCode int, CategoryID int, CategoryName string, data interface{}) CategoryResponse {
	return CategoryResponse{
		CategoryID: CategoryID,
		Category:   CategoryName,
		Data:       data,
	}
}

func BrandResponseWithProduct(statusCode int, BrandID int, BrandName string, data interface{}) BrandResponse {
	return BrandResponse{
		BrandID: BrandID,
		Brand:   BrandName,
		Data:    data,
	}
}
